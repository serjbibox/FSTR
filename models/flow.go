package models

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Flow struct {
	Pass      *Pass
	Parray    []Pass
	ID        string
	RID       string
	Images    *[][]byte
	ImgMapPA  *map[string]string
	ImgMapIA  *map[string][]int
	Err       error
	ErrStatus int
	GetWith   byte
	GetBy     string
	GetByFIO  [3]string
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
	for _, elem := range f.Pass.Images {
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
		if img, err := ioutil.ReadAll(response.Body); err != nil {
			f.Err = fmt.Errorf("%w", err)
			return f
		} else {
			imgArray = append(imgArray, img)
		}
	}
	f.Images = &imgArray
	return f
}

func (f *Flow) ValidateFields() *Flow {
	f.ErrStatus = http.StatusBadRequest
	switch {
	case f.Err != nil:
		return f
	case f.Pass.ID == "":
		f.Err = errors.New("отсутствует ID перевала")
		return f
	case f.Pass.User.ID == "":
		f.Err = errors.New("отсутствует ID пользователя")
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
		f.Err = errors.New("отсутствуют изображения")
		return f
	case len(f.Pass.Images) != 0:
		for idx, elem := range f.Pass.Images {
			if elem.URL == "" {
				s := fmt.Sprintf("отсутствует URL изображения: №%d, описание: %s", idx+1, elem.Title)
				f.Err = errors.New(s)
				return f
			}
		}
	}
	return f
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
