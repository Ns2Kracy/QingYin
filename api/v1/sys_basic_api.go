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

const (
	ERROR   = 7
	SUCCESS = 0
)

//相当于controller层,调用service层方法实现业务逻辑
func (b *basicApi) Feed(c *gin.Context) {
	status := response.Status{StatusCode: 1, StatusMsg: "访问视频Feed流成功"}
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

//用户登录
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

//用户注册
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

func (b *basicApi) Publish(c *gin.Context) {
	c.JSON(http.StatusOK, "测试/publish/action/接口")
}

func (b *basicApi) PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, "测试/publish/list/接口")
}
