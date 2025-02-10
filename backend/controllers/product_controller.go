package controllers

import (
	"database/sql"
	"demo2501/backend/config"
	"demo2501/backend/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kataras/iris/v12"
)

// API: Thêm sản phẩm mới
func AddProduct(ctx iris.Context) {
	db := config.GetDB()
	var product models.Product

	if err := ctx.ReadJSON(&product); err != nil {
		log.Println("Error reading JSON:", err)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid input"})
		return
	}

	query := `INSERT INTO products (product_name, price, cost, quantity, unit, image, category_id, ManufacturingDate, ExpiryDate) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(query, product.ProductName, product.Price, product.Cost, product.Quantity, product.Unit, product.Image, product.CategoryID, product.ManufacturingDate, product.ExpiryDate)
	if err != nil {
		log.Printf("Lỗi khi thêm sản phẩm vào DB: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": fmt.Sprintf("Database error: %v", err)})
		return
	}

	ctx.JSON(iris.Map{"message": "Thêm thành công"})
}

// API: Lấy danh sách sản phẩm
func GetProducts(ctx iris.Context) {
	db := config.GetDB()
	var products []models.Product
	rows, err := db.Query("SELECT product_id, product_name, price, cost, quantity, unit, image, category_id, ManufacturingDate, ExpiryDate FROM products")
	if err != nil {
		log.Println("Database query error:", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Database query failed"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		var categoryID sql.NullInt64

		if err := rows.Scan(&product.ProductID, &product.ProductName, &product.Price, &product.Cost, &product.Quantity, &product.Unit, &product.Image, &categoryID, &product.ManufacturingDate, &product.ExpiryDate); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}

		// Xử lý category_id NULL
		if categoryID.Valid {
			product.CategoryID = &categoryID.Int64
		} else {
			product.CategoryID = nil
		}

		products = append(products, product)
	}

	ctx.JSON(iris.Map{"products": products})
}

// API: cập nhật sản phẩm
func UpdateProduct(ctx iris.Context) {
	db := config.GetDB()
	id := ctx.Params().GetIntDefault("id", 0)
	if id == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	if err := ctx.ReadJSON(&product); err != nil {
		log.Println("Error reading JSON:", err)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid input"})
		return
	}

	query := `UPDATE products SET product_name=?, price=?, cost=?, quantity=?, unit=?, image=?, category_id=?, ManufacturingDate=?, ExpiryDate=? WHERE product_id=?`
	_, err := db.Exec(query, product.ProductName, product.Price, product.Cost, product.Quantity, product.Unit, product.Image, product.CategoryID, product.ManufacturingDate, product.ExpiryDate, id)
	if err != nil {
		log.Printf("Lỗi khi cập nhật sản phẩm: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Database update failed"})
		return
	}

	ctx.JSON(iris.Map{"message": "Cập nhật sản phẩm thành công"})
}

// API: Xóa sản phẩm
func DeleteProduct(ctx iris.Context) {
	db := config.GetDB()
	id := ctx.Params().GetIntDefault("id", 0)
	if id == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid product ID"})
		return
	}

	query := `DELETE FROM products WHERE product_id=?`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Lỗi khi xóa sản phẩm: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Database delete failed"})
		return
	}

	ctx.JSON(iris.Map{"message": "Xóa sản phẩm thành công"})
}
