package controllers

import "net/http"

func MockUsersController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("users controller"))
}
