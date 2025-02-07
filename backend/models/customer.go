package models

import "time"

type Customer struct {
	CustomerID    int        `json:"customer_id" db:"customer_id"`
	FullName      string     `json:"fullname" db:"fullname"`
	Email         string     `json:"email" db:"email"`
	PhoneNumber   *string    `json:"phonenumber" db:"phonenumber"`       // Có thể NULL
	LoyaltyPoints int        `json:"loyalty_points" db:"loyalty_points"` // Mặc định 0
	CreateAt      time.Time  `json:"create_at" db:"create_at"`
	Status        *string    `json:"status" db:"status"`               // Có thể NULL
	CustomerType  *string    `json:"customer_type" db:"customer_type"` // Có thể NULL
	Address       *string    `json:"address" db:"address"`             // Có thể NULL
	BirthDay      *time.Time `json:"birthday" db:"birthday"`           // Có thể NULL
}
