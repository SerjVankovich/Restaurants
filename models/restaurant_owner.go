package models

type RestaurantOwner struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Token     *string `json:"token"`
	Salt      *string `json:"salt"`
	Confirmed bool    `json:"confirmed"`
}
