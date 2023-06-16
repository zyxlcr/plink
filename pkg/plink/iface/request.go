package iface

import (
	"fmt"
	"net/url"
)

/*
	IRequest 接口：
	实际上是把客户端请求的链接信息 和 请求的数据 包装到了 Request里
*/

func NewPsotReq(conn IConnection, uri string, msg IMessage) Request {
	return Request{
		conn:   conn,
		Method: "POST",
		URL:    &url.URL{Path: uri},
		msg:    msg,
	}
}

func NewReq(conn IConnection, header []byte, data []byte) Request {
	msg := NewMsgPackage(header, data)
	return Request{
		conn:   conn,
		Method: "POST",
		URL:    &url.URL{Path: msg.GetUrl()},
		msg:    msg,
	}
}

func NewReqWithMsg(conn IConnection, msg IMessage) Request {
	return Request{
		conn:   conn,
		Method: "POST",
		msg:    msg,
		URL:    &url.URL{Path: msg.GetUrl()},
	}
}

func NewPingReq(conn IConnection) Request {
	h, err := NewHeader("/ping").ToJson()
	if err != nil {
		//
		fmt.Println("error")
	}
	return Request{
		conn:   conn,
		Method: "POST",
		URL:    &url.URL{Path: "/ping"},
		msg:    NewMsgPackage(h, []byte("ping\n")),
	}
}

type HandlerFunc func(ResponseWriter, *Request)
type HandlerFun func(ctx any)

type IRequest interface {
	GetConnection() IConnection //获取请求连接信息
	GetHeaderLen() uint32       //获取请求消息的数据
	GetHeader() []byte          //获取请求消息的数据
	GetBody() []byte            //获取请求消息的数据
	GetBodyLen() uint32         //获取请求消息的数据
	GetMsgID() uint32           //获取请求的消息ID
	GetMsg() IMessage
}

type Request struct {
	conn   IConnection //已经和客户端建立好的 链接
	msg    IMessage    //客户端请求的数据
	Method string
	//ctx           context.Context
	URL *url.URL
	//ContentLength int64
	//RemoteAddr    string
	//RequestURI    string
	//Cancel        <-chan struct{}
	//Response      *Response
}

//获取请求连接信息
func (r *Request) GetConnection() IConnection { // ziface.IConnection {
	return r.conn
}

//获取请求消息的数据
func (r *Request) GetBody() []byte {
	return r.msg.GetBody()
}

func (r *Request) GetHeader() []byte {
	return r.msg.GetHeader()
}

func (r *Request) GetBodyLen() uint32 {
	return r.msg.GetBodyLen()
}

func (r *Request) GetHeaderLen() uint32 {
	return r.msg.GetHeaderLen()
}

//获取请求的消息的ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}

//获取 消息
func (r *Request) GetMsg() IMessage {
	return r.msg
}

// func (r *Request) WithContext(context.Context) {

// }
