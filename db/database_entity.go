package db

import (
	"database/sql"
	"strconv"
)

type DB struct {
	DataBase *sql.DB
}

func (db DB) GetAll(table string, conditions map[string]interface{}, tp interface{}) (interface{}, error) {
	if db.DataBase == nil {
		return nil, dbErr
	}

	query := `SELECT * FROM ` + table + ` WHERE `
	i := 1
	var args []interface{}
	for key, value := range conditions {
		query += key + ` = $` + strconv.Itoa(i) + `, `
		args = append(args, value)
		i += 1
	}

	rows, err := db.DataBase.Query(query, args...)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan()
	}

	return nil, nil

}
