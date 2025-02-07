package main

import (
	"database/sql"
	"demo2501/backend/config"
	"demo2501/backend/controllers"
	"demo2501/backend/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	// Kết nối cơ sở dữ liệu
	db := config.ConnectDB()
	defer db.Close()

	// Cấu hình CORS
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	app.UseRouter(crs)

	// Endpoint kiểm tra kết nối
	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})

	// Endpoint đăng ký & đăng nhập
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)

	// API Thêm sản phẩm
	app.Post("/api/products", func(ctx iris.Context) {
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
	})

	// API Lấy danh sách sản phẩm
	app.Get("/api/products", func(ctx iris.Context) {
		var products []models.Product
		rows, err := db.Query("SELECT product_name, price, cost, quantity, unit, image, category_id, ManufacturingDate, ExpiryDate FROM products")
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

			if err := rows.Scan(&product.ProductName, &product.Price, &product.Cost, &product.Quantity, &product.Unit, &product.Image, &categoryID, &product.ManufacturingDate, &product.ExpiryDate); err != nil {
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
	})

	app.Listen(":8080")
}
