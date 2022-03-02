package dbcontroller

import (
	"encoding/json"
	"fmt"

	"github.com/serjbibox/FSTR/models"
)

//Возвращает JSON структуру карточки перевала из поля raw_data таблицы pereval_added
func GetRow(id string) (p *models.Pereval, status string, err error) {
	err = DbConnect()
	if err != nil {
		return nil, "", fmt.Errorf("%w", err)
	}
	defer DB.Close()
	var pAdded string
	query := "SELECT status, raw_data FROM pereval_added WHERE id = ($1);"
	if err = DB.QueryRow(query, id).Scan(&status, &pAdded); err != nil {
		return nil, "", fmt.Errorf("%w", err)
	}
	if err = json.Unmarshal([]byte(pAdded), &p); err != nil {
		return nil, "", fmt.Errorf("%w", err)
	}
	return p, status, nil
}
