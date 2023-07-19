package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUp() *gin.Engine {
	r := gin.Default()

	r.GET("/Hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "World",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r
}
