package backend

import (
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"webstack/lib"
)

func BookmarkAction(routerGroup *gin.RouterGroup, engin *gin.Engine) {
	routerGroup.GET("/bookmark", func(ctx *gin.Context) {
		if !lib.PermissionCheck(ctx) {
			ctx.Request.URL.Path = lib.GlobalConfig.System.SignInUrl
			engin.HandleContext(ctx)
		}
	}, bookmarkHandle)
	routerGroup.PUT("/bookmark/add-tag", AuthHandleMiddleware, addTagHandle)
	routerGroup.PUT("/bookmark/add-link", AuthHandleMiddleware, addLinkHandle)
	routerGroup.GET("/bookmark/refresh-link/:tagName/:sort", AuthHandleMiddleware, refreshLinkHandle)
	routerGroup.POST("/bookmark/edit-file", AuthHandleMiddleware, editBookmarkTextHandle)
}

func bookmarkHandle(ctx *gin.Context) {
	var bookmark []lib.Tag
	bookmark, _ = lib.LoadBookmark()
	bookmarkText, _ := lib.GetBookmarkTextFrom(bookmark)
	gintemplate.HTML(ctx, http.StatusOK, lib.GlobalRoutes[1].Url, gin.H{
		"title":        lib.GlobalConfig.System.Name + " - " + lib.GlobalRoutes[1].Title,
		"glob_conf":    lib.GlobalConfig,
		"glob_nav":     lib.GlobalRoutes,
		"uri":          lib.GlobalRoutes[1].Url,
		"bookmark":     bookmark,
		"bookmarkText": bookmarkText,
	})
}

func addTagHandle(ctx *gin.Context) {
	title, titlePost := ctx.GetPostForm("title")
	name, namePost := ctx.GetPostForm("name")
	if !titlePost || !namePost || len(title) <= 0 || len(name) <= 0 {
		lib.RespConflict(ctx, "请输入正确的参数")
	}
	icon, _ := ctx.GetPostForm("icon")
	if len(icon) <= 0 {
		icon = "fa-tags"
	}

	var bookmark []lib.Tag
	bookmark, _ = lib.LoadBookmark()
	bookmark = append(bookmark, lib.Tag{Title: title, Name: name, Icon: icon, Links: []lib.Link{}})

	err := lib.WriteYamlFile(lib.BookmarkPath, bookmark)
	if err != nil {
		lib.RespConflict(ctx, err.Error())
	}
	lib.RespJson(ctx, http.StatusCreated, gin.H{"msg": "添加成功!"})
	return
}

func addLinkHandle(ctx *gin.Context) {
	title, _ := ctx.GetPostForm("title")
	url, _ := ctx.GetPostForm("url")
	selectTag, _ := ctx.GetPostForm("tag")
	if len(title) <= 0 || len(url) <= 0 {
		lib.RespConflict(ctx, "请输入正确的参数!")
	}

	result := false
	var bookmark []lib.Tag
	bookmark, _ = lib.LoadBookmark()
	for i, tag := range bookmark {
		if strings.EqualFold(tag.Name, selectTag) {
			newTag := bookmark[i]
			desc, icon, err := lib.GetHtmlInfo(url)
			if err != nil {
				lib.RespConflict(ctx, err.Error())
			}
			newTag.Links = append(newTag.Links, lib.Link{Title: title, Desc: desc, Url: url, Icon: icon})
			bookmark[i] = newTag
			result = true
			break
		}
	}
	if !result {
		lib.RespConflict(ctx, "标签参数错误!")
	}

	err := lib.WriteYamlFile(lib.BookmarkPath, bookmark)
	if err != nil {
		lib.RespConflict(ctx, err.Error())
	}
	lib.RespJson(ctx, http.StatusCreated, gin.H{"msg": "添加成功!"})
	return
}

func refreshLinkHandle(ctx *gin.Context) {
	tagName := ctx.Param("tagName")
	sortStr := ctx.Param("sort")
	sort, err := strconv.ParseInt(sortStr, 10, 64)
	if err != nil {
		lib.RespConflict(ctx, err.Error())
	}

	result := false
	var bookmark []lib.Tag
	bookmark, _ = lib.LoadBookmark()
	for i, tag := range bookmark {
		if strings.EqualFold(tag.Name, tagName) {
			oldTag := bookmark[i]
			desc, icon, err := lib.GetHtmlInfo(oldTag.Links[sort].Url)
			if err != nil {
				lib.RespConflict(ctx, err.Error())
				return
			}
			oldTag.Links[sort].Desc = desc
			oldTag.Links[sort].Icon = icon
			bookmark[i] = oldTag
			result = true
			break
		}
	}
	if !result {
		lib.RespConflict(ctx, "标签参数错误!")
	}
	err = lib.WriteYamlFile(lib.BookmarkPath, bookmark)
	if err != nil {
		lib.RespConflict(ctx, err.Error())
	}
	lib.RespSuccess(ctx, "更新成功!")
	return
}

func editBookmarkTextHandle(ctx *gin.Context) {
	txt, _ := ctx.GetPostForm("content")
	err := lib.SetBookmarkText(txt)
	if err != nil {
		lib.RespConflict(ctx, err.Error())
	}
	lib.RespSuccess(ctx, "保存成功!")
	return
}
