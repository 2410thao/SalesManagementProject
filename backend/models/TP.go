package models

type TP struct {
	TPID   int    `json:"tp_id" db:"tp_id"`
	Name   string `json:"name" db:"name"`
	TinhID int    `json:"tinh_id" db:"tinh_id"`
}
