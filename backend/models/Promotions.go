package models

import "time"

// Promotion ánh xạ bảng Promotions trong cơ sở dữ liệu
type Promotion struct {
	PromotionID   int       `json:"promotion_id" db:"promotion_id"`
	Title         string    `json:"title" db:"title"`
	Description   string    `json:"description" db:"description"`
	DiscountRate  float64   `json:"discount_rate" db:"discount_rate"`
	DiscountValue float64   `json:"discount_value" db:"discount_value"`
	StartDate     time.Time `json:"start_date" db:"start_date"`
	EndDate       time.Time `json:"end_date" db:"end_date"`
	ApplicableTo  string    `json:"applicable_to" db:"applicable_to"`
	Status        string    `json:"status" db:"status"`
}
