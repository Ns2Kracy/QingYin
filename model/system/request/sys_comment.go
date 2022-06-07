package request

//发布评论接口请求
type CommentActionRequest struct {
	Token       string `form:"token"`        //用户鉴权token
	VideoID     uint   `form:"video_id"`     //视频id
	ActionType  int    `form:"action_type"`  //1-发布评论，2-删除评论
	CommentText string `form:"comment_text"` //用户填写的评论内容，在action_type=1的时候使用
}

//删除评论接口请求
type DeleteCommentActionRequest struct {
	Token      string `form:"token"`       //用户鉴权token
	VideoID    uint   `form:"video_id"`    //视频id
	ActionType int    `form:"action_type"` //1-发布评论，2-删除评论
	CommentID  uint   `form:"comment_id"`  //要删除的评论id，在action_type=2的时候使用
}

//评论列表接口请求
type CommentListRequest struct {
	Token   string `form:"token"`    //用户鉴权token
	VideoID uint   `form:"video_id"` //视频id
}
