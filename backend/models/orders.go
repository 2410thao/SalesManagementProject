package models

import "time"

type Order struct {
	OrderID       int        `json:"order_id" db:"order_id"`
	TotalAmount   float64    `json:"total_amount" db:"total_amount"`
	Status        *string    `json:"status" db:"status"`                 // Có thể NULL
	OrderDate     *time.Time `json:"order_date" db:"order_date"`         // Có thể NULL
	DeliveryDate  *time.Time `json:"delivery_date" db:"delivery_date"`   // Có thể NULL
	Address       *string    `json:"address" db:"address"`               // Có thể NULL
	Voucher       *string    `json:"voucher" db:"voucher"`               // Có thể NULL
	VoucherAmount *float64   `json:"voucher_amount" db:"voucher_amount"` // Có thể NULL
	Note          *string    `json:"note" db:"note"`                     // Có thể NULL
	CustomerID    int        `json:"customer_id" db:"customer_id"`
}
