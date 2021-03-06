package v1

import (
	"QingYin/global"
	model "QingYin/model/system"
	"QingYin/model/system/request"
	"QingYin/model/system/response"
	"QingYin/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type basicApi struct{}

//响应状态
const (
	ERROR   = 7
	SUCCESS = 0
)

//相当于controller层,调用service层方法实现业务逻辑

//生成用户返回信息:::::仅限内部调用
func userInfoResponse(userId uint) (error, response.User) {
	//获取用户信息
	err, userInfo := userService.GetUserInfo(userId)
	if err != nil {
		global.GVA_LOG.Error("获取用户信息失败!", zap.Error(err))
		return err, response.User{}
	}
	//用户返回信息
	userRet := response.User{
		ID:            userInfo.ID,
		Username:      userInfo.Username,
		FollowCount:   int64(userInfo.FollowCount),
		FollowerCount: int64(userInfo.FollowerCount),
		IsFollow:      false}
	return nil, userRet
}

//获取视频流信息>>>>未测试<<<<<测试通过
func (b *basicApi) Feed(c *gin.Context) {
	var feedReq request.FeedRequest
	//绑定请求参数
	_ = c.ShouldBind(&feedReq)

	//默认当前时间
	if feedReq.LatestTime == "" {
		feedReq.LatestTime = time.Now().Format("2006-01-02 15:04:05")
	} else {
		//时间戳转换
		value, _ := strconv.ParseInt(feedReq.LatestTime, 10, 64)
		feedReq.LatestTime = time.Unix(value, 0).Format("2006-01-02 15:04:05")
	}

	//获取视频信息
	videoList, getErr := feedService.GetVideoFeed(feedReq.LatestTime)
	if getErr != nil {
		global.GVA_LOG.Error("获取视频流失败!", zap.Error(getErr))
		status := response.Status{StatusCode: ERROR, StatusMsg: "获取视频流失败"}
		c.JSON(http.StatusOK, response.FeedResponse{Status: status, NextTime: 0, VideoList: nil})
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
			c.JSON(http.StatusOK, response.FeedResponse{Status: status, NextTime: 0, VideoList: nil})
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
			IsFavorite:    false}

		//追加视频信息
		videos = append(videos, v)
	}

	//成功返回
	status := response.Status{StatusCode: SUCCESS, StatusMsg: "Feed流获取成功"}
	c.JSON(http.StatusOK, response.FeedResponse{Status: status, NextTime: int(videoList[0].CreatedAt.Unix()), VideoList: videos})
}

//用户信息获取:::::::::::未实现is_follow功能
func (b *basicApi) UserInfo(c *gin.Context) {
	uid := utils.GetUserID(c)
	if err, userInfo := userService.GetUserInfo(uid); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		status := response.Status{StatusCode: ERROR, StatusMsg: "获取失败!"}
		user := response.User{ID: 0, Username: "", FollowCount: 0, FollowerCount: 0, IsFollow: false}
		c.JSON(http.StatusOK, response.UserInfoResponse{Status: status, User: user})
	} else {
		status := response.Status{StatusCode: SUCCESS, StatusMsg: "获取成功!"}
		user := response.User{ID: userInfo.ID, Username: userInfo.Username, FollowCount: int64(userInfo.FollowCount), FollowerCount: int64(userInfo.FollowerCount), IsFollow: false}
		c.JSON(http.StatusOK, response.UserInfoResponse{Status: status, User: user})
	}
}

//用户登录返回d以及Token
func (b *basicApi) Login(c *gin.Context) {
	var l request.LoginRequest
	_ = c.ShouldBind(&l)
	u := &model.SysUser{Username: l.UserName, Password: l.UserPasswd}

	if err, user := userService.Login(u); err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		status := response.Status{StatusCode: ERROR, StatusMsg: "登陆失败"}
		c.JSON(http.StatusOK, response.LoginResponse{Status: status, ID: 0, Token: ""})
	} else {
		err, token := b.tokenIssue(c, *user) //签发Token
		if err != nil {
			status := response.Status{StatusCode: ERROR, StatusMsg: "Token获取失败"}
			c.JSON(http.StatusOK, response.RegisterResponse{Status: status, ID: 0, Token: ""})
		}
		status := response.Status{StatusCode: SUCCESS, StatusMsg: "登陆成功"}
		c.JSON(http.StatusOK, response.LoginResponse{Status: status, ID: user.ID, Token: token})
	}
}

