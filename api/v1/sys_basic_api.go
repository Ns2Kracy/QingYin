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

type basicApi struct{}

//响应状态
const (
	ERROR   = 7
	SUCCESS = 0
)

//相当于controller层,调用service层方法实现业务逻辑
func (b *basicApi) Feed(c *gin.Context) {
	status := response.Status{StatusCode: SUCCESS, StatusMsg: "访问视频Feed流成功"}
	c.JSON(http.StatusOK, response.FeedResponse{Status: status})
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
	// videoReq.Data = c.Request.FormFile("data")
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

func (b *basicApi) PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, "测试/publish/list/接口")
}
