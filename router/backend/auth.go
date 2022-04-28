package backend

import (
	"net/http"
	"strings"
	"webstack/lib"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

func AuthHandleMiddleware(ctx *gin.Context) {
	session := ginsession.FromContext(ctx)
	_, ok := session.Get(lib.GlobalConfig.System.Session)
	if !ok {
		lib.RespUnauthorized(ctx, "没有权限")
	}
	ctx.Next()
}

func SignInAction(app *gin.Engine) {

	app.GET(lib.GlobalConfig.System.SignInUrl, func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.tpl", gin.H{
			"title": lib.GlobalConfig.System.Name + " - " + lib.GlobalRoutes[0].Title,
			"msg":   "please sign in",
		})
	})
	app.POST(lib.GlobalConfig.System.SignInUrl, signInHandle)
	app.GET("/sign-out", func(ctx *gin.Context) {
		session := ginsession.FromContext(ctx)
		session.Delete(lib.GlobalConfig.System.Session)
		ctx.HTML(http.StatusOK, "login.tpl", gin.H{
			"title": lib.GlobalConfig.System.Name + " - " + lib.GlobalRoutes[0].Title,
			"msg":   "please sign in",
		})
	})
}

func signInHandle(ctx *gin.Context) {
	username, namePost := ctx.GetPostForm("username")
	password, pwdPost := ctx.GetPostForm("password")
	if !namePost || !pwdPost || len(username) <= 0 || len(password) <= 0 {
		lib.RespConflict(ctx, "请输入账号和密码")
		return
	}

	config := lib.GlobalConfig
	if !strings.EqualFold(username, config.Auth.Username) || !strings.EqualFold(password, config.Auth.Password) {
		lib.RespUnauthorized(ctx, "账户名或密码不正确")
		return
	}

	session := ginsession.FromContext(ctx)
	session.Set(config.System.Session, 1)
	err := session.Save()
	if err != nil {
		lib.RespJson(ctx, http.StatusBadGateway, gin.H{"msg": err.Error()})
		return
	}

	lib.RespSuccess(ctx, "sign in success.")
}
