package response

//点赞操作接口响应
type LikeActionResponse struct {
	Status
}

//点赞列表接口响应
type LikeListResponse struct {
	Status
	VideoList []Video `json:"video_list"` //视频列表
}
