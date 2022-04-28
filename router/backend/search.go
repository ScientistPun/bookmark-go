package backend

import (
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"net/http"
	"webstack/lib"
)

func SearchAction(routerGroup *gin.RouterGroup, engin *gin.Engine) {
	routerGroup.GET("/search", func(ctx *gin.Context) {
		if !lib.PermissionCheck(ctx) {
			ctx.Request.URL.Path = lib.GlobalConfig.System.SignInUrl
			engin.HandleContext(ctx)
		}
	}, searchHandle)
	routerGroup.POST("/search/edit", AuthHandleMiddleware, editSearchHandle)
}

func searchHandle(ctx *gin.Context) {
	searchEngineText, _ := lib.GetSearchEngineTextFrom()
	gintemplate.HTML(ctx, http.StatusOK, lib.GlobalRoutes[2].Url, gin.H{
		"title":            lib.GlobalConfig.System.Name + " - " + lib.GlobalRoutes[1].Title,
		"glob_conf":        lib.GlobalConfig,
		"glob_nav":         lib.GlobalRoutes,
		"uri":              lib.GlobalRoutes[2].Url,
		"searchEngineText": searchEngineText,
	})
}

func editSearchHandle(ctx *gin.Context) {
	txt, _ := ctx.GetPostForm("content")
	err := lib.SetSearchEngineText(txt)
	if err != nil {
		lib.RespConflict(ctx, err.Error())
	}
	lib.RespSuccess(ctx, "保存成功!")
	return
}
