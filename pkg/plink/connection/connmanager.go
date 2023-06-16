package connection

import (
	"chatcser/pkg/plink/iface"
	"errors"
	"fmt"
	"sync"
)

type HandlerFunc func(iface.ResponseWriter, *iface.Request)

/*
	连接管理抽象层
*/

/*
连接管理模块
*/
type ConnManager struct {
	connectionsTcp map[string]iface.IConnection //管理的连接信息
	connectionsWs  map[string]iface.IConnection //管理的连接信息
	connLock       sync.RWMutex                 //读写连接的读写锁
	AddCh          chan *iface.IConnection
	DelCh          chan *iface.IConnection
}

/*
创建一个链接管理
*/
func NewConnManager() *ConnManager {
	return &ConnManager{
		connectionsTcp: make(map[string]iface.IConnection),
	}
}

// 添加链接
func (connMgr *ConnManager) Add(conn iface.IConnection) {
	//保护共享资源Map 加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//将conn连接添加到ConnMananger中
	connMgr.connectionsTcp[conn.GetConnID()] = conn
	//connMgr.addCh <- &conn
	fmt.Println("connection add to ConnManager successfully", conn.GetConnID(), ": conn num = ", connMgr.LenTcp())
}

// 删除连接
func (connMgr *ConnManager) Remove(conn iface.IConnection) {
	//保护共享资源Map 加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()
	
	//删除连接信息
	delete(connMgr.connectionsTcp, conn.GetConnID())
	//connMgr.delCh <- &conn
	fmt.Println("connection Remove ConnID=", conn.GetConnID(), " successfully: conn num = ", connMgr.LenTcp())
}

// 利用ConnID获取链接
func (connMgr *ConnManager) Get(connID string) (iface.IConnection, error) {
	//保护共享资源Map 加读锁
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()

	if conn, ok := connMgr.connectionsTcp[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

func (connMgr *ConnManager) GetAll() map[string]iface.IConnection {
	//保护共享资源Map 加读锁
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()

	return connMgr.connectionsTcp
}

// 获取当前连接
func (connMgr *ConnManager) LenTcp() int {
	return len(connMgr.connectionsTcp)
}
func (connMgr *ConnManager) LenWs() int {
	return len(connMgr.connectionsWs)
}

// 清除并停止所有连接
func (connMgr *ConnManager) ClearConn() {
	//保护共享资源Map 加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//停止并删除全部的连接信息
	for connID, conn := range connMgr.connectionsTcp {
		//停止
		conn.Stop()
		//删除
		delete(connMgr.connectionsTcp, connID)
	}

	fmt.Println("Clear All Connections successfully: conn num = ", connMgr.LenTcp())
}
