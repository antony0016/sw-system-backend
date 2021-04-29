package v1

import (
	controller "github.com/antony0016/sw-system-backend/controller/v1"
	"github.com/gin-gonic/gin"
)

func SetItemRouter(router *gin.RouterGroup) {
	router.GET("/items", controller.AllItem)
	router.GET("/items/:id", controller.OneItem)
	router.POST("/items", controller.CreateItem)
	router.PUT("/items/:id", controller.UpdateItem)
	router.DELETE("/items/:id", controller.DeleteItem)
}
