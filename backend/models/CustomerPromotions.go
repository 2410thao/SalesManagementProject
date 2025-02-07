package models

import "time"

type CustomerPromotion struct {
	CustomerPromotionID int        `json:"customer_promotion_id" db:"customer_promotion_id"`
	CustomerID          int        `json:"customer_id" db:"customer_id"`
	PromotionID         int        `json:"promotion_id" db:"promotion_id"`
	UsageLimit          *int       `json:"usage_limit" db:"usage_limit"`         // Có thể NULL
	UsageCount          *int       `json:"usage_count" db:"usage_count"`         // Có thể NULL
	ExpirationDate      *time.Time `json:"expiration_date" db:"expiration_date"` // Có thể NULL
	Status              *string    `json:"status" db:"status"`                   // Có thể NULL
}
