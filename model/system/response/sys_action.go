package response

//投稿接口响应
type PublishActionResponse struct {
	Status //状态信息
}

//发布列表接口响应
type PublishListResponse struct {
	Status            //状态信息
	VideoList []Video `json:"video_list"` //视频列表
}
