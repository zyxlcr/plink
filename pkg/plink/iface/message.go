package iface

import (
	"net/url"
)

type IMessage interface {
	GetBodyLen() uint32   //获取消息数据段长度
	GetMsgId() uint32     //获取消息ID
	GetBody() []byte      //获取消息内容
	GetHeaderLen() uint32 //获取头长度
	GetHeader() []byte    //获取头内容

	SetMsgId(uint32)     //设计消息ID
	SetBody([]byte)      //设置消息内容
	SetBodyLen(uint32)   //设置消息数据段长度
	SetHeader([]byte)    //设置头内容
	SetHeaderLen(uint32) //设置头长度

	Pack() ([]byte, error)

	GetUrl() string
	ToReq(conn IConnection) *Request
}

/*
	将请求的一个消息封装到message中，定义抽象层接口
*/

type Message struct {
	Id        uint32 //消息的ID
	HeaderLen uint32 //头的长度
	Header    []byte //头的内容
	BodyLen   uint32 //消息的长度
	Body      []byte //消息的内容
}

// 创建一个Message消息包
func NewMsgPackage(header []byte, data []byte) *Message {
	return &Message{
		Id:        GenerateMessageID(),
		HeaderLen: uint32(len(header)),
		Header:    header,
		BodyLen:   uint32(len(data)),
		Body:      data,
	}
}

func NewMsgPackagewithUrl(url string, data []byte) *Message {
	h, err := NewHeader(url).ToJson()
	if err != nil {
		return nil
	}
	return &Message{
		Id:        GenerateMessageID(),
		HeaderLen: uint32(len(h)),
		Header:    []byte(h),
		BodyLen:   uint32(len(data)),
		Body:      data,
	}
}

func NewMsgPackagewithUrlFromTo(url string, from string, to string, data []byte) *Message {
	h, err := NewHeaderFromTo(url, from, to).ToJson()
	if err != nil {
		return nil
	}
	return &Message{
		Id:        GenerateMessageID(),
		HeaderLen: uint32(len(h)),
		Header:    []byte(h),
		BodyLen:   uint32(len(data)),
		Body:      data,
	}
}

// 获取消息数据段长度
func (msg *Message) GetBodyLen() uint32 {
	return msg.BodyLen
}

// 获取消息ID
func (msg *Message) GetMsgId() uint32 {
	return msg.Id
}

// 获取消息内容
func (msg *Message) GetBody() []byte {
	return msg.Body
}

// 获取消息内容
func (msg *Message) GetHeaderLen() uint32 {
	return msg.HeaderLen
}

// 获取消息内容
func (msg *Message) GetHeader() []byte {
	return msg.Header
}

func (msg *Message) GetUrl() string {
	var h Header
	FromJsonTo(msg.GetHeader(), &h)
	return h.Url
}

// 设置消息数据段长度
func (msg *Message) SetBodyLen(len uint32) {
	msg.BodyLen = len
}

// 设计消息ID
func (msg *Message) SetMsgId(msgId uint32) {
	msg.Id = msgId
}

// 设计消息内容
func (msg *Message) SetBody(data []byte) {
	msg.Body = data
}

// 设置头长度
func (msg *Message) SetHeaderLen(msgId uint32) {
	msg.HeaderLen = msgId
}

// 设置头内容
func (msg *Message) SetHeader(data []byte) {
	msg.Header = data
}

func (msg *Message) Pack() ([]byte, error) {
	//return IDataPack.Pack(msg)
	return []byte(""), nil
}

func (msg *Message) ToReq(conn IConnection) *Request {
	return &Request{
		conn:   conn,
		msg:    msg,
		Method: "POST",
		URL:    &url.URL{Path: msg.GetUrl()},
	}
}
