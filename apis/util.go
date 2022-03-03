package apis

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/serjbibox/FSTR/models"
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
