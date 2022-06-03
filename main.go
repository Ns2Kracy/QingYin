package main

import (
	"QingYin/core"
	"QingYin/global"
	"QingYin/initialize"

	"go.uber.org/zap"
)

func main() {
	global.GVA_VP = core.Viper() //初始化Viper
	global.GVA_LOG = core.Zap()  //初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.GormMysql() //gorm连接MySQL数据库
	if global.GVA_DB != nil {
		//创建数据库表
		initialize.RegisterTables(global.GVA_DB)
		//延迟关闭数据库
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

	//运行服务
	core.RunServer()
}
