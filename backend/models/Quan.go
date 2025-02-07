package models

type Quan struct {
	QuanID int    `json:"quan_id" db:"quan_id"`
	Name   string `json:"name" db:"name"`
	TPID   int    `json:"tp_id" db:"tp_id"`
}
