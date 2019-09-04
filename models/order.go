package models

import "time"

type Order struct {
	Id         int          `json:"id"`
	User       int          `json:"user"`
	Restaurant int          `json:"restaurant"`
	Time       time.Time    `json:"time"`
	Complete   bool         `json:"complete"`
	OrderItems []*OrderItem `json:"order_items"`
}
