package models

import "time"

// PurchaseDetail ánh xạ bảng PurchaseDetails trong cơ sở dữ liệu
type PurchaseDetail struct {
	PurchaseDetailID int       `json:"purchase_detail_id" db:"purchase_detail_id"`
	ProductID        int       `json:"product_id" db:"product_id"`
	PurchaseID       int       `json:"purchase_id" db:"purchase_id"`
	Quantity         int       `json:"quantity" db:"quantity"`
	Cost             float64   `json:"cost" db:"cost"`
	Subtotal         float64   `json:"subtotal" db:"subtotal"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}
