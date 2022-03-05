package models

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"
	"strconv"
	"time"
)

type Flow struct {
	Pass         *Pass
	Parray       []PassLoaded
	PassLoaded   *PassLoaded
	ID           string
	RID          string
	ImagesMap    *[]ImagesMap
	ImagesLoaded []*ImageLoaded
	ImageLoaded  *ImageLoaded
	Images       *[][]byte
	ImgMapPA     *map[string]string
	ImgMapIA     *map[string][]int
	Err          error
	Warning      error
	ImgExpects   []int
	ErrStatus    int
	GetWith      byte
	GetBy        string
	GetByFIO     [3]string
}

func (f *Flow) ImgData() *Flow {
	if f.Err != nil {
		return f
	}
	imgMap := make(map[string][]int)
	for key, title := range *f.ImgMapPA {
		imgId, err := strconv.Atoi(key)
		if err != nil {
			f.Err = fmt.Errorf("%w", err)
			return f
		}
		imgMap[title] = append(imgMap[title], imgId)
	}
	f.ImgMapIA = &imgMap
	return f
}

//Загружает изображения по URL
func (f *Flow) GetImage() *Flow {
	if f.Err != nil {
		return f
	}
	var imgArray [][]byte
	for idx, elem := range f.Pass.Images {
		if f.Warning.Error() != "" {
			flag := false
			for _, elem := range f.ImgExpects {
				if elem == idx {
					flag = true
				}
			}
			if flag {
				continue
			}
		}
		response, err := http.Get(elem.URL)
		if err != nil {
			f.Err = fmt.Errorf("%w", err)
			return f
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			f.Err = fmt.Errorf("%w", errors.New("невозможно загрузить изображение по адресу: "+elem.URL))
			return f

		}
		var bytes []byte
		bytes, f.Err = ioutil.ReadAll(response.Body)
		if f.Err != nil {
			f.Err = fmt.Errorf("%w", err)
			return f
		}
		var base64Encoding string
		mimeType := http.DetectContentType(bytes)
		base64Encoding += "data:" + mimeType + ";base64,"
		base64Encoding += base64.StdEncoding.EncodeToString(bytes)
		imgArray = append(imgArray, []byte(base64Encoding))
	}
	f.Images = &imgArray
	return f
}

func (f *Flow) ValidateFields() *Flow {
	f.ErrStatus = http.StatusBadRequest
	switch {
	case f.Err != nil:
		return f
	case !isValid(f.Pass.User.Email):
		f.Err = errors.New("неверный email адрес")
		return f
	case f.Pass.ID == "":
		f.Err = errors.New("отсутствует ID перевала")
		return f
	case f.Pass.Coords.Height == "":
		f.Err = errors.New("отсутствует координата: coords.Height")
		return f
	case f.Pass.Coords.Latitude == "":
		f.Err = errors.New("отсутствует координата: coords.Latitude")
		return f
	case f.Pass.Coords.Longitude == "":
		f.Err = errors.New("отсутствует координата: coords.Longitude")
		return f
	case f.Pass.User.Name+f.Pass.User.Fam+f.Pass.User.Otc == "":
		f.Err = errors.New("отсутствует имя пользователя")
		return f
	case f.Pass.BeautyTitle+f.Pass.Title+f.Pass.OtherTitles == "":
		f.Err = errors.New("отсутствует название объекта")
		return f
	case len(f.Pass.Images) == 0:
		//f.Err = errors.New("отсутствуют изображения")
		f.Warning = errors.New("warning: отсутствуют изображения")
		//return f
	case len(f.Pass.Images) != 0:
		for idx, elem := range f.Pass.Images {
			if elem.URL == "" {
				s := fmt.Sprintf("warning: отсутствует URL изображения: №%d, описание: %s", idx+1, elem.Title)
				//f.Err = errors.New(s)
				f.Warning = errors.New(s)
				f.ImgExpects = append(f.ImgExpects, idx)
				//return f
			}
		}
		//case f.Pass.User.ID == "":
		//	f.Err = errors.New("отсутствует ID пользователя")
		//	return f
	}
	return f

}

func isValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (f *Flow) ValidateData() *Flow {
	if f.Err != nil {
		return f
	}
	if f.Pass.AddTime == "" {
		f.Pass.AddTime = "NOW()"
	} else if t, err := time.Parse("2006-01-02 15:04:05", f.Pass.AddTime); err != nil {
		f.Err = fmt.Errorf("%w", err)
		return f
	} else {
		f.Pass.AddTime = t.Format("2006-01-02 15:04:05")
	}
	return f
}
