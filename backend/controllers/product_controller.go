package controllers

import (
	"database/sql"

	"log"
	"time"

	"github.com/kataras/iris/v12"
)

var db *sql.DB

func initDB() {
	var err error
	connStr := "user=postgres password=yourpassword dbname=yourdb sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
}

// Product model
type Product struct {
	ID                int        `json:"id"`
	ProductName       string     `json:"productName"`
	Price             float64    `json:"price"`
	Cost              float64    `json:"cost"`
	Quantity          int        `json:"quantity"`
	Unit              string     `json:"unit"`
	CategoryID        int        `json:"categoryId"` // Thêm trường category_id
	CreatedAt         time.Time  `json:"createdAt"`
	UpdatedAt         time.Time  `json:"updatedAt"`
	ImageURL          string     `json:"imageURL"`
	ManufacturingDate *time.Time `json:"manufacturingDate,omitempty"`
	ExpiryDate        *time.Time `json:"expiryDate,omitempty"`
}

// Create product
// Create product
// Create product
func AddProductHandler(ctx iris.Context) {
	var p Product
	if err := ctx.ReadJSON(&p); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{"error": "Invalid JSON data"})
		return
	}

	// Validate required fields
	if p.ProductName == "" || p.Price <= 0 || p.Quantity < 0 || p.Unit == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{"error": "Missing or invalid fields"})
		return
	}

	// Parse dates
	if p.ManufacturingDate != nil && p.ExpiryDate != nil && p.ExpiryDate.Before(*p.ManufacturingDate) {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{"error": "ExpiryDate cannot be earlier than ManufacturingDate"})
		return
	}

	// Insert into database
	query := `INSERT INTO products (product_name, price, cost, quantity, unit, category_id, created_at, updated_at, image_url, manufacturing_date, expiry_date)
VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW(), $7, $8, $9) RETURNING id`
	row := db.QueryRow(query, p.ProductName, p.Price, p.Cost, p.Quantity, p.Unit, p.CategoryID, p.ImageURL, p.ManufacturingDate, p.ExpiryDate)

	// Log error in case of failure
	if err := row.Scan(&p.ID); err != nil {
		// Log chi tiết lỗi và thêm thông tin về query
		log.Printf("Error inserting product into DB: %v\nQuery: %s\n", err, query)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": "Failed to insert product"})
		return
	}

	// Return the newly created product
	ctx.JSON(p)
}

// Get all products
func GetProductsHandler(ctx iris.Context) {
	rows, err := db.Query("SELECT id, product_name, price, cost, quantity, unit, created_at, updated_at, image_url, manufacturing_date, expiry_date FROM products")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": "Failed to fetch products"})
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.ProductName, &p.Price, &p.Cost, &p.Quantity, &p.Unit, &p.CreatedAt, &p.UpdatedAt, &p.ImageURL, &p.ManufacturingDate, &p.ExpiryDate); err != nil {
			log.Println("Scan error:", err)
			continue
		}
		products = append(products, p)
	}
	ctx.JSON(products)
}

// Update product
func UpdateProductHandler(ctx iris.Context) {
	id := ctx.Params().GetIntDefault("id", 0)
	if id == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{"error": "Invalid product ID"})
		return
	}

	var p Product
	if err := ctx.ReadJSON(&p); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{"error": "Invalid JSON data"})
		return
	}

	query := `UPDATE products SET product_name=$1, price=$2, cost=$3, quantity=$4, unit=$5, updated_at=NOW(), image_url=$6, manufacturing_date=$7, expiry_date=$8 WHERE id=$9`
	_, err := db.Exec(query, p.ProductName, p.Price, p.Cost, p.Quantity, p.Unit, p.ImageURL, p.ManufacturingDate, p.ExpiryDate, id)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": "Failed to update product"})
		return
	}
	ctx.JSON(map[string]string{"message": "Product updated successfully"})
}
