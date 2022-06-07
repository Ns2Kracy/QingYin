package service

import (
	"QingYin/global"
	model "QingYin/model/system"

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