//用户注册返回id以及Token
func (b *basicApi) Register(c *gin.Context) {
	var r request.RegisterRequest
	_ = c.ShouldBind(&r)
	user := &model.SysUser{Username: r.UserName, Password: r.UserPasswd}
	err, userRet := userService.Register(*user)

	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		status := response.Status{StatusCode: ERROR, StatusMsg: "注册失败"}
		c.JSON(http.StatusOK, response.RegisterResponse{Status: status, ID: 0, Token: ""})
	} else {
		err, token := b.tokenIssue(c, userRet) //签发Token
		if err != nil {
			status := response.Status{StatusCode: ERROR, StatusMsg: "Token获取失败"}
			c.JSON(http.StatusOK, response.RegisterResponse{Status: status, ID: 0, Token: ""})
		}
		status := response.Status{StatusCode: SUCCESS, StatusMsg: "注册成功"}
		c.JSON(http.StatusOK, response.RegisterResponse{Status: status, ID: userRet.ID, Token: token})
	}

}

//签发Token
func (b *basicApi) tokenIssue(c *gin.Context, user model.SysUser) (error, string) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(request.BaseClaims{
		UserId:   user.ID,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		return err, ""
	}
	return err, token
}

//视频投稿返回是否成功
func (b *basicApi) Publish(c *gin.Context) {
	var videoReq request.PublishActionRequest
	videoReq.Title = c.Request.FormValue("title")
	videoReq.Token = c.Request.FormValue("token")
	var sysVideo model.SysVideo

	//确定视频标题信息
	sysVideo.VideoTitle = videoReq.Title

	//从Token中确定作者信息
	j := utils.NewJWT()
	claims, Parserr := j.ParseToken(videoReq.Token)
	if Parserr != nil {
		global.GVA_LOG.Error("解析Token失败", zap.Error(Parserr))
		status := response.Status{StatusCode: ERROR, StatusMsg: "解析Token失败"}
		c.JSON(http.StatusOK, response.PublishActionResponse{Status: status})
		return
	}
	sysVideo.UserRefer = claims.UserId //赋予作者

	//上传视频并保存信息至数据库
	_, header, err := c.Request.FormFile("data")
	// c.Get("claims")	//保留节目
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		status := response.Status{StatusCode: ERROR, StatusMsg: "上传文件失败"}
		c.JSON(http.StatusOK, response.PublishActionResponse{Status: status})
		return
	}
	pErr := publishService.Action(header, &sysVideo) //考虑用videoReq.Data()替代
	if pErr != nil {
		global.GVA_LOG.Error("上传文件失败!", zap.Error(err))
		status := response.Status{StatusCode: ERROR, StatusMsg: "上传文件失败"}
		c.JSON(http.StatusOK, response.PublishActionResponse{Status: status})
		return
	}
	status := response.Status{StatusCode: SUCCESS, StatusMsg: "上传文件成功"}
	c.JSON(http.StatusOK, response.PublishActionResponse{Status: status})
}

//获取发布列表:::::::::::未实现is_follow功能>>>>未测试<<<<<<测试通过
func (b *basicApi) PublishList(c *gin.Context) {
	var listReq request.PublishListRequest
	_ = c.ShouldBind(&listReq)

	videoList, getErr := publishService.GetPublishList(listReq.UserID)

	//获取用户信息
	err, userInfo := userService.GetUserInfo(listReq.UserID)
	if err != nil {
		global.GVA_LOG.Error("获取用户信息失败!", zap.Error(getErr))
		status := response.Status{StatusCode: ERROR, StatusMsg: "获取用户信息失败"}
		c.JSON(http.StatusOK, response.PublishListResponse{Status: status, VideoList: nil})
		return
	}

	//用户返回信息
	userRet := response.User{
		ID:            userInfo.ID,
		Username:      userInfo.Username,
		FollowCount:   int64(userInfo.FollowCount),
		FollowerCount: int64(userInfo.FollowerCount),
		IsFollow:      false}

	//获取失败直接返回
	if getErr != nil {
		global.GVA_LOG.Error("获取成功失败!", zap.Error(getErr))
		status := response.Status{StatusCode: ERROR, StatusMsg: "获取成功失败"}
		c.JSON(http.StatusOK, response.PublishListResponse{Status: status, VideoList: nil})
		return
	}

	//视频返回信息
	var videos []response.Video
	for _, video := range videoList {
		v := response.Video{
			ID:            video.ID,
			Title:         video.VideoTitle,
			Author:        userRet,
			PlayURL:       video.PlayURL,
			CoverURL:      video.CoverURL,
			FavoriteCount: int(video.FavoriteCount),
			CommentCount:  int(video.CommentCount),
			IsFavorite:    false}

		//追加视频信息
		videos = append(videos, v)
	}

	status := response.Status{StatusCode: SUCCESS, StatusMsg: "获取成功"}
	c.JSON(http.StatusOK, response.PublishListResponse{Status: status, VideoList: videos})
}
