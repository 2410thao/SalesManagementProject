package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB là hàm để kết nối tới cơ sở dữ liệu
func ConnectDB() *sql.DB {
	dsn := "root:1234@tcp(127.0.0.1:3306)/hieu" // Sửa username, password và tên database nếu cần
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Lỗi khi kết nối MySQL: ", err)
	}

	// Kiểm tra kết nối
	if err := db.Ping(); err != nil {
		log.Fatal("Lỗi khi kiểm tra kết nối: ", err)
	}

	fmt.Println("Kết nối đến MySQL thành công!")
	return db
}
