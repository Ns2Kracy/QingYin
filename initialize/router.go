package initialize

import (
	"QinYin/global"
	"QinYin/middleware"
	"QinYin/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	r := gin.Default()

	//提供文件上传文件路径
	r.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path))

	global.GVA_LOG.Info("use middleware logger")

	//Api分组
	apiRouter := r.Group("/douyin")

	//路由初始化
	apiRouter.Use(middleware.JWTAuth())
	{
		router.ApiRouters.InitBasicRouter(apiRouter)
		router.ApiRouters.InitExtraApi_1Router(apiRouter)
		router.ApiRouters.InitExtraApi_2Router(apiRouter)
	}
	global.GVA_LOG.Info("router register success")
	return r
}
