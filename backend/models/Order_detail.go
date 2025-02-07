package models

type OrderDetail struct {
	OrderDetailID int     `json:"order_detail_id" db:"order_detail_id"`
	OrderID       int     `json:"order_id" db:"order_id"`
	Quantity      int     `json:"quantity" db:"quantity"`
	Price         float64 `json:"price" db:"price"`
}
