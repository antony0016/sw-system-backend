package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	CategoryID uint   `gorm:"NOT NULL;" json:"category_id,omitempty"`
	ItemName   string `gorm:"NOT NULL;" json:"item_name,omitempty"`
	Price      int    `gorm:"NOT NULL;" json:"price,omitempty"`
	Category   Category
}
