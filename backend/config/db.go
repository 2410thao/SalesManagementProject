package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import driver MySQL
)

var DB *sql.DB // Biến toàn cục lưu kết nối DB

// InitDB khởi tạo kết nối tới cơ sở dữ liệu MySQL
func InitDB() {
	var err error
	dsn := "root:1234@tcp(127.0.0.1:3306)/hieu" // Thay đổi username, password, database nếu cần

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ Lỗi khi kết nối MySQL: ", err)
	}

	// Kiểm tra kết nối
	if err := DB.Ping(); err != nil {
		log.Fatal("❌ Lỗi khi kiểm tra kết nối: ", err)
	}

	fmt.Println("✅ Kết nối đến MySQL thành công!")
}

// GetDB trả về kết nối DB để sử dụng ở các nơi khác
func GetDB() *sql.DB {
	return DB
}
