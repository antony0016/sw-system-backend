package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "pong",
	})
}

func Reply(c *gin.Context) {
	answer := c.PostForm("question") + ", world"
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"answer": answer,
	})
}
