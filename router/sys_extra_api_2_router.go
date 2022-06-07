package router

import (
	v1 "QingYin/api/v1"

	"github.com/gin-gonic/gin"
)

type extraApi_2Router struct{}

//初始化额外接口路由二
func (ex *extraApi_2Router) InitExtraApi_2Router(Router *gin.RouterGroup) {
	api := v1.ApiGroups
	{
		Router.POST("/favorite/action/", api.FavoriteAction)
		Router.GET("/favorite/list/", api.FavoriteList)
		Router.POST("/comment/action/", api.CommentAction)
		Router.GET("/comment/list/", api.CommentList)
	}
}
