package controllers

import "github.com/gin-gonic/gin"

type ServerController struct{}

func (receiver ServerController) Ping(c *gin.Context) {
	c.JSON(200, "pong")
	c.Abort()
	return
}
