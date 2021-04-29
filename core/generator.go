package core

import (
	"fmt"
	"github.com/antony0016/sw-system-backend/core/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{}, &model.Category{}, &model.Item{},
		&model.Order{}, &model.OrderItem{})
	if err != nil {
		fmt.Println(err)
	}
}
