package model

import "QingYin/global"

type SysComment struct {
	global.GVA_MODEL
	Content    string `gorm:"comment:评论内容;default:NULL"`
	UserRefer  uint   //用户ID
	VideoRefer uint   //视频ID
}

func (SysComment) TableName() string {
	return "sys_comments"
}
