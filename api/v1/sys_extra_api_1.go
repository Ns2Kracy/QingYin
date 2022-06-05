package v1

import (
	"QingYin/global"
	"QingYin/model/system/request"
	"QingYin/model/system/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type extraApi_1 struct{}

var (
	action_type_like        = 1 //说明：1-点赞，2-取消点赞
	action_type_unlike      = 2
	action_type_comment     = 1 // 说明：用户填写的评论内容，在action_type=1的时候使用
	action_type_del_comment = 2 // 说明：要删除的评论id，在action_type=2的时候使用
)

// 登录用户对视频的点赞和取消点赞操作
func (*extraApi_1) FavoriteAction(c *gin.Context) {
	var r request.FavoriteActionRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		global.GVA_LOG.Error("[FavoriteAction] 解析请求参数失败", zap.Error(err))
		return
	}
	videoId := r.Video
	actionType := r.Action_type
	if actionType == uint(action_type_like) {
		// 点赞
		err := feedService.LikeVideo(videoId)
		if err != nil {
			global.GVA_LOG.Error("[FavoriteAction] 点赞失败", zap.Error(err))
			return
		}
		status := response.Status{StatusCode: SUCCESS, StatusMsg: "点赞成功"}
		c.JSON(http.StatusOK, response.FavoriteActionResponse{Status: status})
	}
	if actionType == uint(action_type_unlike) {
		// 取消点赞
		err := feedService.UnLikeVideo(videoId)
		if err != nil {
			global.GVA_LOG.Error("[FavoriteAction] 取消点赞失败", zap.Error(err))
			return
		}
		status := response.Status{StatusCode: SUCCESS, StatusMsg: "取消点赞成功"}
		c.JSON(http.StatusOK, response.FavoriteActionResponse{Status: status})
	}

}

// 获取点赞列表
func (*extraApi_1) FavoriteList(c *gin.Context) {
}

// 评论操作
func (*extraApi_1) CommentAction(c *gin.Context) {

}

// 查看评论列表，按时间排序
func (*extraApi_1) CommentList(c *gin.Context) {
}
