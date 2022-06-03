package request

//视频流接口请求
type FeedRequest struct {
	LatestTime int    //限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      string //用户登录状态下设置Token
}
