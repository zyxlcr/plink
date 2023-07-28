package main

import (
	"chatcser/config"
	"chatcser/internal/app"
	"chatcser/pkg/plink/iface"
	"chatcser/pkg/plink/server"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 初始化配置
	config.GVA_VP = app.Viper()
	// 初始化日志
	config.GVA_LOG = app.Zap()
	zap.ReplaceGlobals(config.GVA_LOG)
	// 初始化redis
	config.GVA_REDIS = app.Reids()

	// 初始化数据库
	config.GVA_DB = app.Gorm() // gorm连接数据库
	if config.GVA_DB != nil {
		// 程序结束前关闭数据库链接
		db, _ := config.GVA_DB.DB()
		defer db.Close()
	}
	// 初始化数据库
	config.GVA_AORM = app.Aorm() // gorm连接数据库
	if config.GVA_AORM != nil {
		// 程序结束前关闭数据库链接
		config.GVA_LOG.Info("db con != nil2")
		defer config.GVA_AORM.Close()
	}
	app.RegisterTablesAorm(config.GVA_AORM, false)
	s := app.NewService()
	s.Run()

}

func main23() {

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})

	//r := router2.NewPrefixTree()
	s := server.NewServer()
	//s.Router = r
	//s.AddRouter2(10, &PingRouter{})
	//r.Insert("/ping", helloWorld)

	//s.SetWsHandle(r)
	s.AddRouter("/ping", helloWorld2)
	s.WsServer.Gin.GET("/he", he)

	fmt.Println("Welcome")
	s.Serve()
	fmt.Println("Welcomewww")

}

func helloWorld(r any) {
	fmt.Println("1236666 helloWorld dddd")
	//w.Write([]byte("Hello World"))
}

func helloWorld2(res iface.ResponseWriter, req *iface.Request) {
	println("1236666 helloWorld\n")
	//println("1236666 helloWorld\n", res)
	req.GetConnection().SendMsgWithUrl("/ping/ack", []byte("333")) //.Write([]byte("Hello World"))
	//w.Write([]byte("Hello World"))
}

// ping test 自定义路由
type PingRouter struct {
}

// Test PreHandle
func (this *PingRouter) PreHandle(request iface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

// Test Handle
func (this *PingRouter) Handle(request iface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

// Test PostHandle
func (this *PingRouter) PostHandle(request iface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping .....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func he(ctx *gin.Context) {
	log.Println("hello")
}
