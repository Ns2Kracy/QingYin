package request

import "os"

//投稿接口请求
type PublishActionRequest struct {
	Data  *os.File //视频数据
	Token string   //用户鉴权Token
	Title string   //视频标题
}

//发布列表接口请求
type PublishListRequest struct {
	Token string //用户鉴权Token
	ID    uint   //用户ID
}
