package service

type ServiceGroup struct {
	UserService
	FeedService
	PublishService
	RelationService
}

var ServiceGroups = new(ServiceGroup)
