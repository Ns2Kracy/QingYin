package initialize

import (
	"QingYin/global"
	model "QingYin/model/system"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//初始化MySQL数据库
func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	//数据库名为空返回nil
	if m.Dbname == "" {
		return nil
	}
	//加载MySQL自定义配置
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), //数据源
		DefaultStringSize:         191,     //string类型默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		return db
	}
}

// 自动建立表结构
func RegisterTables(db *gorm.DB) {
	//这一步很重要
	db.SetupJoinTable(&model.SysUser{}, "Comments", &model.UserCommentVideo{})
	err := db.AutoMigrate(
		&model.SysUser{},
		&model.SysVideo{},
		&model.UserCommentVideo{},
	)

	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
