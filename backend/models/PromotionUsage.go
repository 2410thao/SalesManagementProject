package models

import "time"

// PromotionUsage ánh xạ bảng PromotionUsage trong cơ sở dữ liệu
type PromotionUsage struct {
	UsageID        int       `json:"usage_id" db:"usage_id"`
	CustomerID     int       `json:"customer_id" db:"customer_id"`
	PromotionID    int       `json:"promotion_id" db:"promotion_id"`
	OrderID        int       `json:"order_id" db:"order_id"`
	DiscountAmount float64   `json:"discount_amount" db:"discount_amount"`
	AppliedAt      time.Time `json:"applied_at" db:"applied_at"`
}
