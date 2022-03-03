package apis

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/serjbibox/FSTR/models"
	"github.com/serjbibox/FSTR/services"
)

func SendHttp(w http.ResponseWriter, v ResponseInterface) {
	v.Send(w)
}

func SendErr(w http.ResponseWriter, statusCode int, err error) {
	er := ErrResponse{
		HTTPStatusCode: statusCode,
		ErrorText:      fmt.Sprint(err),
	}
	er.Send(w)
}

func GetImage(p *models.Pass) ([][]byte, error) {
	var i [][]byte
	for _, elem := range p.Images {
		response, err := http.Get(elem.URL)

		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			return nil, errors.New("невозможно загрузить изображение по адресу: " + elem.URL)
		}
		if img, err := ioutil.ReadAll(response.Body); err != nil {
			return nil, err
		} else {
			i = append(i, img)
		}

	}
	return i, nil
}

func imgData(m map[string]string) (*map[string][]int, error) {
	imgMap := make(map[string][]int)
	var err error
	for key, title := range m {
		imgId, err := strconv.Atoi(key)
		if err != nil {
			return nil, err
		}
		imgMap[title] = append(imgMap[title], imgId)
	}
	return &imgMap, err
}

func Validate(p *models.Pass, s *services.PassService) error {
	err := s.ValidateFields(p)
	if err != nil {
		return err
	}
	err = s.ValidateData(p)
	if err != nil {
		return err
	}
	return nil
}
