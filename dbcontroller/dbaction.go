package dbcontroller

import (
	"encoding/json"
	"fmt"

	"github.com/serjbibox/FSTR/models"
)

const (
	new      = "new"
	pending  = "pending"
	resolved = "resolved"
	accepted = "accepted"
	rejected = "rejected"
)

func GetRow(id string) (p *models.Pereval, err error) {
	err = DbConnect()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer DB.Close()
	var status, pAdded, foo string
	query := "SELECT * FROM pereval_added WHERE id = ($1);"
	if err = DB.QueryRow(query, id).Scan(&foo, &foo, &pAdded, &foo, &status); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	if err = json.Unmarshal([]byte(pAdded), &p); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return p, nil
}

func AddData(p *models.Pereval, ai *map[string][]int) (id string, err error) {
	err = DbConnect()
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	defer DB.Close()
	pa := models.NewPerevalAdded(p)
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

func AddImage(img [][]byte, p *models.Pereval) (m map[string]string, err error) {
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
