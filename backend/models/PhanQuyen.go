package models

type PhanQuyen struct {
	PQID  int     `json:"pq_id" db:"pq_id"`
	Name  string  `json:"name" db:"name"`
	CapDo *string `json:"cap_do" db:"cap_do"` // Có thể NULL
	MoTa  *string `json:"mo_ta" db:"mo_ta"`   // Có thể NULL
}
