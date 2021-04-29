package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string `gorm:"not null;unique" json:"category_name,omitempty"`
}
