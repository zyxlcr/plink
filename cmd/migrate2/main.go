package main

import (
	"chatcser/config"
	"chatcser/internal/app"
)

func main() {
	// 初始化配置
	config.GVA_VP = app.Viper()
	config.GVA_AORM = app.Aorm() // gorm连接数据库
	if config.GVA_AORM != nil {
		print("!= nil")
		app.RegisterTablesAorm(config.GVA_AORM, true) // 初始化表
		// 程序结束前关闭数据库链接
		defer config.GVA_AORM.Close()
	}
}
