package service

type ServiceGroup struct {
	UserService
	FeedService
	PublishService
}

var ServiceGroups = new(ServiceGroup)
