package v1

import (
	"encoding/json"
	"fmt"
	"github.com/antony0016/sw-system-backend/core/model"
	"github.com/antony0016/sw-system-backend/pkg/httpresponse"
	"github.com/antony0016/sw-system-backend/services/pgdb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AllOrderItem(orderID uint) ([]model.OrderItem, error) {
	var orderItems []model.OrderItem
	db := pgdb.DB
	err := db.Where("order_id=?", orderID).Find(&orderItems).Error
	return orderItems, err
}

func CreateOrderItems(orderID uint, orderItems []model.OrderItem) error {
	var err error
	db := pgdb.DB
	for _, orderItem := range orderItems {
		orderItem.OrderID = orderID
		orderItem.ItemID = orderItem.ID
		orderItem.ID = 0
		err = db.Save(&orderItem).Error
		if err != nil {
			return err
		}
	}
	return err
}

func DeleteOrderItems(orderID int) error {
	db := pgdb.DB
	var orderItems []model.OrderItem
	var err error
	db.Where("order_id=?", orderID).First(&orderItems)
	if len(orderItems) == 0 {
		return err
	}
	err = db.Unscoped().Where("order_id=?", orderID).Delete(&model.OrderItem{}).Error
	return err
}

func AllOrder(c *gin.Context) {
	var ordersData []map[string]interface{}
	var orders []model.Order
	var response model.Response
	db := pgdb.DB
	if err := db.Find(&orders).Error; err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
		httpresponse.MakeResponse(c, response)
		return
	}
	if len(orders) == 0 {
		response.ErrorCode = http.StatusNoContent
		response.ErrorMessage = "No orders found."
		httpresponse.MakeResponse(c, response)
		return
	}
	for _, order := range orders {
		var items []model.OrderItem
		var err error
		items, err = AllOrderItem(order.ID)
		if err != nil {
			fmt.Println(err)
			response.ErrorCode = http.StatusInternalServerError
			httpresponse.MakeResponse(c, response)
			return
		}
		ordersData = append(ordersData, map[string]interface{}{
			"ID":      order.ID,
			"user_id": order.UserID,
			"price":   order.Price,
			"status":  order.Status,
			"date":    order.Date,
			"mileage": order.Mileage,
			"items":   items,
		})
	}
	//data, _ := json.Marshal(ordersData)
	response.Data = ordersData
	httpresponse.MakeResponse(c, response)
}

//func OneOrder(c *gin.Context) {
//	var orders []model.Order
//	var response model.Response
//	db := pgdb.DB
//	orderID, _ := strconv.Atoi(c.Param("id"))
//	if err := db.First(&orders, orderID).Error; err != nil {
//		fmt.Println(err)
//		response.ErrorCode = http.StatusInternalServerError
//	}
//	response.Data = orders
//	httpresponse.MakeResponse(c, response)
//}

func CreateOrder(c *gin.Context) {
	var items []model.OrderItem
	var response model.Response
	db := pgdb.DB
	userID, _ := strconv.Atoi(c.PostForm("user_id"))
	price, _ := strconv.Atoi(c.PostForm("price"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	mileage, _ := strconv.Atoi(c.PostForm("mileage"))
	itemsJsonStr := c.PostForm("items")
	_ = json.Unmarshal([]byte(itemsJsonStr), &items)
	order := model.Order{
		UserID:  uint(userID),
		Price:   price,
		Status:  status,
		Mileage: mileage,
		Date:    c.PostForm("date"),
	}
	if err := db.Save(&order).Error; err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
		httpresponse.MakeResponse(c, response)
		return
	}

	if err := CreateOrderItems(order.ID, items); err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
	}
	response.Message = "Create order successfully."
	httpresponse.MakeResponse(c, response)
}

func UpdateOrder(c *gin.Context) {
	var order model.Order
	var response model.Response
	db := pgdb.DB
	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&order, id)
	if order.ID == 0 {
		response.ErrorCode = http.StatusNoContent
		response.ErrorMessage = "Order not found."
		httpresponse.MakeResponse(c, response)
		return
	}
	status, _ := strconv.Atoi(c.PostForm("status"))
	if err := db.Model(&order).Update("status", status).Error; err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
	}
	response.Message = "Order updated successfully."
	httpresponse.MakeResponse(c, response)
}

func DeleteOrder(c *gin.Context) {
	var (
		order    model.Order
		response model.Response
	)
	db := pgdb.DB
	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&order, id)
	if order.ID == 0 {
		response.ErrorCode = http.StatusNoContent
		response.ErrorMessage = "Order not found."
		httpresponse.MakeResponse(c, response)
		return
	}
	if err := DeleteOrderItems(id); err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
		httpresponse.MakeResponse(c, response)
		return
	}
	if err := db.Unscoped().Delete(&order).Error; err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
	}
	response.Message = "order deleted successfully."
	httpresponse.MakeResponse(c, response)
}
