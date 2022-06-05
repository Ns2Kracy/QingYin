package router

import (
	v1 "QingYin/api/v1"

	"github.com/gin-gonic/gin"
)

type extraApi_2Router struct{}

//初始化额外接口路由二
func (ex *extraApi_1Router) InitExtraApi_2Router(Router *gin.RouterGroup) {
	api := v1.ApiGroups
	{
		Router.POST("/douyin/relation/action/", api.Relation)
		Router.GET("/douyin/relation/follow/list/", api.FollowList)
		Router.GET("/douyin/relation/follower/list/", api.FollowerList)
	}
}
