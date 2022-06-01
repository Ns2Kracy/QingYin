package model

type UserCommentVideo struct {
	SysUserID  uint   `gorm:"primary_key;autoIncrement:false"`
	SysVideoID uint   `gorm:"primary_key;autoIncrement:false"`
	Content    string `gorm:"comment:评论内容"`
}

func (UserCommentVideo) TableName() string {
	return "user_comment_videos"
}
