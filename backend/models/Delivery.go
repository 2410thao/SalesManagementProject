package models

import "time"

type Delivery struct {
	DeliveryID     int        `json:"delivery_id" db:"delivery_id"`
	OrderID        int        `json:"order_id" db:"order_id"`
	DeliveryStatus *string    `json:"delivery_status" db:"delivery_status"` // Có thể NULL
	CourierName    *string    `json:"courier_name" db:"courier_name"`       // Có thể NULL
	TrackingNumber *string    `json:"tracking_number" db:"tracking_number"` // Có thể NULL
	DeliveryDate   *time.Time `json:"delivery_date" db:"delivery_date"`     // Có thể NULL
}
