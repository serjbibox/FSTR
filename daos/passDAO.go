package daos

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/serjbibox/FSTR/models"
)

type PassDAO struct{}

func NewPassDAO() *PassDAO {
	return &PassDAO{}
}

func (dao *PassDAO) Get(id string) (*models.Pass, error) {
	var p *models.Pass
	var err error
	if p, err = GetPass(id); err != nil {
		return nil, err
	}
	return p, err
}

//Возвращает JSON структуру карточки перевала из поля raw_data таблицы pereval_added
func GetPass(id string) (p *models.Pass, err error) {

	if err = DbConnect(); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer DB.Close()
	var pAdded, status string
	query := "SELECT status, raw_data FROM pereval_added WHERE id = ($1);"
	if err = DB.QueryRow(query, id).Scan(&status, &pAdded); err != nil {
		log.Println("ошибка тут")
		return nil, fmt.Errorf("%w", err)
	}
	if err = json.Unmarshal([]byte(pAdded), &p); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	p.Status = status
	return p, nil
}

func (dao *PassDAO) Create(r *http.Request) (*models.Pass, error) {
	p := models.Pass{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (dao *PassDAO) ValidateFields(p *models.Pass) error {
	switch {
	case p.ID == "":
		return errors.New("отсутствует ID перевала")
	case p.User.ID == "":
		return errors.New("отсутствует ID пользователя")
	case p.Coords.Height == "":
		return errors.New("отсутствует координата: coords.Height")
	case p.Coords.Latitude == "":
		return errors.New("отсутствует координата: coords.Latitude")
	case p.Coords.Longitude == "":
		return errors.New("отсутствует координата: coords.Longitude")
	case p.User.Name+p.User.Fam+p.User.Otc == "":
		return errors.New("отсутствует имя пользователя")
	case p.BeautyTitle+p.Title+p.OtherTitles == "":
		return errors.New("отсутствует название объекта")
	case len(p.Images) == 0:
		return errors.New("отсутствуют изображения")
	case len(p.Images) != 0:
		for idx, elem := range p.Images {
			if elem.URL == "" {
				s := fmt.Sprintf("отсутствует URL изображения: №%d, описание: %s", idx+1, elem.Title)
				return errors.New(s)
			}
		}
	}
	return nil
}

func (dao *PassDAO) ValidateData(p *models.Pass) error {
	if p.AddTime == "" {
		p.AddTime = "NOW()"
	} else if t, err := time.Parse("2006-01-02 15:04:05", p.AddTime); err != nil {
		return fmt.Errorf("%w", err)
	} else {
		p.AddTime = t.Format("2006-01-02 15:04:05")
	}
	return nil
}

func (dao *PassDAO) Insert(p *models.Pass, ai *map[string][]int) (id string, err error) {
	err = DbConnect()
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	defer DB.Close()
	pa := NewPassAdded(p)
	pData, err := json.Marshal(pa)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	iData, err := json.Marshal(ai)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	status := new
	err = DB.QueryRow("INSERT INTO pereval_added (date_added, raw_data, images, status) VALUES ($1, $2, $3, $4) RETURNING ID;",
		p.AddTime, pData, iData, status).Scan(&id)
	//result, err := db.Exec("INSERT INTO pereval_added (date_added, raw_data, status) VALUES ($1, $2, $3);", t, data, Status)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	return id, nil
}

func (dao *PassDAO) InsertImage(p *models.Pass, img [][]byte) (m map[string]string, err error) {
	var id []string
	err = DbConnect()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer DB.Close()
	m = make(map[string]string)
	for idx, elem := range img {
		id = append(id, "")
		err = DB.QueryRow("INSERT INTO pereval_images (date_added, img) VALUES ($1, $2) RETURNING ID;",
			p.AddTime, elem).Scan(&id[idx])
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		m[id[idx]] = p.Images[idx].Title
	}
	return m, nil
}

func NewPassAdded(p *models.Pass) models.PassAdded {
	return models.PassAdded{
		ID:          p.ID,
		BeautyTitle: p.BeautyTitle,
		Title:       p.Title,
		OtherTitles: p.OtherTitles,
		Connect:     p.Connect,
		AddTime:     p.AddTime,
		Coords:      p.Coords,
		Type:        p.Type,
		Level:       p.Level,
		User:        p.User,
	}
}
