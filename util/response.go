package util

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type UserLoginResponse struct {
	Response
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}
