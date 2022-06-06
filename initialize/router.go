package initialize

import (
	"QingYin/global"
	"QingYin/middleware"
	"QingYin/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	//提供文件上传文件路径
	r.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path))
	global.GVA_LOG.Info("文件上传路径:", zap.String("filePath:", global.GVA_CONFIG.Local.Path))

	global.GVA_LOG.Info("use middleware logger")

	//Api分组
	apiRouter := r.Group("/douyin")
	//公共路由不受登陆状态限制?
	// publicRouter := r.Group("")
	// publicRouter.GET("/douyin/feed", v1.ApiGroups.Feed)

	//路由初始化
	apiRouter.Use(middleware.JWTAuth())
	{
		router.ApiRouters.InitBasicRouter(apiRouter)
		router.ApiRouters.InitExtraApi_1Router(apiRouter)
		// router.ApiRouters.InitExtraApi_2Router(apiRouter)
	}
	global.GVA_LOG.Info("router register success")
	return r
}
