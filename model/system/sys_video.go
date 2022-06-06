package model

import "QingYin/global"

type SysVideo struct {
	global.GVA_MODEL
	Author        SysUser `json:"author"`
	VideoTitle    string  `json:"title" gorm:"comment:视频标题"`                      //视频标题
	PlayURL       string  `json:"play_url" gorm:"comment:视频播放地址"`                 //视频播放地址
	CoverURL      string  `json:"cover_url" gorm:"comment:视频封面地址"`                //视频封面地址
	FavoriteCount uint    `json:"favorite_count" gorm:"comment:视频点赞总数;default:0"` //视频点赞总数
	CommentCount  uint    `json:"comment_count" gorm:"comment:视频评论总数;default:0"`  //视频评论总数
	IsFavorite    bool    `json:"is_favorite" gorm:"comment:视频点赞;default:false"`  //视频是否点赞
	UserRefer     uint    //用户发布联系
}

type UserFavoriteVideo struct {
	VideoID uint `json:"video_id" gorm:"comment:视频ID"`
	UserID  uint `json:"user_id" gorm:"comment:用户ID"`
}

//自定义表名
func (SysVideo) TableName() string {
	return "sys_videos"
}

func (UserFavoriteVideo) TableName() string {
	return "user_like_videos"
}
