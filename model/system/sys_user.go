package model

import (
	"QingYin/global"
)

//用户
type SysUser struct {
	global.GVA_MODEL
	Username      string       `json:"userName" gorm:"comment:用户登录名;uniqueIndex"`                           //用户名
	Password      string       `json:"-"  gorm:"comment:用户登录密码"`                                            //用户密码
	FollowerCount uint         `json:"follower_count" gorm:"comment:粉丝总数;default:0"`                        //粉丝总数
	FollowCount   uint         `json:"follow_count" gorm:"comment:关注总数;default:0"`                          //关注总数
	SysVideos     []SysVideo   `gorm:"foreignKey:UserRefer;default:NULL"`                                   //发布联系
	Focus         []*SysUser   `gorm:"many2many:user_focus;default:NULL"`                                   //关注联系:自引用
	Comments      []SysComment `gorm:"foreignKey:UserRefer;default:NULL"`                                   //评论关系
	Likes         []SysVideo   `gorm:"many2many:user_like_videos;ForeignKey:id;References:id;default:NULL"` //点赞关系
}

func (SysUser) TableName() string {
	return "sys_users"
}
