package entity

import "time"

type Order struct {
	ID           uint
	BuyerEmail   string
	BuyerAddress string
	OrderDate    time.Time
}
