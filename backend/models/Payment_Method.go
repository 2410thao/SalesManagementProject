package models

import "time"

type PaymentMethod struct {
	PaymentMethodID int       `json:"payment_method_id" db:"payment_method_id"`
	Name            string    `json:"name" db:"name"`
	Description     *string   `json:"description" db:"description"` // Có thể NULL
	IsActive        *bool     `json:"is_active" db:"is_active"`     // Có thể NULL, mặc định TRUE (1)
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}
