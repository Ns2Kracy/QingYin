package response

//评论操作接口响应
type CommentActionResponse struct {
	Status
	Comment
}

//评论列表接口响应
type CommentListResponse struct {
	Status
	CommentList []Comment `json:"comment_list"` //评论列表
}
