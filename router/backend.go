package router

import (
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"webstack/lib"
	"webstack/router/backend"
)

var backendGroup *gin.RouterGroup

func InitBackend(app *gin.Engine) {
	backMW := gintemplate.NewMiddleware(gintemplate.TemplateConfig{
		Root:         "./admin",
		Extension:    ".tpl",
		Master:       "layouts/layout",
		Partials:     []string{},
		DisableCache: true,
	})

	backendGroup = app.Group("/manage", backMW)
	{
		backendGroup.Static("/static", "./admin/style")
		backendGroup.GET("/", func(ctx *gin.Context) {
			ctx.Request.URL.Path = "/manage/" + lib.GlobalRoutes[0].Url
			app.HandleContext(ctx)
		})

		backend.SignInAction(app)
		backend.SettingAction(backendGroup, app)
		backend.SearchAction(backendGroup, app)
		backend.BookmarkAction(backendGroup, app)
	}
}
