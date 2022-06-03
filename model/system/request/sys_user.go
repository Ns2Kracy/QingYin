package request

//用户请求

//注册请求
type RegisterRequest struct {
	UserName   string `form:"username"` //注册用户名
	UserPasswd string `form:"password"` //注册密码
}

//登录请求
type LoginRequest struct {
	UserName   string `form:"username"` //登录用户名
	UserPasswd string `form:"password"` //登陆密码
}

//用户信息请求
type UserInfoRequest struct {
	ID    uint   `form:"user_id"` //用户ID
	Token string `form:"token"`   //用户鉴权Token
}
