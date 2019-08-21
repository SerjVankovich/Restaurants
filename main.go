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

	_, err := db.Connect()

	var apiRoute = routes.ApiRoute()

	err = http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, apiRoute))

	if err != nil {
		panic(err)
	}

}
