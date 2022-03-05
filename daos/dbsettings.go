package daos

import (
	"database/sql"
	"errors"
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
