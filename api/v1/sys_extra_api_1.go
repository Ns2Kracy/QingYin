package v1

import "github.com/gin-gonic/gin"

type extraApi_1 struct{}

// 登录用户对视频的点赞和取消点赞操作
func (*extraApi_1) FavoriteAction(c *gin.Context) {
	// 点赞
	// 取消点赞
}

// 获取点赞列表
func (*extraApi_1) FavoriteList(c *gin.Context) {
	// 获取点赞列表
}

// 评论操作
func (*extraApi_1) CommentAction(c *gin.Context) {
	// 发表评论
	// 删除评论
}

// 查看评论列表，按时间排序
func (*extraApi_1) CommentList(c *gin.Context) {
	// 获取评论列表
}
