package tcpserver

import (
	"chatcser/pkg/plink/config"
	"chatcser/pkg/plink/connection"
	"chatcser/pkg/plink/iface"
	"chatcser/pkg/plink/route"
	"net/http"

	"fmt"
	"net"
	"sync"
	"time"
)

type TcpServer struct {
	Config  *config.Config
	PServer iface.IServer

	Router iface.IRouter
	Route  *route.Router

	//当前Server的消息管理模块，用来绑定MsgId和对应的处理方法
	MsgHandler iface.IMsgHandle
	//当前Server的链接管理器
	ConnMgr iface.IConnManager
	//该Server的连接创建时Hook函数
	OnConnStart func(conn iface.IConnection)
	//该Server的连接断开时的Hook函数
	OnConnStop func(conn iface.IConnection)

	SessionsMtx sync.RWMutex
}

// NewServer creates a new server with the given options.
func NewTcpServer(s iface.IServer) *TcpServer {
	//r := router.NewRouter()
	return &TcpServer{
		Router:     s.GetRouter(),
		Route:      s.GetRoute(),
		MsgHandler: s.GetMsgHandler(),
		ConnMgr:    s.GetConnMgr(),
		Config:     s.GetConfig(),
	}
}

// 开启网络服务
func (s *TcpServer) Start() {
	fmt.Printf("[START] Server name: %s,listenner at IP: %s, Port %d is starting\n", s.Config.Name, s.Config.IP, s.Config.TcpConfig.TcpPort)

	//1 获取一个TCP的Addr
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", s.Config.IP, s.Config.TcpConfig.TcpPort))
	if err != nil {
		fmt.Println("resolve tcp addr err: ", err)
		return
	}

	//2 监听服务器地址
	listenner, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("listen", "err", err)
		return
	}

	//已经监听成功
	fmt.Println("start Zinx server  ", s.Config.Name, " succ, now listenning...")

	//3 启动server网络连接业务
	for {
		//3.1 阻塞等待客户端建立连接请求
		conn, err := listenner.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err ", err)
			continue
		}

		//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
		if s.ConnMgr.LenTcp() >= int(s.Config.TcpConfig.MaxConn) { //GlobalObject.MaxConn
			fmt.Println("超了")
			conn.Close()
			continue
		}

		//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
		dealConn := connection.NewTcpConntion(s, conn)
		//cid++

		//3.4 启动当前链接的处理业务
		go dealConn.StartTcp()

	}

}

func (s *TcpServer) DispatchLoop() {
	// 心跳事件
	ticker := time.NewTicker(s.Config.Heartbeat)
	defer ticker.Stop()

	cm := s.ConnMgr.(*connection.ConnManager)

	for {
		select {
		case c := <-cm.AddCh:
			fmt.Println("add ok", c)
		case c := <-cm.DelCh:
			fmt.Println("del ok", c)
		case <-ticker.C:
			data := []byte("ping ok!!!")
			//fmt.Println("send ping ")
			for _, conn := range cm.GetAll() {
				if err := conn.SendMsgWithUrl("/ping", data); err != nil {
					cm.Remove(conn.(*connection.Connection))
				}
				//conn.Write(data)
			}
		}
	}
}
func (s *TcpServer) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Config.Name)

	//将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.ConnMgr.ClearConn()
}

func (s *TcpServer) Serve() {
	done := make(chan struct{})

	// 启动其他 goroutine 并执行一些任务

	// 执行清理工作、打印日志等操作
	s.Start()

	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加

	//阻塞,否则主Go退出， listenner的go将会退出
	<-done // 等待收到信号量
	//select {}
}

func (s *TcpServer) AddRouter(path string, handler iface.HandlerFunc) {
	fmt.Println("TcpServer AddRouter ")
	//s.Router.Insert(path, handler)
	s.Router.Post(path, handler)
}
func (s *TcpServer) AddRouter2(path uint32, handler iface.IRouter) {
	fmt.Println("AddRouter2 TcpServer")
	s.MsgHandler.AddRouter2(path, handler)
}

// 得到链接管理
func (s *TcpServer) GetConnMgr() iface.IConnManager {
	return s.ConnMgr
}

func (s *TcpServer) GetConfig() *config.Config {
	return s.Config
}

func (s *TcpServer) GetRouter() iface.IRouter {
	return s.Router
}
func (s *TcpServer) GetRoute() *route.Router {
	return s.Route
}

func (s *TcpServer) GetMsgHandler() iface.IMsgHandle {
	return s.MsgHandler
}

func (s *TcpServer) SetWsHandle(h http.Handler) {
	//s.Wshandler = h
}

// 设置该Server的连接创建时Hook函数
func (s *TcpServer) SetOnConnStart(hookFunc func(iface.IConnection)) {
	s.OnConnStart = hookFunc
}

// 设置该Server的连接断开时的Hook函数
func (s *TcpServer) SetOnConnStop(hookFunc func(iface.IConnection)) {
	s.OnConnStop = hookFunc
}

// 调用连接OnConnStart Hook函数
func (s *TcpServer) CallOnConnStart(conn iface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

// 调用连接OnConnStop Hook函数
func (s *TcpServer) CallOnConnStop(conn iface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}
