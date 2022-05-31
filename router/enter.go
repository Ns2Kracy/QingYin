package router

//路由封装
type ApiRouter struct {
	basicApiRouter
	extraApi_1Router
	extraApi_2Router
}

//路由组实例化
var ApiRouters = new(ApiRouter)
