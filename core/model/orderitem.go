package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID    uint `gorm:"NOT NULL;" json:"order_id,omitempty"`
	CategoryID uint `gorm:"NOT NULL;" json:"category_id,omitempty"`
	ItemID     uint `gorm:"NOT NULL;" json:"item_id,omitempty"`
	Count      int  `gorm:"NOT NULL;" json:"count,omitempty"`
	Price      int  `gorm:"NOT NULL;" json:"price,omitempty"`
	Item       Item
	Category   Category
	Order      Order
}
