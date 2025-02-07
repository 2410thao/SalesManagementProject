package models

type Phuong struct {
	PhuongID int    `json:"phuong_id" db:"phuong_id"`
	Name     string `json:"name" db:"name"`
	QuanID   int    `json:"quan_id" db:"quan_id"`
}
