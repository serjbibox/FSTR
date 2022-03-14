package daos

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
	FSTR_DB_NAME  = "dci84nq5jb8beh"
)

var DB *sql.DB
var DbConnString string

func ConnStringConfig() (string, error) {
	m := map[string]string{
		FSTR_DB_LOGIN: "",
		FSTR_DB_PASS:  "",
		FSTR_DB_HOST:  "",
		FSTR_DB_PORT:  "",
	}
	for key := range m {
		if d := readEnvironment(key); d == "" {
			return "", errors.New("Системная переменная не найдена: " + key)
		} else {
			m[key] = d
		}
	}
	connString := "postgres://" +
		m[FSTR_DB_LOGIN] +
		":" +
		m[FSTR_DB_PASS] +
		"@" +
		m[FSTR_DB_HOST] +
		":" +
		m[FSTR_DB_PORT] +
		"/" +
		FSTR_DB_NAME // + "?sslmode=disable"

	log.Println(connString)

	return connString, nil

}

func readEnvironment(key string) string {
	if env, ok := os.LookupEnv(key); !ok {
		return ""
	} else {
		return env
	}
}
