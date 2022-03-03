package daos

import (
	"database/sql"
	"log"
)

func DbConnect() error {
	var err error
	if DbConnString == "" {
		if DbConnString, err = ConnStringConfig(); err != nil {
			return err
		}
	}
	DB, err = sql.Open("postgres", DbConnString)
	if err != nil {
		return err
	} else {
		log.Println("Установлено соединение с БД")
	}
	return nil
}
