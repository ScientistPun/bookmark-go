package backend

import (
	"fmt"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"webstack/lib"
)

func SettingAction(routerGroup *gin.RouterGroup, engin *gin.Engine) {
	routerGroup.GET("/setting", func(ctx *gin.Context) {
		if !lib.PermissionCheck(ctx) {
			ctx.Request.URL.Path = lib.GlobalConfig.System.SignInUrl
			engin.HandleContext(ctx)
		}
	}, settingHandle)
	routerGroup.POST("/setting/basic", AuthHandleMiddleware, editBasicHandle)
	routerGroup.POST("/setting/auth", AuthHandleMiddleware, editAuthHandle)
}

func settingHandle(ctx *gin.Context) {
	lib.LoadConfiguration()
	skins, err := lib.LoadSkins()
	if err != nil {
		fmt.Println(err)
	}
	gintemplate.HTML(ctx, http.StatusOK, lib.GlobalRoutes[0].Url, gin.H{
		"title":     lib.GlobalConfig.System.Name + " - " + lib.GlobalRoutes[0].Title,
		"glob_conf": lib.GlobalConfig,
		"glob_nav":  lib.GlobalRoutes,
		"skins":     skins,
		"useSkin":   lib.GlobalConfig.Service.Skin,
		"uri":       lib.GlobalRoutes[0].Url,
	})
}

func editBasicHandle(ctx *gin.Context) {
	config := lib.GlobalConfig

	port, b := ctx.GetPostForm("port")
	if !b || len(port) <= 1 {
		lib.RespConflict(ctx, "请输入正确的端口号！")
	}
	portInt, _ := strconv.ParseInt(port, 10, 64)
	if portInt < 0 {
		lib.RespConflict(ctx, "请输入正确的端口号！")
	}
	config.Service.Port = port

	skin, b := ctx.GetPostForm("skin")
	if !b {
		lib.RespConflict(ctx, "请选择使用的皮肤")
	}
	config.Service.Skin = skin

	_ssl, _ := ctx.GetPostForm("ssl")
	ssl, _ := strconv.ParseInt(_ssl, 10, 64)
	config.Service.SSLMode = ssl == 1

	csr, b1 := ctx.GetPostForm("csr")
	key, b2 := ctx.GetPostForm("key")
	if config.Service.SSLMode && (!b1 || !b2) && len(csr) < 1 && len(key) < 1 {
		lib.RespConflict(ctx, "SSL文件地址有误！")
	}
	config.Service.Csr = csr
	config.Service.Key = key

	err := lib.WriteYamlFile(lib.ConfigPath, config)
	if err != nil {
		lib.RespConflict(ctx, err.Error())
	}

	lib.GlobalConfig.Service.Skin = skin
	lib.RespSuccess(ctx, "保存成功!")
	return
}

func editAuthHandle(ctx *gin.Context) {
	config := lib.GlobalConfig

	username, b := ctx.GetPostForm("username")
	if !b && len(strings.TrimSpace(username)) < 1 {
		lib.RespConflict(ctx, "用户名不能为空")
	}
	config.Auth.Username = username

	pwd, _ := ctx.GetPostForm("password")
	if len(pwd) > 0 {
		config.Auth.Password = pwd
	}
	err := lib.WriteYamlFile(lib.ConfigPath, &config)
	if err != nil {
		lib.RespConflict(ctx, err.Error())
	}
	lib.RespSuccess(ctx, "保存成功!")
	return
}
