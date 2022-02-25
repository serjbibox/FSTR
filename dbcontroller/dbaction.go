package dbcontroller

import (
	"encoding/json"
	"log"
	"time"

	"github.com/serjbibox/FSTR/jsoncontroller"
)

const (
	new = "new"

	pending  = "pending"
	resolved = "resolved"
	accepted = "accepted"
	rejected = "rejected"
)

func AddData(p *jsoncontroller.Pereval) (id string, err error) {
	var t time.Time
	Connect()
	defer DB.Close()
	if t, err = time.Parse("2006-01-02 15:04:05", p.AddTime); err != nil {
		log.Println(err)
		return "", err
	}
	data, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
		return "", err
	}
	status := new
	err = DB.QueryRow("INSERT INTO pereval_added (date_added, raw_data, status) VALUES ($1, $2, $3) RETURNING ID;",
		t, data, status).Scan(&id)
	//result, err := db.Exec("INSERT INTO pereval_added (date_added, raw_data, status) VALUES ($1, $2, $3);", t, data, Status)
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(id)
	return id, nil
}
