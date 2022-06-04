package response

//视频流接口响应
type FeedResponse struct {
	Status            //状态信息
	NextTime  int     `json:"next_time"`  //本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	VideoList []Video `json:"video_list"` //视频列表
}
