package dbcontroller

import (
	"encoding/json"
	"fmt"
	"time"

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
	var t time.Time
	err = DbConnect()
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	defer DB.Close()
	if err := p.Validate(); err != nil {
		return "", fmt.Errorf("%w", err)
	}
	if t, err = time.Parse("2006-01-02 15:04:05", p.AddTime); err != nil {
		return "", fmt.Errorf("%w", err)
	}
	data, err := json.Marshal(p)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	status := new
	err = DB.QueryRow("INSERT INTO pereval_added (date_added, raw_data, status) VALUES ($1, $2, $3) RETURNING ID;",
		t, data, status).Scan(&id)
	//result, err := db.Exec("INSERT INTO pereval_added (date_added, raw_data, status) VALUES ($1, $2, $3);", t, data, Status)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	return id, nil
}
