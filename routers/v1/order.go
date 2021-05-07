package v1

import (
	controller "github.com/antony0016/sw-system-backend/controller/v1"
	"github.com/gin-gonic/gin"
)

func SetOrderRouter(router *gin.RouterGroup) {
	router.GET("/orders", controller.AllOrder)
	//router.GET("/orders/:id", controller.OneOrder)
	router.POST("/orders", controller.CreateOrder)
	router.PUT("/orders/:id", controller.UpdateOrder)
	router.DELETE("/orders/:id", controller.DeleteOrder)
}
