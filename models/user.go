package models

type Token string
type PrefRest string

type User struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	PrefRest  *PrefRest `json:"pref_rest"`
	Token     *Token    `json:"token"`
	Salt      string    `json:"salt"`
	Confirmed bool      `json:"confirmed"`
}

type Confirm struct {
	IsOk        bool   `json:"isOk"`
	AccessToken string `json:"access_token"`
	ConfirmHash string `json:"confirm_hash"`
}
