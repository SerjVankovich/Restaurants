package db

import (
	"../utils"
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func Connect() (*sql.DB, error) {
	path, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	dbConfig, err := utils.ParseDbConfig(path + "\\db\\dbconfig.json")

	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", "user="+dbConfig.User+" password="+dbConfig.Password+
		" dbname="+dbConfig.Dbname+" sslmode="+dbConfig.Sslmode)

	return db, err
}
