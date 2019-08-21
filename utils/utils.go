package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type DbConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Sslmode  string `json:"sslmode"`
}

func ParseDbConfig(path string) (*DbConfig, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	bytevalue, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	var dbConfig DbConfig

	err = json.Unmarshal(bytevalue, &dbConfig)

	if err != nil {
		return nil, err
	}

	return &dbConfig, nil
}
