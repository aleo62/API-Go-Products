package model

type Product struct {
	ID     int     `json:"id"`
	Name   string  `json:"product_name"`
	Price  float64 `json:"price"`
}

