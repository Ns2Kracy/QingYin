package database

import (
	"DouYin/config"
	"DouYin/model"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Engine() *gorm.DB {
	initConfig := config.InitConfig()
	if initConfig == nil {
		return nil
	}

	database := initConfig.MySQL
	dataSourceName := database.User + ":" + database.Pwd + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.Database + "?charset=utf8"
	// 连接数据库
	engine, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err == nil {
		engine.AutoMigrate(
			&model.User{},
		)
	}
	return nil
}
