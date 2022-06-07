package response

//关注操作接口响应
type FocusActionResponse struct {
	Status
}

//关注列表接口响应
type FollowListResponse struct {
	Status
	UserList []User `json:"user_list"` //用户信息列表
}

//粉丝列表接口响应
type FollowerListResponse struct {
	Status
	UserList []User `json:"user_list"` //用户信息列表
}
