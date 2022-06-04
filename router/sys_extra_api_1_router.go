package router

import (
	v1 "QingYin/api/v1"
	"github.com/gin-gonic/gin"
)

type extraApi_1Router struct{}

//初始化额外接口路由
func (ex *extraApi_1Router) InitExtraApi_1Router(Router *gin.RouterGroup) {
	api := v1.ApiGroups
	{
		Router.POST("/favorite/action/", api.FavoriteAction)
		Router.GET("/favorite/list/", api.FavoriteList)
		Router.POST("/comment/action", api.CommentAction)
		Router.GET("/comment/list", api.CommentList)
	}
}
