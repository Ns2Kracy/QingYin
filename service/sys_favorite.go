package service

import (
	"QingYin/global"
	model "QingYin/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FavoriteService struct{}

type userLike struct {
	SysUserID  uint
	SysVideoID uint
}

//点赞操作
func (*FavoriteService) LikeAction(UserId uint, VideoId uint) error {
	err := global.GVA_DB.Table("user_like_videos").Create(&userLike{SysUserID: UserId, SysVideoID: VideoId}).Error
	if err != nil {
		global.GVA_LOG.Error("Like Action Save Failed", zap.Error(err))
		return err
	}

	upErr := global.GVA_DB.Model(&model.SysVideo{}).Where("id = ?", VideoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	if upErr != nil {
		global.GVA_LOG.Error("Update like relation Failed", zap.Error(upErr))
		return upErr
	}
	return nil
}

//取消点赞操作
func (*FavoriteService) CancelLikeAction(UserId uint, VideoId uint) error {
	err := global.GVA_DB.Table("user_like_videos").Delete(&userLike{}, "sys_user_id = ? and sys_video_id = ?", UserId, VideoId).Error
	if err != nil {
		global.GVA_LOG.Error("Cancel Like Action Save Failed", zap.Error(err))
		return err
	}

	upErr := global.GVA_DB.Model(&model.SysVideo{}).Where("id = ?", VideoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
	if upErr != nil {
		global.GVA_LOG.Error("Update cancel like relation Failed", zap.Error(upErr))
		return upErr
	}
	return nil
}

//获取点赞列表操作
func (*FavoriteService) GetLikeList(UserId uint) (error, []model.SysVideo) {
	var videos []model.SysVideo
	err := global.GVA_DB.Where("id in (?)", global.GVA_DB.Table("user_like_videos").Select("sys_video_id").Where("sys_user_id = ?", UserId)).Find(&videos).Error
	if err != nil {
		global.GVA_LOG.Error("Get Like List Failed", zap.Error(err))
		return err, nil
	}
	return nil, videos

}
