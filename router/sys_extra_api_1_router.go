package router

import (
	v1 "QingYin/api/v1"

	"github.com/gin-gonic/gin"
)

type extraApi_1Router struct{}

//初始化额外接口路由一
func (ex *extraApi_1Router) InitExtraApi_1Router(Router *gin.RouterGroup) {
	api := v1.ApiGroups
	{
		Router.POST("/douyin/favorite/action/", api.FavoriteAction)
		Router.GET("/douyin/favorite/list/", api.FavoriteList)
		Router.POST("/douyin/comment/action/", api.CommentAction)
		Router.GET("/douyin/comment/list/", api.CommentList)
	}
}
