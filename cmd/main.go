package main

import (
	"chatcser/pkg/plink/iface"
	"chatcser/pkg/plink/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

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