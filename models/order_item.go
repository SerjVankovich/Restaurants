package models

type OrderItem struct {
	Id         int `json:"id"`
	Product    int `json:"product"`
	NumProduct int `json:"num_product"`
	Order      int `json:"order"`
}
