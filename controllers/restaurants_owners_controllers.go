package controllers

import "net/http"

func MockRestaurantsOwnerController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("restaurants owner controller"))
}
