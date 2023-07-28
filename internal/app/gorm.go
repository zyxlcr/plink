package app

import (
	"chatcser/config"
	"chatcser/pkg/friend"
	"chatcser/pkg/notification"
	"chatcser/pkg/user"
	"os"

	"github.com/tangpanqing/aorm"
	"github.com/tangpanqing/aorm/base"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch config.GVA_CONFIG.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}
func Aorm() *base.Db {
	switch config.GVA_CONFIG.DbType {
	case "mysql":
		return AormMysql()
	default:
		return AormMysql()
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := db.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(
		// 系统模块表
		user.BaseUser{},
		friend.Friend{},
		user.UserInfo{},
		notification.Notification{},

	//app.App{}, // app表注册
	)
	if err != nil {
		os.Exit(0)
	}
}

func RegisterTablesAorm(db *base.Db, isMg bool) {
	var arr []any
	u := &user.BaseUserAorm{}
	//f := &friend.Friend{}
	//uInfo := &user.UserInfo{}
	//ntf := &notification.Notification{}

	arr = append(arr, u)
	//arr = append(arr, f)
	//arr = append(arr, uInfo)
	//arr = append(arr, ntf)
	//arr = append(arr, u)

	//保存实例
	aorm.Store(arr...)

	if isMg {
		aorm.Migrator(db).AutoMigrate(
			arr...,
		)
	}

}
