package lib

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

type Route struct {
	Title string `yaml:"title"`
	Url   string `yaml:"url"`
	Icon  string `yaml:"icon"`
}

var GlobalRoutes []Route

func InitRoutes() {
	routes := []Route{
		{Title: "系统设置", Icon: "linecons-cog", Url: "setting"},
		{Title: "导航管理", Icon: "linecons-tag", Url: "bookmark"},
		{Title: "搜索引擎", Icon: "linecons-search", Url: "search"},
	}
	WriteYamlFile(RoutesPath, routes)
	GlobalRoutes = routes
}

func LoadRoutes() {
	if !IsFileExists(RoutesPath) {
		InitRoutes()
	} else {
		err := LoadYamlFile(RoutesPath, &GlobalRoutes)
		if err != nil {
			log.Printf("file get err: %v ", err.Error())
		}
	}
}

func PermissionCheck(ctx *gin.Context) bool {
	session := ginsession.FromContext(ctx)
	_, ok := session.Get(GlobalConfig.System.Session)
	return ok
}

func RespSuccess(ctx *gin.Context, msg string) {
	RespJson(ctx, http.StatusOK, gin.H{"msg": msg})
}

func RespUnauthorized(ctx *gin.Context, msg string) {
	RespJson(ctx, http.StatusUnauthorized, gin.H{"msg": msg})
}

func RespConflict(ctx *gin.Context, msg string) {
	RespJson(ctx, http.StatusConflict, gin.H{"msg": msg})
}

func RespJson(ctx *gin.Context, statusCode int, json interface{}) {
	succList := []int{http.StatusOK, http.StatusCreated}
	if InIntArray(statusCode, succList) {
		ctx.JSON(statusCode, json)
	} else {
		ctx.AbortWithStatusJSON(statusCode, json)
	}
}
