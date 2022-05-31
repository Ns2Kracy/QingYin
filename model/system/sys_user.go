package model

import (
	"QingYin/global"
)

//用户
type SysUser struct {
	global.GVA_MODEL
	Username string `json:"userName" gorm:"comment:用户登录名"` //用户名
	Password string `json:"-"  gorm:"comment:用户登录密码"`      //用户密码
}
