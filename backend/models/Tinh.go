package models

type Tinh struct {
	TinhID int    `json:"tinh_id" db:"tinh_id"`
	Name   string `json:"name" db:"name"`
}
