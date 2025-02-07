package models

import "time"

type Cart struct {
	CartID    int       `json:"cart_id" db:"cart_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
	Price     float64   `json:"price" db:"price"`
	OrderID   *int      `json:"order_id" db:"order_id"`     // Có thể NULL
	ProductID *int      `json:"product_id" db:"product_id"` // Có thể NULL
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
