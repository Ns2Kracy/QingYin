package router

import (
	v1 "QingYin/api/v1"

	"github.com/gin-gonic/gin"
)

type basicApiRouter struct{}

// 初始化基本接口路由
func (b *basicApiRouter) InitBasicRouter(Router *gin.RouterGroup) {
	api := v1.ApiGroups
	{
		Router.GET("/feed/", api.Feed)
		Router.GET("/user/", api.UserInfo)
		Router.POST("/user/register/", api.Register)
		Router.POST("/user/login/", api.Login)
		Router.POST("/publish/action/", api.Publish)
		Router.GET("/publish/list/", api.PublishList)
	}
}

// basic apis
// apiRouter.GET("/feed/", controller.Feed)
// apiRouter.GET("/user/", controller.UserInfo)
// apiRouter.POST("/user/register/", controller.Register)
// apiRouter.POST("/user/login/", controller.Login)
// apiRouter.POST("/publish/action/", controller.Publish)
// apiRouter.GET("/publish/list/", controller.PublishList)
