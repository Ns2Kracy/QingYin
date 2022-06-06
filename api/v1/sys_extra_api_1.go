package v1

import (
	"QingYin/global"
	"QingYin/model/system/request"
	"QingYin/model/system/response"
	"QingYin/utils"
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

// 登录用户对视频的点赞和取消点赞操作>>>未测试
func (*extraApi_1) FavoriteAction(c *gin.Context) {
	var req request.FavoriteActionRequest
	// 先获取token
	token := c.Query("token")
	// 判断token是否为空
	if token == "" {
		c.JSON(http.StatusOK, response.Status{StatusCode: ERROR, StatusMsg: "token为空"})
		global.GVA_LOG.Info("token为空")
		return
	}
	// 解析token
	j := utils.NewJWT()
	_, err := j.ParseToken(token)
	if err != nil {
		// 过期中断
		if err == utils.TokenExpired {
			c.JSON(http.StatusOK, response.Status{StatusCode: ERROR, StatusMsg: "token过期"})
			global.GVA_LOG.Error("token过期", zap.Error(err))
			return
		}
		// 其他错误
		c.JSON(http.StatusOK, response.Status{StatusCode: ERROR, StatusMsg: err.Error()})
		global.GVA_LOG.Error("token解析失败", zap.Error(err))
		return
	}
	req.Token = token
	// 获取视频id
	video_id := c.Query("video_id")
	req.Video, _ = utils.StringToUint(video_id)
	// 然后获取action_type
	action_type := c.Query("action_type")
	req.Action_type, _ = utils.StringToUint(action_type)

	// 判断action_type是否为空
	if action_type == "" {
		c.JSON(http.StatusOK, response.Status{StatusCode: ERROR, StatusMsg: "action_type为空"})
		global.GVA_LOG.Info("action_type为空")
		return
	}
	// 判断action_type如果为1，则点赞，如果为2，则取消点赞
	if req.Action_type == uint(action_type_like) {
		// 点赞
		err := feedService.LikeVideo(req.Token, req.Video)
		if err != nil {
			global.GVA_LOG.Error("[FavoriteAction] 点赞失败", zap.Error(err))
			return
		}
		c.JSON(http.StatusOK, response.Status{StatusCode: SUCCESS, StatusMsg: "点赞成功"})
	}
	if req.Action_type == uint(action_type_unlike) {
		// 取消点赞
		err := feedService.UnLikeVideo(req.Token, req.Video)
		if err != nil {
			global.GVA_LOG.Error("[FavoriteAction] 取消点赞失败", zap.Error(err))
			return
		}
		c.JSON(http.StatusOK, response.Status{StatusCode: SUCCESS, StatusMsg: "取消点赞成功"})
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
