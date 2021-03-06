package model

import "QingYin/global"

type SysVideo struct {
	global.GVA_MODEL
	VideoTitle    string       `json:"title" gorm:"comment:视频标题"`                      //视频标题
	PlayURL       string       `json:"play_url" gorm:"comment:视频播放地址;default:NULL"`    //视频播放地址
	CoverURL      string       `json:"cover_url" gorm:"comment:视频封面地址;default:NULL"`   //视频封面地址
	FavoriteCount uint         `json:"favorite_count" gorm:"comment:视频点赞总数;default:0"` //视频点赞总数
	CommentCount  uint         `json:"comment_count" gorm:"comment:视频评论总数;default:0"`  //视频评论总数
	UserRefer     uint         //用户ID
	Comments      []SysComment `gorm:"foreignKey:VideoRefer;default:NULL"` //评论列表
}

//自定义表名
func (SysVideo) TableName() string {
	return "sys_videos"
}
