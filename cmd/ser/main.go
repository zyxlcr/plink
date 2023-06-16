package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	idMutex     sync.Mutex
	lastID      uint32
	maxSequence uint16 = 0xFFFF
)

func GenerateMessageID() uint32 {
	idMutex.Lock()
	defer idMutex.Unlock()

	now := uint32(time.Now().UnixNano() / int64(time.Millisecond))

	if now == lastID {
		sequence := uint16(rand.Intn(int(maxSequence)))
		return (now << 16) | uint32(sequence)
	}

	lastID = now
	return (now << 16)
}

func main() {

	da := []byte("34534513")
	fmt.Println("headData = ", da)

	binaryData := make([]byte, 12)
	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer(binaryData)

	//写msgID
	if err := binary.Write(dataBuff, binary.LittleEndian, int(4196270080)); err != nil {
	}

	//写headerLen
	if err := binary.Write(dataBuff, binary.LittleEndian, uint32(55)); err != nil {
	}

	//写bodyLen
	binary.Write(dataBuff, binary.LittleEndian, uint32(3))

	fmt.Println("Data = ", dataBuff.Bytes())

	//binaryData := make([]byte, 12)
	//创建一个从输入二进制数据的ioReader
	dataBuff2 := bytes.NewReader(dataBuff.Bytes())

	//dataBuff2 = dataBuff
	var id uint32 = 1
	var headerLen uint32 = 1
	var bodyLen uint32 = 1

	//读msgID
	if err := binary.Read(dataBuff2, binary.LittleEndian, &id); err != nil {
		//return nil, err
	}
	//println(id)
	fmt.Println("id = ", id)

	if err := binary.Read(dataBuff2, binary.LittleEndian, &headerLen); err != nil {
		//return nil, err
	}

	fmt.Println("headerLen = ", headerLen)

	//读bodyLen
	if err := binary.Read(dataBuff2, binary.LittleEndian, &bodyLen); err != nil {
		//return nil, err
	}
	println(bodyLen)

}
