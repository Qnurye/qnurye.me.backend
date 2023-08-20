package routers

import (
	"github.com/gin-gonic/gin"
	"qnurye/qnurye.me/controllers"
)

type Api struct {
	router *gin.RouterGroup
}

func (a Api) Setup() *gin.RouterGroup {
	server := new(controllers.ServerController)
	posts := new(controllers.PostController)

	a.router.Any("/ping", server.Ping)
	p := a.router.Group("posts")
	{
		p.GET("", posts.Index)
	}

	return a.router
}
