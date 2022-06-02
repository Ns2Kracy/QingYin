package service

import (
	"QingYin/global"
	model "QingYin/model/system"
	"QingYin/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

//针对接口:
// /douyin/user/register
// /douyin/user/login
// /douyin/user

type UserService struct{}

//用户注册
func (userService *UserService) Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	u.Password = utils.BcryptHash(u.Password)
	err = global.GVA_DB.Create(&u).Error
	return err, u
}

//用户登录
func (userService *UserService) Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	if nil == global.GVA_DB {
		return fmt.Errorf("db not init"), nil
	}
	var user model.SysUser
	err = global.GVA_DB.Where("username = ?", u.Username).First(&user).Error

	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return errors.New("用户密码错误"), nil
		}
	}
	return err, &user
}

//获取用户信息
func (userService *UserService) GetUserInfo(userId uint) (err error, user model.SysUser) {
	var reUser model.SysUser
	err = global.GVA_DB.Where("id = ?", userId).First(&reUser).Error
	if err != nil {
		return err, reUser
	}
	return err, reUser
}
