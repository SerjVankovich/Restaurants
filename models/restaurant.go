package models

type Restaurant struct {
	Id          int32   `json:"id"`
	Name        string  `json:"name"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	Description string  `json:"description"`
	Owner       int32   `json:"owner"`
}
