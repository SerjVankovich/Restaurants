package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	PrefRest string `json:"pref_rest"`
	Token    string `json:"token"`
}
