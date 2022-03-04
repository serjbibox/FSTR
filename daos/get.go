package daos

import (
	"database/sql"
	"encoding/json"
	"fmt"

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
		var rows *sql.Rows
		query := "SELECT raw_data FROM pereval_added WHERE raw_data::TEXT SIMILAR TO '(%" + f.GetBy + "%)'"
		rows, f.Err = DB.Query(query)
		if f.Err != nil {
			return f
		}
		defer rows.Close()
		rowsScan(rows, f)
	case f.GetWith == fioMode:
		var rows *sql.Rows
		fam := f.GetByFIO[0]
		name := f.GetByFIO[1]
		otc := f.GetByFIO[2]
		query := "SELECT raw_data FROM pereval_added WHERE raw_data::TEXT SIMILAR TO '(%" +
			fam + "%)(%" + name + "%)(%" + otc + "%)'"
		rows, f.Err = DB.Query(query)
		if f.Err != nil {
			return f
		}
		defer rows.Close()
		rowsScan(rows, f)
	}

	f.Pass.Status = status
	return f
}

func rowsScan(rows *sql.Rows, f *models.Flow) {
	for rows.Next() {
		var pAdded string
		f.Err = rows.Scan(&pAdded)
		if f.Err != nil {
			return
		}
		f.Err = json.Unmarshal([]byte(pAdded), &f.Pass)
		if f.Err != nil {
			return
		}
		f.Parray = append(f.Parray, *f.Pass)
	}
}
