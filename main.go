package main

import (
	"./db"
	"./routes"
	"database/sql"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

var (
	dataBase *sql.DB
)

func main() {

	dataBase, err := db.Connect()
	defer dataBase.Close()

	var apiRoute = routes.ApiRoute(dataBase)

	err = http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, apiRoute))

	if err != nil {
		panic(err)
	}

}
