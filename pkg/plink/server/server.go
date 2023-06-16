package server

import (
	"chatcser/pkg/plink/config"
	"chatcser/pkg/plink/connection"
	"chatcser/pkg/plink/grpc"
	"chatcser/pkg/plink/iface"
	"chatcser/pkg/plink/router"
	"chatcser/pkg/plink/tcpserver"
	ws "chatcser/pkg/plink/websocket"
	"fmt"
	"net/http"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const StatusNotFound = 404

type Server struct {
	Config *config.Config

	TcpServer  *tcpserver.TcpServer
	WsServer   *ws.WsServer
	GrpcServer *grpc.GrpcService

	Router iface.IRouter

	Wshandler http.Handler

	//当前Server的消息管理模块，用来绑定MsgId和对应的处理方法
	MsgHandler iface.IMsgHandle
	//当前Server的链接管理器
	ConnMgr iface.IConnManager
	//该Server的连接创建时Hook函数
	OnConnStart func(conn iface.IConnection)
	//该Server的连接断开时的Hook函数
	OnConnStop func(conn iface.IConnection)

	SessionsMtx sync.RWMutex

	mqttClient mqtt.Client
}

// NewServer creates a new server with the given options.
func NewServer() *Server {
	r := router.NewRouter()
	c := config.NewConfig()
	s := Server{
		Config: c,

		// grpcServer: grpc.NewServer(),
		Router: r,

		ConnMgr: connection.NewConnManager(),
	}
	s.MsgHandler = NewMsgHandle(&s)
	s.TcpServer = tcpserver.NewTcpServer(&s)
	s.WsServer = ws.Start(&s)
	s.GrpcServer = grpc.NewGrpcServer(&s)

	return &s
}

// 开启网络服务
func (s *Server) Start() {
	fmt.Printf("[START] Server name: %s,listenner at IP: %s, is starting\n", s.Config.Name, s.Config.IP)
	// fmt.Printf("[Zinx] Version: %s, MaxConn: %d, MaxPacketSize: %d\n",
	// 	GlobalObject.Version,
	// 	GlobalObject.MaxConn,
	// 	GlobalObject.MaxPacketSize)

	//0 启动worker工作池机制
	s.MsgHandler.StartWorkerPool()

	//开启一个go去做服务端 tcp 业务
	go func() {
		s.TcpServer.Start()

	}()

	//websocket
	go func() {
		err := ws.Listen(s, s.WsServer.Gin)
		if err != nil {
			fmt.Println(err)
		}
	}()

	//grpc
	go func() {
		err := grpc.Start(s.GrpcServer)
		if err != nil {
			fmt.Println(err)
		}
	}()

	// 心跳事件
	go s.dispatchLoop()
}

func (s *Server) dispatchLoop() {
	s.TcpServer.DispatchLoop()
}
func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Config.Name)

	//将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.ConnMgr.ClearConn()
}

func (s *Server) Serve() {
	done := make(chan struct{})

	// 启动其他 goroutine 并执行一些任务

	// 执行清理工作、打印日志等操作
	s.Start()

	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加

	//阻塞,否则主Go退出， listenner的go将会退出
	<-done // 等待收到信号量
	//select {}
}

func (s *Server) AddRouter(path string, handler iface.HandlerFunc) {
	fmt.Println("TcpServer AddRouter: ", path)
	//s.Router.Insert(path, handler)
	s.Router.Post(path, handler)
}
func (s *Server) AddRouter2(path uint32, handler iface.IRouter) {
	fmt.Println("AddRouter2 TcpServer")
	s.MsgHandler.AddRouter2(path, handler)
}

// 得到链接管理
func (s *Server) GetConnMgr() iface.IConnManager {
	return s.ConnMgr
}

func (s *Server) GetConfig() *config.Config {
	return s.Config
}

func (s *Server) GetRouter() iface.IRouter {
	return s.Router
}

func (s *Server) GetMsgHandler() iface.IMsgHandle {
	return s.MsgHandler
}

func (s *Server) SetWsHandle(h http.Handler) {
	s.Wshandler = h
}

// 设置该Server的连接创建时Hook函数
func (s *Server) SetOnConnStart(hookFunc func(iface.IConnection)) {
	s.OnConnStart = hookFunc
}

// 设置该Server的连接断开时的Hook函数
func (s *Server) SetOnConnStop(hookFunc func(iface.IConnection)) {
	s.OnConnStop = hookFunc
}

// 调用连接OnConnStart Hook函数
func (s *Server) CallOnConnStart(conn iface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

// 调用连接OnConnStop Hook函数
func (s *Server) CallOnConnStop(conn iface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}
