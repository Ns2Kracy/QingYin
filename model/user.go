package model

// User 创建一个抖音用户
type User struct {
	ID       int    `json:"id,omitempty" gorm:"primary_key" `
	Nickname string `json:"nickname,omitempty" gorm:"type:varchar(255);not null"`
	Password string `json:"password,omitempty" gorm:"type:varchar(255);not null"`
}

type RegisterUser struct {
	NickName string `json:"nick_name,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginUser struct {
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
}
