package v1

import (
	"fmt"
	"github.com/antony0016/sw-system-backend/core/model"
	"github.com/antony0016/sw-system-backend/pkg/httpresponse"
	"github.com/antony0016/sw-system-backend/services/pgdb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteItemRelate(c *gin.Context, response model.Response, itemID int) error {
	db := pgdb.DB
	//err := db.Unscoped().Where(
	//	"item_id=?", itemID).Delete(&model.OrderItem{}).Error
	//if err != nil {
	//	fmt.Println(err)
	//	response.ErrorCode = http.StatusInternalServerError
	//	httpresponse.MakeResponse(c, response)
	//	return err
	//}
	err := db.Unscoped().Where(
		"item_id=?", itemID).Delete(&model.OrderItem{}).Error
	if err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
		httpresponse.MakeResponse(c, response)
	}
	return err
}

func AllItem(c *gin.Context) {
	var items []model.Item
	var response model.Response
	db := pgdb.DB
	if err := db.Find(&items).Error; err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
	}
	response.Data = items
	httpresponse.MakeResponse(c, response)
}

func OneItem(c *gin.Context) {
	var items []model.Item
	var response model.Response
	db := pgdb.DB
	itemID, _ := strconv.Atoi(c.Param("id"))
	if err := db.First(&items, itemID); err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
	}
	response.Data = items
	httpresponse.MakeResponse(c, response)
}

func CreateItem(c *gin.Context) {
	var item model.Item
	var response model.Response
	db := pgdb.DB
	categoryID, _ := strconv.Atoi(c.PostForm("category_id"))
	price, _ := strconv.Atoi(c.PostForm("price"))
	db.Where("category_id=?", categoryID).
		Where("item_name=?", c.PostForm("item_name")).
		First(&item)
	if item.ID != 0 {
		response.ErrorCode = http.StatusConflict
		response.ErrorMessage = "item already exists"
		httpresponse.MakeResponse(c, response)
		return
	}
	item = model.Item{
		CategoryID: uint(categoryID),
		ItemName:   c.PostForm("item_name"),
		Price:      price,
	}
	if err := db.Save(&item).Error; err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
	}
	response.Message = "Create item successful."
	httpresponse.MakeResponse(c, response)
}

func UpdateItem(c *gin.Context) {
	var item model.Item
	var response model.Response
	db := pgdb.DB
	itemID := c.Param("id")
	db.First(&item, itemID)
	if item.ID == 0 {
		response.ErrorCode = http.StatusNoContent
		response.ErrorMessage = "Item not found."
		httpresponse.MakeResponse(c, response)
		return
	}
	price, _ := strconv.Atoi(c.PostForm("price"))
	err := db.Model(&item).Update("item_name", c.PostForm("item_name")).Error
	err = db.Model(&item).Update("price", price).Error
	if err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
	}
	response.Message = "Item updated successfully."
	httpresponse.MakeResponse(c, response)
}

func DeleteItem(c *gin.Context) {
	var item model.Item
	var response model.Response
	db := pgdb.DB
	itemID, _ := strconv.Atoi(c.Param("id"))
	db.First(&item, itemID)
	if item.ID == 0 {
		response.ErrorCode = http.StatusNoContent
		response.ErrorMessage = "Item not found."
		httpresponse.MakeResponse(c, response)
		return
	}
	if err := DeleteItemRelate(c, response, itemID); err != nil {
		return
	}
	if err := db.Unscoped().Delete(&item).Error; err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
	}
	response.Message = "Item deleted successfully"
	httpresponse.MakeResponse(c, response)
}
