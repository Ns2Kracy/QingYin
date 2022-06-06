package service

import (
	"QingYin/global"
	model "QingYin/model/system"
	"QingYin/utils"

	"go.uber.org/zap"
)

//针对接口:
// /douyin/feed

type FeedService struct{}

const (
	MAX_VIDEOS = 30
)

//获取指定时间戳之前的视频信息
func (*FeedService) GetVideoFeed(last_time string) ([]model.SysVideo, error) {
	var videos []model.SysVideo
	err := global.GVA_DB.Where("created_at < ?", last_time).Limit(MAX_VIDEOS).Order("created_at desc").Find(&videos).Error
	//select * from sys_videos where created_at < last_time order by created_at limit MAX_VIDEOS
	if err != nil {
		global.GVA_LOG.Error("查询视频信息失败", zap.Error(err))
		return videos, err
	}
	return videos, nil
}

func (s *FeedService) LikeVideo(UserId string, VedioId uint) error {
	// 保存视频点赞信息
	var video model.SysVideo
	err := global.GVA_DB.First(&video, VedioId).Error
	if err != nil {
		global.GVA_LOG.Error("查询视频信息失败", zap.Error(err))
		return err
	}
	video.FavoriteCount++
	//更新视频isFavorite字段
	err = global.GVA_DB.Save(&video).Update("isFavorite", true).Error
	if err != nil {
		global.GVA_LOG.Error("保存点赞信息失败", zap.Error(err))
		return err
	}
	// 保存用户点赞信息到user_favorite_videos表
	var userFavoriteVideo model.UserFavoriteVideo
	userFavoriteVideo.UserID, _ = utils.StringToUint(UserId)
	userFavoriteVideo.VideoID = VedioId
	err = global.GVA_DB.Save(userFavoriteVideo).Error
	if err != nil {
		global.GVA_LOG.Error("保存用户点赞信息失败", zap.Error(err))
		return err
	}

	return nil
}

func (s *FeedService) UnLikeVideo(UserId string, VedioId uint) error {
	// 取消视频点赞信息
	var video model.SysVideo
	err := global.GVA_DB.First(&video, VedioId).Error
	if err != nil {
		global.GVA_LOG.Error("查询视频信息失败", zap.Error(err))
		return err
	}
	video.FavoriteCount--
	//更新视频isFavorite字段
	err = global.GVA_DB.Save(&video).Update("isFavorite", false).Error
	if err != nil {
		global.GVA_LOG.Error("取消点赞信息失败", zap.Error(err))
		return err
	}
	// 取消用户点赞信息到user_favorite_videos表
	var userFavoriteVideo model.UserFavoriteVideo
	userFavoriteVideo.UserID, _ = utils.StringToUint(UserId)
	userFavoriteVideo.VideoID = VedioId
	err = global.GVA_DB.Delete(userFavoriteVideo).Error
	if err != nil {
		global.GVA_LOG.Error("取消用户点赞信息失败", zap.Error(err))
		return err
	}

	return nil
}
