package main

import (
	"fmt"
	"net/http"
	"webstack/lib"
	"webstack/router"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

func init() {
	lib.LoadConfiguration()
	lib.LoadRoutes()
}

func main() {
	config := lib.GlobalConfig

	app := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	app.Use(ginsession.New())

	skinDir := "./skin"
	app.LoadHTMLFiles("./admin/404.tpl", "./admin/login.tpl", fmt.Sprintf("%v/%v/index.html", skinDir, lib.GlobalConfig.Service.Skin))

	app.Static("/skin", "./skin")
	app.NoRoute(func(ctx *gin.Context) {
		// 实现内部重定向
		ctx.HTML(http.StatusNotFound, "404.html", gin.H{})
	})

	router.InitFrontend(app)
	router.InitBackend(app)
	addr := ":" + config.Service.Port
	if config.Service.SSLMode {
		csrKey := fmt.Sprintf("./data/%s", config.Service.Csr)
		priKey := fmt.Sprintf("./data/%s", config.Service.Key)
		app.RunTLS(addr, csrKey, priKey)
	} else {
		app.Run(addr)
	}
}
