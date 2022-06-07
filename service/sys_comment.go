package service

import (
	"QingYin/global"
	model "QingYin/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CommentService struct{}

//发布评论操作>>>>>>>>>未测试>>>>>>>>优化建议:使用事务进行提交
func (*CommentService) CommentAction(comment model.SysComment) (error, model.SysComment) {
	err := global.GVA_DB.Create(&comment).Error
	if err != nil {
		global.GVA_LOG.Error("Create comment Failed", zap.Error(err))
		return err, model.SysComment{}
	}

	//更新视频数据表
	upErr := global.GVA_DB.Model(&model.SysVideo{}).Where("id = ?", comment.VideoRefer).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	if upErr != nil {
		global.GVA_LOG.Error("Update comment count Failed", zap.Error(upErr))
		return upErr, model.SysComment{}
	}

	return nil, comment
}

//删除评论操作>>>>>>>>>未测试>>>>>>>>优化建议:使用事务进行提交
func (*CommentService) DeleteCommentAction(comment model.SysComment) error {
	err := global.GVA_DB.Delete(&model.SysComment{}, comment.ID).Error
	if err != nil {
		global.GVA_LOG.Error("Delete comment Failed", zap.Error(err))
		return err
	}

	//更新视频数据表
	upErr := global.GVA_DB.Model(&model.SysVideo{}).Where("id = ?", comment.VideoRefer).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error
	if upErr != nil {
		global.GVA_LOG.Error("Update comment count Failed", zap.Error(upErr))
		return upErr
	}
	return nil
}

//获取评论列表>>>>>>>>>未测试
func (*CommentService) GetCommentList(videoID uint) (error, []model.SysComment) {
	var comments []model.SysComment
	err := global.GVA_DB.Where("video_refer = ?", videoID).Order("created_at desc").Find(&comments).Error
	if err != nil {
		global.GVA_LOG.Error("Get comment list Failed", zap.Error(err))
		return err, nil
	}
	return nil, comments
}
