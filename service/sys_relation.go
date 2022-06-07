package service

import (
	"QingYin/global"
	model "QingYin/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RelationService struct{}

type userFocus struct {
	SysUserID uint
	FocuID    uint
}

//关注操作
func (*RelationService) Focus(UserId uint, ToUserID uint) error {
	Crerr := global.GVA_DB.Table("user_focus").Create(&userFocus{SysUserID: UserId, FocuID: ToUserID}).Error
	if Crerr != nil {
		global.GVA_LOG.Error("Create Focus relation Failed", zap.Error(Crerr))
		return Crerr
	}

	UpMainErr := global.GVA_DB.Model(&model.SysUser{}).Where("id = ?", UserId).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error
	if UpMainErr != nil {
		global.GVA_LOG.Error("Update Main Focus relation Failed", zap.Error(UpMainErr))
		return UpMainErr
	}

	UpSlavErr := global.GVA_DB.Model(&model.SysUser{}).Where("id = ?", ToUserID).Update("follower_count", gorm.Expr("follow_count + ?", 1)).Error
	if UpSlavErr != nil {
		global.GVA_LOG.Error("Update Slave Focus relation Failed", zap.Error(UpSlavErr))
		return UpSlavErr
	}

	return nil
}

//取消关注操作
func (*RelationService) UnFocus(UserId uint, ToUserID uint) error {
	err := global.GVA_DB.Table("user_focus").Delete(&userFocus{}, "sys_user_id = ? and focu_id = ?", UserId, ToUserID).Error
	if err != nil {
		global.GVA_LOG.Error("Cancel Focus relation Failed", zap.Error(err))
		return err
	}

	UpMainErr := global.GVA_DB.Model(&model.SysUser{}).Where("id = ?", UserId).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error
	if UpMainErr != nil {
		global.GVA_LOG.Error("Update Main Focus relation Failed", zap.Error(UpMainErr))
		return UpMainErr
	}

	UpSlavErr := global.GVA_DB.Model(&model.SysUser{}).Where("id = ?", ToUserID).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error
	if UpSlavErr != nil {
		global.GVA_LOG.Error("Update Slave Focus relation Failed", zap.Error(UpSlavErr))
		return UpSlavErr
	}

	return nil
}

//获取关注列表操作:我关注了谁
func (*RelationService) GetFollowList(UserID uint) (error, []model.SysUser) {
	var userList []model.SysUser
	err := global.GVA_DB.Where("id in (?)", global.GVA_DB.Table("user_focus").Select("focu_id").Where("sys_user_id = ?", UserID)).Find(&userList).Error
	if err != nil {
		global.GVA_LOG.Error("Search Focus relation Failed", zap.Error(err))
		return err, nil
	}
	return err, userList
}

//获取粉丝列表操作:谁关注了我
func (*RelationService) GetFollowerList(UserID uint) (error, []model.SysUser) {
	var userList []model.SysUser
	err := global.GVA_DB.Where("id in (?)", global.GVA_DB.Table("user_focus").Select("sys_user_id").Where("focu_id = ?", UserID)).Find(&userList).Error
	if err != nil {
		global.GVA_LOG.Error("Search Fans relation Failed", zap.Error(err))
		return err, nil
	}
	return err, userList
}
