package request

import (
	"mime/multipart"
)

//投稿接口请求
type PublishActionRequest struct {
	Data  *multipart.FileHeader `form:"data"`  //视频数据
	Token string                `form:"token"` //用户鉴权Token
	Title string                `form:"title"` //视频标题
}

//发布列表接口请求
type PublishListRequest struct {
	Token  string `form:"token"`   //用户鉴权Token
	UserID uint   `form:"user_id"` //用户ID
}

type FavoriteActionRequest struct {
	Token       string `form:"token"`       //用户鉴权Token
	Video       uint   `form:"video"`       //视频ID
	Action_type uint   `form:"action_type"` //操作类型
}

type CommentActionRequest struct {
	Token string `form:"token"` //用户鉴权Token
	Video uint   `form:"video"` //视频ID
	Text  string `form:"text"`  //评论内容
}

type FollowActionRequest struct {
	Token       string `form:"token"`       //用户鉴权Token
	Video       uint   `form:"video"`       //视频ID
	Action_type uint   `form:"action_type"` //操作类型
}
