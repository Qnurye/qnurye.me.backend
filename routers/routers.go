package routers

import (
	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	router := gin.New()

	Api{router.Group("api/v1")}.Setup()

	return router
}
