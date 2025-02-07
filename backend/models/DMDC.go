package models

import "time"

type DMDC struct {
	DMDCID  int       `json:"dmdc_id" db:"dmdc_id"`
	Name    string    `json:"name" db:"name"`
	Loai    *string   `json:"loai" db:"loai"` // Có thể NULL
	NgayTao time.Time `json:"ngay_tao" db:"ngay_tao"`
}
