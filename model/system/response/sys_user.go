package response

//用户接口响应

//用户信息接口响应
type UserInfoResponse struct {
	Status      //状态信息
	User   User `json:"user"` //用户信息
}

//用户注册接口响应
type RegisterResponse struct {
	Status        // 状态信息
	ID     uint   `json:"user_id"` //用户ID
	Token  string `json:"token"`   //用户鉴权Token
}

//用户登录接口响应
type LoginResponse struct {
	Status        // 状态信息
	ID     uint   `json:"user_id"` //用户ID
	Token  string `json:"token"`   //用户鉴权Token
}
