package models

type SaleDetail struct {
	SaleDetailID int     `json:"sale_detail_id" db:"sale_detail_id"`
	SaleID       int     `json:"sale_id" db:"sale_id"`
	ProductID    int     `json:"product_id" db:"product_id"`
	Quantity     int     `json:"quantity" db:"quantity"`
	Price        float64 `json:"price" db:"price"`
	Subtotal     float64 `json:"subtotal" db:"subtotal"`
}
