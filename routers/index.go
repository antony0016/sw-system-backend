package routers

import (
	"github.com/antony0016/sw-system-backend/routers/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.RouterGroup) {
	v1Router := router.Group("/v1")
	v1.SetTestRouter(v1Router)
	v1.SetUserRouter(v1Router)
	v1.SetItemRouter(v1Router)
	v1.SetCategoryRouter(v1Router)
	v1.SetOrderRouter(v1Router)
}
