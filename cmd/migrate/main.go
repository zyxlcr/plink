package main

import (
	"chatcser/config"
	"chatcser/internal/app"
)

func main() {
	// 初始化配置
	config.GVA_VP = app.Viper()
	config.GVA_DB = app.Gorm() // gorm连接数据库
	if config.GVA_DB != nil {
		print("!= nil")
		app.RegisterTables(config.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := config.GVA_DB.DB()
		defer db.Close()
	}
}
