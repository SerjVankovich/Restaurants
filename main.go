package main

import (
	"./db"
	"./routes"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
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
