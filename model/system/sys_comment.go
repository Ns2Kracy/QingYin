package model

type UserCommentVideo struct {
	SysUserID  uint   `gorm:"primary_key;autoIncrement:false"`
	SysVideoID uint   `gorm:"primary_key;autoIncrement:false"`
	Content    string `gorm:"comment:评论内容;default:NULL"`
}

type UserCommentVideoModel struct {
	UserId  uint   `json:"user_id" gorm:"comment:用户ID"`
	VideoId uint   `json:"video_id" gorm:"comment:视频ID"`
	Content string `json:"content" gorm:"comment:评论内容;default:NULL"`
}

func (UserCommentVideo) TableName() string {
	return "user_comment_videos"
}

func (UserCommentVideoModel) TableName() string {
	return "user_comment_videos"
}
