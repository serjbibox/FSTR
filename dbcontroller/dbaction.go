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

func AddData(p *models.Pereval) (id string, err error) {
	//var t time.Time
	err = DbConnect()
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	defer DB.Close()
	pData, err := json.Marshal(p)
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
