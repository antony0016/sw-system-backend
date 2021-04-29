package model

import "gorm.io/gorm"

// JSON tag use snake_case.

type User struct {
	gorm.Model
	Username     string `gorm:"NOT NULL;" json:"username,omitempty"`
	LicensePlate string `gorm:"NOT NULL;" json:"license_plate,omitempty"`
	Mileage      int    `gorm:"NOT NULL;" json:"mileage,omitempty"`
	Phone        string `json:"phone,omitempty"`
}
