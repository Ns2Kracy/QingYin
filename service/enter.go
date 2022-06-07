package service

type ServiceGroup struct {
	UserService
	FeedService
	PublishService
	RelationService
	FavoriteService
	CommentService
}

var ServiceGroups = new(ServiceGroup)
