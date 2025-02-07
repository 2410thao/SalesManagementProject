package models

import "time"

// TieuThuong đại diện cho bảng TieuThuong trong cơ sở dữ liệu
type TieuThuong struct {
	TieuThuongID int       `json:"tieu_thuong_id" gorm:"primaryKey;autoIncrement"`
	TenCuaHang   string    `json:"ten_cua_hang" gorm:"type:varchar(255);not null"`
	SoDienThoai  string    `json:"so_dien_thoai" gorm:"type:varchar(15)"`
	Email        string    `json:"email" gorm:"type:varchar(255)"`
	DiaChi       string    `json:"dia_chi" gorm:"type:varchar(255)"`
	TinhID       int       `json:"tinh_id" gorm:"index"`   // Khóa ngoại
	TPID         int       `json:"tp_id" gorm:"index"`     // Khóa ngoại
	PhuongID     int       `json:"phuong_id" gorm:"index"` // Khóa ngoại
	QuanID       int       `json:"quan_id" gorm:"index"`   // Khóa ngoại
	MatKhau      string    `json:"mat_khau" gorm:"type:varchar(255);not null"`
	NgayTao      time.Time `json:"ngay_tao" gorm:"autoCreateTime"`
	NgayCapNhat  time.Time `json:"ngay_cap_nhat" gorm:"autoUpdateTime"`
}
