package models

type Restaurant struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	Description float32 `json:"description"`
	Owner       int     `json:"owner"`
}
