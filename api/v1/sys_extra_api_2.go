package v1

import (
	"QingYin/global"
	"QingYin/model/system/request"
	"QingYin/model/system/response"
	"QingYin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type extraApi_2 struct{}

const (
	focus_action   = 1 //关注
	unfocus_action = 2 //取消关注
)

//关注操作>>>>>>>>>>未测试<<<<<<<<<<<postman测试通过
func (*extraApi_2) Relation(c *gin.Context) {
	var foucusActionReq request.FocusActionRequest
	_ = c.ShouldBind(&foucusActionReq)

	//获取主体用户信息
	UserId := utils.GetUserID(c)

	//判断操作类型
	switch foucusActionReq.ActionType {
	case focus_action:
		focusErr := relationService.Focus(UserId, foucusActionReq.ToUserID)
		if focusErr != nil {
			global.GVA_LOG.Error("Focus Failed", zap.Error(focusErr))
			status := response.Status{StatusCode: ERROR, StatusMsg: "关注失败"}
			c.JSON(http.StatusOK, response.FocusActionResponse{Status: status})
		}
		status := response.Status{StatusCode: SUCCESS, StatusMsg: "关注成功"}
		c.JSON(http.StatusOK, response.FocusActionResponse{Status: status})
	case unfocus_action:
		unfocusErr := relationService.UnFocus(UserId, foucusActionReq.ToUserID)
		if unfocusErr != nil {
			global.GVA_LOG.Error("Cancel Focus Failed", zap.Error(unfocusErr))
			status := response.Status{StatusCode: ERROR, StatusMsg: "取消关注失败"}
			c.JSON(http.StatusOK, response.FocusActionResponse{Status: status})
		}
		status := response.Status{StatusCode: SUCCESS, StatusMsg: "取消关注成功"}
		c.JSON(http.StatusOK, response.FocusActionResponse{Status: status})
	}
}

// 关注列表>>>>>>>>>>未测试<<<<<<<<<postman测试通过
func (*extraApi_2) FollowList(c *gin.Context) {
	var followListReq request.FollowListRequest
	_ = c.ShouldBind(&followListReq)

	err, userList := relationService.GetFollowList(followListReq.UserID)
	if err != nil {
		global.GVA_LOG.Error("Get Focus List Failed", zap.Error(err))
		status := response.Status{StatusCode: ERROR, StatusMsg: "获取失败"}
		c.JSON(http.StatusOK, response.FollowListResponse{Status: status, UserList: nil})
		return
	}

	var users []response.User
	for _, user := range userList {
		u := response.User{
			ID:            user.ID,
			Username:      user.Username,
			FollowCount:   int64(user.FollowCount),
			FollowerCount: int64(user.FollowerCount),
			IsFollow:      true,
		}
		users = append(users, u)
	}

	status := response.Status{StatusCode: SUCCESS, StatusMsg: "获取成功"}
	c.JSON(http.StatusOK, response.FollowListResponse{Status: status, UserList: users})
}

// 粉丝列表>>>>>>>>>>未测试<<<<<<<<<postman测试通过
func (*extraApi_2) FollowerList(c *gin.Context) {
	var followerListReq request.FollowerListRequest
	_ = c.ShouldBind(&followerListReq)

	err, userList := relationService.GetFollowerList(followerListReq.UserID)
	if err != nil {
		global.GVA_LOG.Error("Get Focus List Failed", zap.Error(err))
		status := response.Status{StatusCode: ERROR, StatusMsg: "获取失败"}
		c.JSON(http.StatusOK, response.FollowerListResponse{Status: status, UserList: nil})
		return
	}

	var users []response.User
	for _, user := range userList {
		u := response.User{
			ID:            user.ID,
			Username:      user.Username,
			FollowCount:   int64(user.FollowCount),
			FollowerCount: int64(user.FollowerCount),
			IsFollow:      false,
		}
		users = append(users, u)
	}

	status := response.Status{StatusCode: SUCCESS, StatusMsg: "获取成功"}
	c.JSON(http.StatusOK, response.FollowerListResponse{Status: status, UserList: users})
}
