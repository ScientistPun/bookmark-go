package router

import (
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"net/http"
	"webstack/lib"
)

func InitFrontend(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/index"
		app.HandleContext(ctx)
	})
	app.GET("/index", indexHandle)

	app.GET("/api/config", func(ctx *gin.Context) {
		lib.RespJson(ctx, http.StatusOK, lib.GlobalConfig.System)
	})

	app.GET("/api/bookmark", bookmarkHandle)
}

func indexHandle(ctx *gin.Context) {
	gintemplate.HTML(ctx, http.StatusOK, "index.html", gin.H{})
}

func bookmarkHandle(ctx *gin.Context) {
	var bookmark []lib.Tag
	bookmark, _ = lib.LoadBookmark()

	searchEngine, _ := lib.LoadSearchEngine()

	lib.RespJson(ctx, http.StatusOK, gin.H{"bookmark": bookmark, "search": searchEngine})
}
