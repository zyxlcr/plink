package iface

import (
	"chatcser/pkg/plink/config"
	"chatcser/pkg/plink/route"
	"net"
	"net/http"
)

type IServer interface {
	//启动服务器方法
	Start()
	//停止服务器方法
	Stop()
	//开启业务服务方法
	Serve()
	//路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
	AddRouter2(msgId uint32, router IRouter)
	//AddRouter(path string, h HandlerFunc)
	AddRouter(path string, h HandlerFunc)
	//得到链接管理
	GetConnMgr() IConnManager

	GetRouter() IRouter
	GetRoute() *route.Router
	GetConfig() *config.Config
	GetMsgHandler() IMsgHandle
	SetWsHandle(h http.Handler)

	//设置该Server的连接创建时Hook函数
	SetOnConnStart(func(IConnection))
	//设置该Server的连接断开时的Hook函数
	SetOnConnStop(func(IConnection))
	//调用连接OnConnStart Hook函数
	CallOnConnStart(conn IConnection)
	//调用连接OnConnStop Hook函数
	CallOnConnStop(conn IConnection)
}

type IConnManager interface {
	Add(conn IConnection)                   //添加链接
	Remove(conn IConnection)                //删除连接
	Get(connID string) (IConnection, error) //利用ConnID获取链接
	LenTcp() int                            //获取当前连接
	LenWs() int                             //获取当前连接
	ClearConn()                             //删除并停止所有链接
	GetAll() map[string]IConnection
}

type IConnection interface {
	//启动连接，让当前连接开始工作
	StartTcp()
	StartWs()
	//停止连接，结束当前连接状态M
	Stop()
	//获取当前连接ID
	GetConnID() string

	SendMsg([]byte, []byte) error

	SendMsgWithUrl(string, []byte) error
	//SendMsgWithUrl(url string, data []byte) error

	// ReadMessage() ([]byte, error)

	GetTCPConnection() *net.TCPConn

	Write(b []byte)
}

type IMsgHandle interface {
	DoMsgHandler(request IRequest)           //马上以非阻塞方式处理消息
	AddRouter2(msgId uint32, router IRouter) //为消息添加具体的处理逻辑
	//AddRouter(path string, h HandlerFunc)
	AddRouter(path string, h HandlerFunc)
	StartWorkerPool()                    //启动worker工作池
	SendMsgToTaskQueue(request IRequest) //将消息交给TaskQueue,由worker进行处理
}

/*
封包数据和拆包数据
直接面向TCP连接中的数据流,为传输数据添加头部信息，用于处理TCP粘包问题。
*/
type IDataPack interface {
	GetHeadLen() uint32              //获取包头长度方法
	Pack(IMessage) ([]byte, error)   //封包方法
	Unpack([]byte) (IMessage, error) //拆包方法
}

type IRouter interface {
	//GetHandler(req *Request) HandlerFunc
	GetHandlerWithUrl(string) HandlerFunc
	Use(...Middleware)
	Post(string, HandlerFunc)
	GroupWithMore(prefix string, middleware ...Middleware) IRouter
	Group(prefix string) IRouter
}

type IRouter2 interface {
	PreHandle(request IRequest)  //在处理conn业务之前的钩子方法
	Handle(request IRequest)     //处理conn业务的方法
	PostHandle(request IRequest) //处理conn业务之后的钩子方法
}

type Middleware func(HandlerFunc) HandlerFunc
