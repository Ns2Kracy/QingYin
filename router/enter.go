package router

type ApiRouter struct {
	basicApiRouter
	extraApi_1Router
	extraApi_2Router
}

var ApiRouters = new(ApiRouter)
