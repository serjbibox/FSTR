package dbcontroller

import (
	"database/sql"
	"errors"
	"log"
	"os"
)

const DbName = "FSTR"

var DB *sql.DB
var DbConnString string

func DbConnect() {
	if DbConnString == "" {
		DbConnString = ConnStringConfig()
	}
	db, err := sql.Open("postgres", DbConnString)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("DataBase connection OK")
	}
	DB = db
}

func ConnStringConfig() string {
	//os.Setenv("FSTR_DB_HOST", "@127.0.0.1:5432/")
	//os.Setenv("FSTR_DB_HOST", "@35.239.250.100:5432/")
	//os.Setenv("FSTR_DB_PORT", "5432")
	//os.Setenv("FSTR_DB_LOGIN", "fstr")
	//os.Setenv("FSTR_DB_PASS", "123456")
	return "postgres://" +
		readEnvironment("FSTR_DB_LOGIN") +
		":" +
		readEnvironment("FSTR_DB_PASS") +
		"@" +
		readEnvironment("FSTR_DB_HOST") +
		":" +
		readEnvironment("FSTR_DB_PORT") +
		"/" +
		DbName + "?sslmode=disable"
}
func readEnvironment(key string) string {
	if env, ok := os.LookupEnv(key); !ok {
		log.Println(errors.New("no such environment"))
		return ""

	} else {
		return env
	}
}
