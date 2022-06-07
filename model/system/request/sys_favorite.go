package request

//点赞操作接口请求
type LikeActionRequest struct {
	Token      string `form:"token"`       //用户鉴权token
	VideoID    uint   `form:"video_id"`    //视频id
	ActionType int    `form:"action_type"` //1-点赞，2-取消点赞
}

//点赞列表接口请求
type LikeListRequest struct {
	UserID uint   `form:"user_id"` //用户id
	Token  string `form:"token"`   //用户鉴权token
}
