package models

// User đại diện cho bảng users trong cơ sở dữ liệu
type User struct {
	ID       int    `json:"id"`       // Trường ID, kiểu INT trong MySQL
	Username string `json:"username"` // Trường username, kiểu VARCHAR trong MySQL
	Password string `json:"password"` // Trường password, kiểu VARCHAR trong MySQL
}
