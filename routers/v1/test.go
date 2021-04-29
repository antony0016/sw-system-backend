package v1

import (
	"github.com/antony0016/sw-system-backend/controller/v1"
	"github.com/gin-gonic/gin"
)

func SetTestRouter(router *gin.RouterGroup) {
	router.GET("/ping", v1.Ping)
	router.POST("/reply", v1.Reply)
}
