package main

import (
	
	"demo2501/backend/config"
	"demo2501/backend/controllers"
	
	_ "github.com/go-sql-driver/mysql"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func main() {

	// Khởi tạo kết nối DB
	config.InitDB()

	app := iris.New()

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

	app.Get("/api/products", controllers.GetProducts)
	app.Post("/api/products", controllers.AddProduct)
	app.Put("/api/products/{id:int}", controllers.UpdateProduct)
	app.Delete("/api/products/{id:int}", controllers.DeleteProduct)

	app.Get("/api/categories", controllers.GetCategories)
	app.Post("/api/categories", controllers.AddCategory)
	app.Put("/api/categories/{id:int}", controllers.UpdateCategory)
	app.Delete("/api/categories/{id:int}", controllers.DeleteCategory)

	app.Listen(":8080")
}
