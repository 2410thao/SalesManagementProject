package models

type Category struct {
	CategoryID  int    `json:"category_id" db:"category_id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}
