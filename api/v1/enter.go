package v1

import "QingYin/service"

type ApiGroup struct {
	basicApi
	extraApi_1
	extraApi_2
}

//service层实例化
var (
	userService    = service.ServiceGroups.UserService
	feedService    = service.ServiceGroups.FeedService
	publishService = service.ServiceGroups.PublishService
)

var ApiGroups = new(ApiGroup)
