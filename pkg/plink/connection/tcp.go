package connection

import (
	"chatcser/pkg/plink/datapack"
	"chatcser/pkg/plink/iface"
	"errors"
	"fmt"
	"io"
)

/*
写消息Goroutine， 用户将数据发送给客户端
*/
func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running]")
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
func (c *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Reader exit!]")
	defer c.Stop()

	for {
		// 创建拆包解包的对象
		dp := datapack.NewDataPack()

		//读取客户端的Msg head
		headData := make([]byte, dp.GetHeadLen())
		if c.GetTCPConnection() != nil {
			if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
				fmt.Println("read msg head error ", err)
				break
			}
		}

		//拆包，得到msgid 和 datalen 放在msg中
		msg, err := dp.Unpack(headData)
		fmt.Println("headData = ", headData)
		fmt.Println("msgID = ", msg.GetMsgId(), ", GetHeaderLen = ", msg.GetHeaderLen(), ", GetBodyLen = ", msg.GetBodyLen())
		if err != nil {
			fmt.Println("unpack error ", err)
			break
		}

		//根据 dataLen 读取 data，放在msg.Data中
		var data = make([]byte, msg.GetHeaderLen())
		if msg.GetHeaderLen() > 0 {
			data = make([]byte, msg.GetHeaderLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error ", err)
				break
			}
		}
		fmt.Println("data = ", string(data))
		msg.SetHeader(data)
		if msg.GetHeaderLen() > 0 {
			data = make([]byte, msg.GetBodyLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error ", err)
				break
			}
		}
		msg.SetBody(data)

		fmt.Printf("读到的消息头:%v 消息体:%v \n", string(msg.GetHeader()), string(msg.GetBody()))

		//得到当前客户端请求的Request数据
		req := iface.NewReqWithMsg(c, msg)

		// if utils.GlobalObject.WorkerPoolSize > 0 {
		// 	//已经启动工作池机制，将消息交给Worker处理
		//c.MsgHandler.SendMsgToTaskQueue(&req)
		// } else {
		// 	//从绑定好的消息和对应的处理方法中执行对应的Handle方法
		go c.MsgHandler.DoMsgHandler(&req)
		// }
	}
}

// 直接将Message数据发送数据给远程的TCP客户端
func (c *Connection) SendMsg(header []byte, data []byte) error {
	if c.IsTcpClosed == true {
		return errors.New("Connection closed when send msg")
	}
	//将data封包，并且发送
	imsg := iface.NewMsgPackage(header, data)
	msg, err := datapack.NewDataPack().Pack(imsg)
	if err != nil {
		fmt.Println("Pack error msg id = ", imsg.GetMsgId())
		return errors.New("Pack error msg ")
	}

	//写回客户端
	c.MsgChanTcp <- msg

	return nil
}

func (c *Connection) SendMsgWithUrl(url string, data []byte) error {
	if c.IsTcpClosed == true {
		return errors.New("Connection closed when send msg")
	}
	//将data封包，并且发送
	imsg := iface.NewMsgPackagewithUrl(url, data)

	//fmt.Println("msgid", imsg.GetMsgId(), "headerLen", imsg.GetHeaderLen(), "bodyLen", imsg.GetBodyLen())
	//fmt.Println("data", data, "string", string(data))
	msg, err := datapack.NewDataPack().Pack(imsg)
	if err != nil {
		fmt.Println("Pack error msg id = ", imsg.GetMsgId())
		return errors.New("Pack error msg ")
	}

	//写回客户端
	c.MsgChanTcp <- msg

	return nil
}

func (c *Connection) SendMsgWithUrlFromTo(url string, from string, to string, data []byte) error {
	if c.IsTcpClosed == true {
		return errors.New("Connection closed when send msg")
	}
	//将data封包，并且发送
	imsg := iface.NewMsgPackagewithUrlFromTo(url, from, to, data)

	//fmt.Println("msgid", imsg.GetMsgId(), "headerLen", imsg.GetHeaderLen(), "bodyLen", imsg.GetBodyLen())
	//fmt.Println("data", data, "string", string(data))
	msg, err := datapack.NewDataPack().Pack(imsg)
	if err != nil {
		fmt.Println("Pack error msg id = ", imsg.GetMsgId())
		return errors.New("Pack error msg ")
	}

	//写回客户端
	c.MsgChanTcp <- msg

	return nil
}

func (c *Connection) SendBuffMsg(header []byte, data []byte) error {
	if c.IsTcpClosed == true {
		return errors.New("Connection closed when send msg")
	}
	//将data封包，并且发送
	imsg := iface.NewMsgPackage(header, data)
	msg, err := datapack.NewDataPack().Pack(imsg)
	if err != nil {
		fmt.Println("Pack error msg id = ", imsg.GetMsgId())
		return errors.New("Pack error msg ")
	}

	//写回客户端
	c.MsgBuffChanTcp <- msg

	return nil
}
