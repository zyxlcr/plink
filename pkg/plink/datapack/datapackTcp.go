package datapack

import (
	"bytes"
	"chatcser/pkg/plink/iface"
	"encoding/binary"
)

// 封包拆包类实例，暂时不需要成员
type DataPack struct{}

// 封包拆包实例初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

// 获取包头长度方法
func (dp *DataPack) GetHeadLen() uint32 {
	//Id uint32(4字节) +  HeaderLen uint32(4字节) +  BodyLen uint32(4字节)
	return 12
}

func (dp *DataPack) Pack(msg iface.IMessage) ([]byte, error) {
	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//写msgID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//写headerLen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetHeaderLen()); err != nil {
		return nil, err
	}

	//写bodyLen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetBodyLen()); err != nil {
		return nil, err
	}
	//写header
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetHeader()); err != nil {
		return nil, err
	}

	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetBody()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

// 拆包方法(解压数据)
func (dp *DataPack) Unpack(binaryData []byte) (iface.IMessage, error) {
	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	//只解压head的信息，得到dataLen和msgID
	msg := &iface.Message{}

	//读msgID
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	//读headerLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.HeaderLen); err != nil {
		return nil, err
	}

	//读bodyLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.BodyLen); err != nil {
		return nil, err
	}

	//读header
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Header); err != nil {
		return nil, err
	}
	// //读body
	// if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Body); err != nil {
	// 	return nil, err
	// }

	//判断dataLen的长度是否超出我们允许的最大包长度
	// if (utils.GlobalObject.MaxPacketSize > 0 && msg.DataLen > utils.GlobalObject.MaxPacketSize) {
	// 	return nil, errors.New("Too large msg data recieved")
	// }

	//这里只需要把head的数据拆包出来就可以了，然后再通过head的长度，再从conn读取一次数据
	return msg, nil
}
