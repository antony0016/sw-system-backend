package v1

import (
	controller "github.com/antony0016/sw-system-backend/controller/v1"
	"github.com/gin-gonic/gin"
)

func SetCategoryRouter(router *gin.RouterGroup) {
	router.GET("/categories", controller.AllCategory)
	router.GET("/categories/:id", controller.OneCategory)
	router.POST("/categories", controller.CreateCategory)
	router.PUT("/categories/:id", controller.UpdateCategory)
	router.DELETE("/categories/:id", controller.DeleteCategory)
}
