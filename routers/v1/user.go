package v1

import (
	controller "github.com/antony0016/sw-system-backend/controller/v1"
	"github.com/gin-gonic/gin"
)

func SetUserRouter(router *gin.RouterGroup) {
	router.GET("/users", controller.AllUser)
	router.GET("/users/:id", controller.OneUser)
	router.POST("/users", controller.CreateUser)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)
}
