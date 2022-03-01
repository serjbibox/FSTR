package dbcontroller

import (
	"encoding/json"
	"fmt"
	"log"

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
	log.Println(status, p.Title, p.AddTime)
	return p, nil
}

func AddData(p *models.Pereval) (id string, err error) {
	//var t time.Time
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
	iData, err := json.Marshal(p.Images)
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
