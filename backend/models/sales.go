package models

import "time"

type Sale struct {
	SaleID      int       `json:"sale_id" db:"sale_id"`
	CustomerID  int       `json:"customer_id" db:"customer_id"`
	Date        time.Time `json:"date" db:"date"`
	TotalAmount float64   `json:"total_amount" db:"total_amount"`
	TotalCost   float64   `json:"total_cost" db:"total_cost"`
	Profit      float64   `json:"profit" db:"profit"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	Status      string    `json:"status" db:"status"`
}
