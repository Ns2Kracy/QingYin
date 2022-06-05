package request

//关注操作接口请求
type FocusActionRequest struct {
	Token      string `form:"token"`       //用户鉴权token
	ToUserID   uint   `form:"to_user_id"`  //对方用户id
	ActionType int    `form:"action_type"` //1-关注，2-取消关注
}

//关注列表接口请求
type FollowListRequest struct {
	UserID uint   `form:"user_id"` //用户id
	Token  string `form:"token"`   //用户鉴权token
}

//粉丝列表接口请求
type FollowerListRequest struct {
	UserID uint   `form:"user_id"` //用户id
	Token  string `form:"token"`   //用户鉴权token
}
