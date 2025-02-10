package controllers

import (
	"demo2501/backend/config"
	"demo2501/backend/models"
	"log"

	"github.com/kataras/iris/v12"
)

// API: Lấy danh sách danh mục
func GetCategories(ctx iris.Context) {
	db := config.GetDB()
	rows, err := db.Query("SELECT category_id, name, description FROM categories")
	if db == nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": "Database connection is not initialized"})
		return
	}

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to fetch categories"})
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.CategoryID, &c.Name, &c.Description); err != nil {
			log.Println("Scan error:", err)
			continue
		}
		categories = append(categories, c)
	}
	ctx.JSON(iris.Map{"categories": categories})
}

// API: Thêm danh mục mới
func AddCategory(ctx iris.Context) {
	var category models.Category
	db := config.GetDB()
	// Đọc dữ liệu từ JSON
	if err := ctx.ReadJSON(&category); err != nil {
		log.Println("Lỗi đọc JSON:", err)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Dữ liệu không hợp lệ"})
		return
	}

	// Kiểm tra dữ liệu đầu vào
	if category.Name == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Tên danh mục không được để trống"})
		return
	}

	// Câu lệnh SQL để thêm danh mục
	query := `INSERT INTO categories (name, description) VALUES (?, ?)`
	result, err := db.Exec(query, category.Name, category.Description)
	if err != nil {
		log.Printf("Lỗi khi thêm danh mục vào DB: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Lỗi khi thêm danh mục"})
		return
	}

	// Lấy ID của danh mục vừa thêm
	categoryID, _ := result.LastInsertId()
	category.CategoryID = int(categoryID)

	// Trả về phản hồi JSON
	ctx.JSON(iris.Map{
		"message":  "Thêm danh mục thành công",
		"category": category,
	})
}

// API: Cập nhật danh mục
func UpdateCategory(ctx iris.Context) {
	db := config.GetDB()
	// Lấy ID từ URL
	id := ctx.Params().GetIntDefault("id", 0)
	if id == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "ID danh mục không hợp lệ"})
		return
	}

	var category models.Category

	// Đọc dữ liệu JSON từ request
	if err := ctx.ReadJSON(&category); err != nil {
		log.Println("Lỗi đọc JSON:", err)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Dữ liệu không hợp lệ"})
		return
	}

	// Kiểm tra đầu vào
	if category.Name == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Tên danh mục không được để trống"})
		return
	}

	// Cập nhật danh mục trong cơ sở dữ liệu
	query := `UPDATE categories SET name=?, description=? WHERE category_id=?`
	_, err := db.Exec(query, category.Name, category.Description, id)
	if err != nil {
		log.Printf("Lỗi khi cập nhật danh mục: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Không thể cập nhật danh mục"})
		return
	}

	ctx.JSON(iris.Map{"message": "Cập nhật danh mục thành công"})
}

// API: Xóa danh mục
func DeleteCategory(ctx iris.Context) {
	db := config.GetDB()
	// Lấy ID từ URL
	id := ctx.Params().GetIntDefault("id", 0)
	if id == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "ID danh mục không hợp lệ"})
		return
	}

	// Kiểm tra xem danh mục có sản phẩm nào không (Ràng buộc khóa ngoại)
	var count int
	checkQuery := `SELECT COUNT(*) FROM products WHERE category_id=?`
	err := db.QueryRow(checkQuery, id).Scan(&count)
	if err != nil {
		log.Printf("Lỗi khi kiểm tra danh mục: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Không thể kiểm tra danh mục"})
		return
	}

	// Nếu danh mục có sản phẩm thì không cho phép xóa
	if count > 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Không thể xóa danh mục vì vẫn còn sản phẩm thuộc danh mục này"})
		return
	}

	// Xóa danh mục trong cơ sở dữ liệu
	deleteQuery := `DELETE FROM categories WHERE category_id=?`
	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		log.Printf("Lỗi khi xóa danh mục: %v", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Không thể xóa danh mục"})
		return
	}

	ctx.JSON(iris.Map{"message": "Xóa danh mục thành công"})
}
