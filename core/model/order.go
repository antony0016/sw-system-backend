package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID  uint   `gorm:"NOT NULL;" json:"user_id,omitempty"`
	Price   int    `gorm:"NOT NULL;" json:"price,omitempty"`
	Status  int    `gorm:"NOT NULL;" json:"status,omitempty"`
	Date    string `gorm:"NOT NULL;" json:"date,omitempty"`
	Mileage int    `gorm:"NOT NULL;" json:"mileage,omitempty"`
	User    User
}
