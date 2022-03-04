package daos

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/serjbibox/FSTR/models"
)

const (
	singleMode = iota
	mailMode
	phoneMode
	fioMode
	statusMode
)

func (dao *PassDAO) Get(f *models.Flow) *models.Flow {
	var err error
	if f.Err != nil {
		return f
	}
	if err = DbConnect(); err != nil {
		f.Err = fmt.Errorf("%w", err)
		return f
	}
	defer DB.Close()
	var status string

	switch {
	case f.GetWith == statusMode:
		query := "SELECT status FROM pereval_added WHERE id = ($1);"
		f.Err = DB.QueryRow(query, f.ID).Scan(&status)
		if f.Err != nil {
			return f
		}
	case f.GetWith == singleMode:
		var pAdded string
		query := "SELECT status, raw_data FROM pereval_added WHERE id = ($1);"
		f.Err = DB.QueryRow(query, f.ID).Scan(&status, &pAdded)
		if f.Err != nil {
			return f
		}
		f.Err = json.Unmarshal([]byte(pAdded), &f.Pass)
		if f.Err != nil {
			return f
		}
	case f.GetWith == mailMode, f.GetWith == phoneMode:
		log.Println("mailMode ", f.GetBy)
		var rows *sql.Rows
		query := "SELECT raw_data FROM pereval_added WHERE raw_data::TEXT SIMILAR TO '(%" + f.GetBy + "%)'"
		rows, f.Err = DB.Query(query)
		if f.Err != nil {
			log.Println("ошибка в запросе")
			return f
		}
		defer rows.Close()
		for rows.Next() {
			var pAdded string
			f.Err = rows.Scan(&pAdded)
			if f.Err != nil {
				log.Println("ошибка в скане")
				return f
			}
			f.Err = json.Unmarshal([]byte(pAdded), &f.Pass)
			if f.Err != nil {
				log.Println("ошибка в marshall")
				return f
			}
			f.Parray = append(f.Parray, *f.Pass)
		}
	case f.GetWith == fioMode:

		var rows *sql.Rows
		fam := f.GetByFIO[0]
		name := f.GetByFIO[1]
		otc := f.GetByFIO[2]
		log.Println("fioMode ", fam, name, otc)
		query := "SELECT raw_data FROM pereval_added WHERE raw_data::TEXT SIMILAR TO '(%" +
			fam + "%)(%" + name + "%)(%" + otc + "%)'"
		rows, f.Err = DB.Query(query)
		if f.Err != nil {
			log.Println("ошибка в запросе")
			return f
		}
		defer rows.Close()
		for rows.Next() {
			var pAdded string
			f.Err = rows.Scan(&pAdded)
			if f.Err != nil {
				log.Println("ошибка в скане")
				return f
			}
			f.Err = json.Unmarshal([]byte(pAdded), &f.Pass)
			if f.Err != nil {
				log.Println("ошибка в marshall")
				return f
			}
			f.Parray = append(f.Parray, *f.Pass)
		}
	}

	f.Pass.Status = status
	return f
}
