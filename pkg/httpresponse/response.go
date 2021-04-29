package httpresponse

import (
	"github.com/antony0016/sw-system-backend/core/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// MakeResponse make a http response by one struct
func MakeResponse(c *gin.Context, response model.Response) {
	if response.ErrorCode != 0 {
		if response.ErrorCode == http.StatusInternalServerError {
			response.ErrorMessage = "Database error."
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  response.ErrorCode,
			"message": response.ErrorMessage,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": response.Message,
		"data":    response.Data,
	})
}
