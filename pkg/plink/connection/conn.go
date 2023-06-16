package connection

import (
	"chatcser/pkg/plink/iface"
	"fmt"
	"net"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/rs/xid"
)

type Connection struct {
	//当前Conn属于哪个Server
	Server iface.IServer
	//当前连接的socket TCP套接字
	ConnTcp *net.TCPConn

	ConnWs *websocket.Conn
	//当前连接的ID 也可以称作为SessionID，ID全局唯一
	ConnID string
	//当前连接的关闭状态
	IsTcpClosed bool

	IsWsClosed bool
	//消息管理MsgId和对应处理方法的消息管理模块
	MsgHandler iface.IMsgHandle

	//告知该链接已经退出/停止的channel
	ExitBuffChanTcp chan bool
	//无缓冲管道，用于读、写两个goroutine之间的消息通信
	MsgChanTcp chan []byte
	//有关冲管道，用于读、写两个goroutine之间的消息通信
	MsgBuffChanTcp chan []byte

	//告知该链接已经退出/停止的channel
	ExitBuffChanWs chan bool
	//无缓冲管道，用于读、写两个goroutine之间的消息通信
	MsgChanWs chan []byte
	//有关冲管道，用于读、写两个goroutine之间的消息通信
	MsgBuffChanWs chan []byte

	//链接属性
	property map[string]interface{}
	//保护链接属性修改的锁
	propertyLock sync.RWMutex
}

// 创建连接的方法
func NewTcpConntion(server iface.IServer, conn *net.TCPConn) *Connection {
	//初始化Conn属性
	c := &Connection{
		Server:          server,
		ConnTcp:         conn,
		ConnID:          xid.New().String(),
		IsTcpClosed:     false,
		IsWsClosed:      true,
		MsgHandler:      server.GetMsgHandler(),
		ExitBuffChanTcp: make(chan bool, 1),
		MsgChanTcp:      make(chan []byte),
		MsgBuffChanTcp:  make(chan []byte, server.GetConfig().MaxMsgChanLen),
		property:        make(map[string]interface{}),
	}

	//将新创建的Conn添加到链接管理中
	c.Server.GetConnMgr().Add(c)
	return c
}

func NewWsConntion(server iface.IServer, conn *websocket.Conn) *Connection {
	//初始化Conn属性
	c := &Connection{
		Server:         server,
		ConnWs:         conn,
		ConnID:         xid.New().String(),
		IsTcpClosed:    true,
		IsWsClosed:     false,
		MsgHandler:     server.GetMsgHandler(),
		ExitBuffChanWs: make(chan bool, 1),
		MsgChanWs:      make(chan []byte),
		MsgBuffChanWs:  make(chan []byte, server.GetConfig().MaxMsgChanLen),
		property:       make(map[string]interface{}),
	}

	//将新创建的Conn添加到链接管理中
	c.Server.GetConnMgr().Add(c)
	return c
}

// 获取当前连接ID
func (c *Connection) GetConnID() string {
	return c.ConnID
}

// 启动连接，让当前连接开始工作
func (c *Connection) StartTcp() {
	//1 开启用户从客户端读取数据流程的Goroutine
	go c.StartReader()
	//2 开启用于写回客户端数据流程的Goroutine
	go c.StartWriter()
	//按照用户传递进来的创建连接时需要处理的业务，执行钩子方法
	c.Server.CallOnConnStart(c)
}

// 启动连接，让当前连接开始工作
func (c *Connection) StartWs() {
	//1 开启用户从客户端读取数据流程的Goroutine
	go c.StartWsReader()
	//2 开启用于写回客户端数据流程的Goroutine
	go c.StartWsWriter()
	//按照用户传递进来的创建连接时需要处理的业务，执行钩子方法
	c.Server.CallOnConnStart(c)
}

// 停止连接，结束当前连接状态M
func (c *Connection) Stop() {
	fmt.Println("Conn Stop()...ConnID = ", c.ConnID)
	//如果当前链接已经关闭
	if c.IsTcpClosed == true || c.IsWsClosed == true {
		return
	} else if c.IsTcpClosed == false {
		c.IsTcpClosed = true

		//如果用户注册了该链接的关闭回调业务，那么在此刻应该显示调用
		c.Server.CallOnConnStop(c)

		// 关闭socket链接
		c.ConnTcp.Close()
		//关闭Writer
		c.ExitBuffChanTcp <- true

		//将链接从连接管理器中删除
		c.Server.GetConnMgr().Remove(c)

		//关闭该链接全部管道
		close(c.ExitBuffChanTcp)
		close(c.MsgBuffChanTcp)
	} else if c.IsWsClosed == false {
		c.IsWsClosed = true

		//如果用户注册了该链接的关闭回调业务，那么在此刻应该显示调用
		c.Server.CallOnConnStop(c)

		// 关闭socket链接
		c.ConnWs.Close()
		//关闭Writer
		c.ExitBuffChanWs <- true

		//将链接从连接管理器中删除
		c.Server.GetConnMgr().Remove(c)

		//关闭该链接全部管道
		close(c.ExitBuffChanTcp)
		close(c.MsgBuffChanTcp)
	}

}

// 从当前连接获取原始的socket TCPConn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.ConnTcp
}

func (c *Connection) GetWsConnection() *websocket.Conn {
	return c.ConnWs
}

// 获取远程客户端地址信息
func (c *Connection) RemoteAddr() net.Addr {
	if c.ConnTcp != nil {
		return c.ConnTcp.RemoteAddr()
	} else {
		return c.ConnWs.RemoteAddr()
	}

}

func (c *Connection) Write(b []byte) {
	if c.ConnTcp != nil {
		c.ConnTcp.Write(b)
	}

}
