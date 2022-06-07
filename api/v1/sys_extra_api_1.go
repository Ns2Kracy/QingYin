package v1

import (
	"QingYin/global"
	model "QingYin/model/system"
	"QingYin/model/system/request"
	"QingYin/model/system/response"
	"QingYin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type extraApi_1 struct{}

const (
	like_action   = 1 //点赞
	unlike_action = 2 //取消点赞
)

// 登录用户对视频的点赞和取消点赞操作
// 点赞
// 取消点赞
//>>>>>>未测试<<<<<<<<postman测试通过
func (*extraApi_1) FavoriteAction(c *gin.Context) {
	var favoriteActionReq request.LikeActionRequest
	_ = c.ShouldBind(&favoriteActionReq)

	//获取主体用户信息
	UserId := utils.GetUserID(c)

	//判断操作类型进行响应操作
	switch favoriteActionReq.ActionType {
	case like_action:
		likeErr := favoriteService.LikeAction(UserId, favoriteActionReq.VideoID)
		if likeErr != nil {
			global.GVA_LOG.Error("Like Action Failed", zap.Error(likeErr))
			status := response.Status{StatusCode: ERROR, StatusMsg: "点赞失败"}
			c.JSON(http.StatusOK, response.LikeActionResponse{Status: status})
		}
		status := response.Status{StatusCode: SUCCESS, StatusMsg: "点赞成功"}
		c.JSON(http.StatusOK, response.LikeActionResponse{Status: status})
	case unlike_action:
		unlikeErr := favoriteService.CancelLikeAction(UserId, favoriteActionReq.VideoID)
		if unlikeErr != nil {
			global.GVA_LOG.Error("Cancel Like Action Failed", zap.Error(unlikeErr))
			status := response.Status{StatusCode: ERROR, StatusMsg: "取消点赞失败"}
			c.JSON(http.StatusOK, response.LikeActionResponse{Status: status})
		}
		status := response.Status{StatusCode: SUCCESS, StatusMsg: "取消点赞成功"}
		c.JSON(http.StatusOK, response.LikeActionResponse{Status: status})
	}
}

// 获取点赞列表>>>>>>未测试>>>>>>>postman测试通过
func (*extraApi_1) FavoriteList(c *gin.Context) {
	var favoriteListReq request.LikeListRequest
	_ = c.ShouldBind(&favoriteListReq)

	err, videoList := favoriteService.GetLikeList(favoriteListReq.UserID)
	if err != nil {
		global.GVA_LOG.Error("Get like list failed", zap.Error(err))
		status := response.Status{StatusCode: ERROR, StatusMsg: "获取点赞列表失败"}
		c.JSON(http.StatusOK, response.LikeListResponse{Status: status, VideoList: nil})
		return
	}

	var videos []response.Video
	//生成视频返回信息
	for _, video := range videoList {
		//生成视频用户信息
		err, userRet := userInfoResponse(video.UserRefer)
		if err != nil {
			global.GVA_LOG.Error("获取用户信息失败!", zap.Error(err))
			status := response.Status{StatusCode: ERROR, StatusMsg: "获取用户信息失败"}
			c.JSON(http.StatusOK, response.LikeListResponse{Status: status, VideoList: nil})
			return
		}

		//生成视频返回信息
		v := response.Video{
			ID:            video.ID,
			Title:         video.VideoTitle,
			Author:        userRet,
			PlayURL:       video.PlayURL,
			CoverURL:      video.CoverURL,
			FavoriteCount: int(video.FavoriteCount),
			CommentCount:  int(video.CommentCount),
			IsFavorite:    true}

		//追加视频信息
		videos = append(videos, v)
	}

	status := response.Status{StatusCode: SUCCESS, StatusMsg: "获取点赞列表成功"}
	c.JSON(http.StatusOK, response.LikeListResponse{Status: status, VideoList: videos})
}

const (
	comment_action        = "1" //发布评论
	delete_comment_action = "2" //删除评论
)

// 评论操作>>>>>>>>>未测试<<<<<<<<<<postman测试通过
func (*extraApi_1) CommentAction(c *gin.Context) {
	action_type := c.Request.FormValue("action_type")

	//获取主体用户信息
	UserId := utils.GetUserID(c)
	err, userRet := userInfoResponse(UserId)
	if err != nil {
		global.GVA_LOG.Error("获取用户信息失败!", zap.Error(err))
		status := response.Status{StatusCode: ERROR, StatusMsg: "获取用户信息失败"}
		c.JSON(http.StatusOK, response.LikeListResponse{Status: status, VideoList: nil})
		return
	}
	switch action_type {
	case comment_action:
		var commentReq request.CommentActionRequest
		_ = c.ShouldBind(&commentReq)

		comment := &model.SysComment{Content: commentReq.CommentText, UserRefer: UserId, VideoRefer: commentReq.VideoID}

		err, commentRet := commentService.CommentAction(*comment)
		if err != nil {
			global.GVA_LOG.Error("Create comment Failed", zap.Error(err))
			status := response.Status{StatusCode: ERROR, StatusMsg: "评论失败"}
			c.JSON(http.StatusOK, response.CommentActionResponse{Status: status, Comment: response.Comment{}})
			return
		}
		status := response.Status{StatusCode: SUCCESS, StatusMsg: "评论成功"}
		c.JSON(http.StatusOK, response.CommentActionResponse{Status: status, Comment: response.Comment{ID: commentRet.ID, User: userRet, Content: commentReq.CommentText, CreateDate: commentRet.CreatedAt.Format("2006-01-02")}})
	case delete_comment_action:
		var delCommentReq request.DeleteCommentActionRequest
		_ = c.ShouldBind(&delCommentReq)

		comment := &model.SysComment{VideoRefer: delCommentReq.VideoID}
		comment.ID = delCommentReq.CommentID

		err := commentService.DeleteCommentAction(*comment)
		if err != nil {
			global.GVA_LOG.Error("Delete comment Failed", zap.Error(err))
			status := response.Status{StatusCode: ERROR, StatusMsg: "删除评论失败"}
			c.JSON(http.StatusOK, response.CommentActionResponse{Status: status, Comment: response.Comment{}})
			return
		}
		status := response.Status{StatusCode: SUCCESS, StatusMsg: "删除评论成功"}
		c.JSON(http.StatusOK, response.CommentActionResponse{Status: status, Comment: response.Comment{}})
	}
}

// 查看评论列表，按时间排序
// 获取评论列表>>>>>>>未测试<<<<<<<<<<postman测试通过
func (*extraApi_1) CommentList(c *gin.Context) {
	var commentListReq request.CommentListRequest
	_ = c.ShouldBind(&commentListReq)

	err, commentList := commentService.GetCommentList(commentListReq.VideoID)
	if err != nil {
		global.GVA_LOG.Error("Get comment list Failed", zap.Error(err))
		status := response.Status{StatusCode: ERROR, StatusMsg: "获取评论列表失败"}
		c.JSON(http.StatusOK, response.CommentListResponse{Status: status, CommentList: nil})
		return
	}

	var comments []response.Comment
	for _, comment := range commentList {
		err, userRet := userInfoResponse(comment.UserRefer)
		if err != nil {
			global.GVA_LOG.Error("获取用户信息失败!", zap.Error(err))
			status := response.Status{StatusCode: ERROR, StatusMsg: "获取用户信息失败"}
			c.JSON(http.StatusOK, response.LikeListResponse{Status: status, VideoList: nil})
			return
		}

		c := response.Comment{
			ID:         comment.ID,
			User:       userRet,
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("2006-01-02"),
		}

		comments = append(comments, c)
	}

	//成功返回
	status := response.Status{StatusCode: SUCCESS, StatusMsg: "获取评论列表成功"}
	c.JSON(http.StatusOK, response.CommentListResponse{Status: status, CommentList: comments})
}
