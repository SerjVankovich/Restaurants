package controllers

import "net/http"

func MockRestaurantsController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("restaurants controller"))
}
