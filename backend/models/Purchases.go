package models

import "time"

type Purchase struct {
	PurchaseID int       `json:"purchase_id" db:"purchase_id"`
	SupplierID int       `json:"supplier_id" db:"supplier_id"`
	Date       time.Time `json:"date" db:"date"`
	TotalCost  float64   `json:"total_cost" db:"total_cost"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
