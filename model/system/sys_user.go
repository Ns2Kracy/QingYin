package model

import (
	"QingYin/global"
)

//用户
type SysUser struct {
	global.GVA_MODEL
	Username      string     `json:"userName" gorm:"comment:用户登录名"`                          //用户名
	Password      string     `json:"-"  gorm:"comment:用户登录密码"`                               //用户密码
	FollowerCount uint       `json:"follower_count" gorm:"comment:粉丝总数"`                     //粉丝总数
	FollowCount   uint       `json:"follow_count" gorm:"comment:关注总数"`                       //关注总数
	SysVideos     []SysVideo `gorm:"foreignKey:UserRefer"`                                   //发布联系
	Focus         []*SysUser `gorm:"many2many:user_focus"`                                   //关注联系
	Comments      []SysVideo `gorm:"many2many:user_comment_videos"`                          //评论关系
	Likes         []SysVideo `gorm:"many2many:user_like_videos;ForeignKey:id;References:id"` //点赞关系
}

func (SysUser) TableName() string {
	return "sys_users"
}
