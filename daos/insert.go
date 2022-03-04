package daos

import (
	"encoding/json"
	"fmt"

	"github.com/serjbibox/FSTR/models"
)

func (dao *PassDAO) InsertTo(f *models.Flow, table string) *models.Flow {
	if f.Err != nil {
		return f
	}
	var err error
	if err = DbConnect(); err != nil {
		f.Err = fmt.Errorf("%w", err)
		return f
	}
	defer DB.Close()

	switch table {
	case "pereval_added":
		pa := NewPassAdded(f.Pass)
		pData, err := json.Marshal(pa)
		if err != nil {
			f.Err = fmt.Errorf("%w", err)
			return f
		}
		iData, err := json.Marshal(f.ImgMapIA)
		if err != nil {
			f.Err = fmt.Errorf("%w", err)
			return f
		}
		status := new
		query := ""
		if f.RID == "" {
			query = "INSERT INTO pereval_added (date_added, raw_data, images, status) VALUES ($1, $2, $3, $4) RETURNING ID;"
		} else {
			query = "UPDATE pereval_added SET date_added = $1, raw_data = $2, images = $3, status = $4" +
				" WHERE id = " + f.RID + " RETURNING ID;"
		}
		err = DB.QueryRow(query, f.Pass.AddTime, pData, iData, status).Scan(&f.ID)
		if err != nil {
			f.Err = fmt.Errorf("%w", err)
			return f
		}
	case "pereval_images":
		var id []string
		m := make(map[string]string)
		for idx, elem := range *f.Images {
			id = append(id, "")
			err = DB.QueryRow("INSERT INTO pereval_images (date_added, img) VALUES ($1, $2) RETURNING ID;",
				f.Pass.AddTime, elem).Scan(&id[idx])
			if err != nil {
				f.Err = fmt.Errorf("%w", err)
				return f
			}
			m[id[idx]] = f.Pass.Images[idx].Title
		}
		f.ImgMapPA = &m
	}

	return f
}
