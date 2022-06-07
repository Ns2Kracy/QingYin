package response

type Status struct {
	StatusCode int    `json:"status_code"` //状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  //返回状态描述
}

type Video struct {
	ID            uint   `json:"id"`             //视频唯一标识
	Title         string `json:"title"`          //视频标题
	Author        User   `json:"author"`         //视频作者信息
	PlayURL       string `json:"play_url"`       //视频播放地址
	CoverURL      string `json:"cover_url"`      //视频封面地址
	FavoriteCount int    `json:"favorite_count"` //视频的点赞总数
	CommentCount  int    `json:"comment_count"`  //视频的评论总数
	IsFavorite    bool   `json:"is_favorite"`    //true-已点赞，false-未点赞
}

type User struct {
	ID            uint   `json:"id"`             //用户ID
	Username      string `json:"name"`           //用户名
	FollowCount   int64  `json:"follow_count"`   //关注总数
	FollowerCount int64  `json:"follower_count"` //粉丝总数
	IsFollow      bool   `json:"is_follow"`      //是否关注:true已关注,false未关注
}

type Comment struct {
	ID         uint   `json:"id"`          //评论id
	User       User   `json:"user"`        //评论用户信息
	Content    string `json:"content"`     //评论内容
	CreateDate string `json:"create_date"` //评论发布日期，格式 mm-dd
}
