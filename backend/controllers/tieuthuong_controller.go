package controllers

import (
	"database/sql"
	"demo2501/backend/models"
	"log"

	"github.com/kataras/iris/v12"
)

// Hàm đăng ký tiểu thương
func RegisterTieuThuong(ctx iris.Context, db *sql.DB) {
	var tt models.TieuThuong
	err := ctx.ReadJSON(&tt)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Dữ liệu không hợp lệ"})
		return
	}

	// Thêm tiểu thương mới vào cơ sở dữ liệu
	query := `INSERT INTO TieuThuong 
		(Ten_Cua_Hang, So_Dien_Thoai, Email, Dia_Chi, Tinh_ID, TP_ID, Phuong_ID, Quan_ID) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = db.Exec(query, tt.TenCuaHang, tt.SoDienThoai, tt.Email, tt.DiaChi, tt.TinhID, tt.TPID, tt.PhuongID, tt.QuanID)
	if err != nil {
		log.Println("Database error:", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Lỗi thêm tiểu thương"})
		return
	}

	ctx.JSON(iris.Map{"message": "Đăng ký tiểu thương thành công"})
}

// Ví dụ thêm các controller khác
func GetTieuThuongList(ctx iris.Context, db *sql.DB) {
	rows, err := db.Query("SELECT * FROM TieuThuong")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Lỗi truy vấn dữ liệu"})
		return
	}
	defer rows.Close()

	var tieuthuongList []models.TieuThuong
	for rows.Next() {
		var tt models.TieuThuong
		err := rows.Scan(&tt.TieuThuongID, &tt.TenCuaHang, &tt.SoDienThoai, &tt.Email, &tt.DiaChi, &tt.TinhID, &tt.TPID, &tt.PhuongID, &tt.QuanID)
		if err != nil {
			log.Println("Scan error:", err)
			continue
		}
		tieuthuongList = append(tieuthuongList, tt)
	}

	ctx.JSON(tieuthuongList)
}
