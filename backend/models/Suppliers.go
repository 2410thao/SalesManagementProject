package models

import "time"

type Supplier struct {
	SupplierID int       `json:"supplier_id" db:"supplier_id"`
	Name       string    `json:"name" db:"name"`
	Phone      *int      `json:"phone" db:"phone"`     // Có thể NULL
	Address    *string   `json:"address" db:"address"` // Có thể NULL
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
