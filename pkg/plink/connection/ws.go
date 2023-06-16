package connection

import (
	"chatcser/pkg/plink/datapack"
	"chatcser/pkg/plink/iface"
	"fmt"
)

/*
写消息Goroutine， 用户将数据发送给客户端
*/
func (c *Connection) StartWsWriter() {
	fmt.Println("[Writer ws Goroutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Writer exit!]")

	for {
		select {
		case data := <-c.MsgChanTcp:
			//有数据要写给客户端
			if _, err := c.ConnTcp.Write(data); err != nil {
				c.Server.GetConnMgr().Remove(c)
				fmt.Println("Send Data error:, ", err, " Conn Writer exit")
				return
			}
		case data, ok := <-c.MsgBuffChanTcp:
			if ok {
				//有数据要写给客户端
				if _, err := c.ConnTcp.Write(data); err != nil {
					c.Server.GetConnMgr().Remove(c)
					fmt.Println("Send Buff Data error:, ", err, " Conn Writer exit")
					return
				}
			} else {
				break
				//fmt.Println("msgBuffChan is Closed")
			}
		case <-c.ExitBuffChanTcp:
			return
		}
	}
}

/*
读消息Goroutine，用于从客户端中读取数据
*/
func (c *Connection) StartWsReader() {
	fmt.Println("[Reader ws Goroutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn ws Reader exit!]")
	defer c.Stop()

	var bodyData = make([]byte, 1024)
	var n = 1
	var tempMsp = iface.Message{}
	for {

		msgType, msgbyte, err := c.ConnWs.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		fmt.Printf("recv: %s, type: %d\n", string(msgbyte), msgType)

		// 创建拆包解包的对象
		dp := datapack.NewDataPack()

		//读取客户端的Msg head
		headData := make([]byte, dp.GetHeadLen())
		headData = msgbyte[:dp.GetHeadLen()]

		//拆包，得到msgid 和 datalen 放在msg中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error ", err)
			break
		}
		tempMsp.SetHeader(headData)

		//根据 HeaderLen 读取 data，放在msg.header中
		var data = make([]byte, 1024)
		if tempMsp.GetHeaderLen() > 0 {
			data = make([]byte, tempMsp.GetHeaderLen())
			data = msgbyte[dp.GetHeadLen()-1 : tempMsp.GetHeaderLen()]
		}
		msg.SetHeader(data)

		//
		if tempMsp.GetHeaderLen()+tempMsp.GetBodyLen()+dp.GetHeadLen() > uint32(len(msgbyte)) && tempMsp.GetHeaderLen()+tempMsp.GetBodyLen()+dp.GetHeadLen() > uint32(1024*n) {
			//tempMsp = msg
			if n == 1 {
				bodyData = bodyData[dp.GetHeadLen()+tempMsp.GetHeaderLen()-1:]
			} else {
				bodyData = append(bodyData, msgbyte...)
			}
			n++
		} else {
			if n == 1 {
				if msg.GetBodyLen() > 0 {
					data = make([]byte, msg.GetBodyLen())
					data = msgbyte[dp.GetHeadLen()+msg.GetHeaderLen()-1:]
				}
				msg.SetBody(data)

				fmt.Printf("读到的消息头:%v 消息体:%v \n", string(msg.GetHeader()), string(msg.GetBody()))

				//得到当前客户端请求的Request数据
				req := iface.NewReqWithMsg(c, msg)

				go c.MsgHandler.DoMsgHandler(&req)
			} else {
				bodyData = append(bodyData, msgbyte...)
			}

			msg.SetHeader([]byte(""))
			bodyData = []byte("")
			n = 1
		}

	}
}
