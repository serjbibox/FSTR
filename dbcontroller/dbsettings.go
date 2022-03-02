package dbcontroller

import (
	"database/sql"
	"errors"
	"log"
	"os"
)

const (
	new      = "new"
	pending  = "pending"
	resolved = "resolved"
	accepted = "accepted"
	rejected = "rejected"
)

const (
	FSTR_DB_LOGIN = "FSTR_DB_LOGIN"
	FSTR_DB_PASS  = "FSTR_DB_PASS"
	FSTR_DB_HOST  = "FSTR_DB_HOST"
	FSTR_DB_PORT  = "FSTR_DB_PORT"
	FSTR_DB_NAME  = "FSTR"
)

var DB *sql.DB
var DbConnString string

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

func ConnStringConfig() (string, error) {
	m := map[string]string{
		FSTR_DB_LOGIN: "",
		FSTR_DB_PASS:  "",
		FSTR_DB_HOST:  "",
		FSTR_DB_PORT:  "",
	}
	// Установка системных переменных для тестирования
	// После теста удалить!
	os.Setenv(FSTR_DB_LOGIN, "fstr")
	os.Setenv(FSTR_DB_PASS, "123456")
	os.Setenv(FSTR_DB_HOST, "35.239.250.100")
	os.Setenv(FSTR_DB_PORT, "5432")
	//////////////////////////////////////////////////
	for key := range m {
		if d := readEnvironment(key); d == "" {
			return "", errors.New("Системная переменная не найдена: " + key)
		} else {
			m[key] = d
		}
	}

	return "postgres://" +
		m[FSTR_DB_LOGIN] +
		":" +
		m[FSTR_DB_PASS] +
		"@" +
		m[FSTR_DB_HOST] +
		":" +
		m[FSTR_DB_PORT] +
		"/" +
		FSTR_DB_NAME + "?sslmode=disable", nil
}

func readEnvironment(key string) string {
	if env, ok := os.LookupEnv(key); !ok {
		return ""
	} else {
		return env
	}
}
