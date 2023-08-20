package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"qnurye/qnurye.me/pkg/config"
	"qnurye/qnurye.me/routers"
)

func main() {
	cfg := config.Get()
	r := routers.SetUp()
	gin.SetMode(cfg.GetString("server.run_mode"))

	err := r.Run(fmt.Sprintf("%s:%d", cfg.GetString("server.host"), cfg.GetInt("server.port")))

	if err != nil {
		return
	}
}
