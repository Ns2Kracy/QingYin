package request

//用户请求

//注册请求
type RegisterRequest struct {
	UserName   string //注册用户名
	UserPasswd string //注册密码
}

//登录请求
type LoginRequest struct {
	UserName   string //登录用户名
	UserPasswd string //登陆密码
}

//用户信息请求
type UserInfoRequest struct {
	ID    uint   //用户ID
	Token string //用户鉴权Token
}
