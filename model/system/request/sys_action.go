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
