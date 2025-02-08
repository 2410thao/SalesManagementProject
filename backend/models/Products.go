// package models

// import "time"

// type Product struct {
// 	ProductID   int       `json:"product_id" db:"product_id"`
// 	ProductName string    `json:"product_name" db:"product_name"`
// 	Price       float64   `json:"price" db:"price"`
// 	Cost        float64   `json:"cost" db:"cost"`
// 	Quantity    int       `json:"quantity" db:"quantity"`
// 	Unit        string    `json:"unit" db:"unit"`
// 	CreatedAt   time.Time `json:"created_at" db:"created_at"`
// 	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
// 	Image       string    `json:"image" db:"image"`
// 	CategoryID  *int      `json:"category_id" db:"category_id"` // Nullable foreign key
// }

package models

type Product struct {
	ProductID         int     `json:"product_id" db:"product_id"`
	ProductName       string  `json:"product_name"`
	Price             float64 `json:"price"`
	Cost              float64 `json:"cost"`
	Quantity          int     `json:"quantity"`
	Unit              string  `json:"unit"`
	Image             string  `json:"image"`
	CategoryID        *int64  `json:"category_id"` // Dùng con trỏ để xử lý NULL
	ManufacturingDate string  `json:"manufacturing_date"`
	ExpiryDate        string  `json:"expiry_date"`
}
