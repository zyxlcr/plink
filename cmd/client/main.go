package main

import (
	"chatcser/pkg/plink/datapack"
	"chatcser/pkg/plink/iface"
	"fmt"
	"net"
	"time"
)

/*
模拟客户端
*/
func ClientTest() {

	fmt.Println("Client Test ... start")
	//3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8989")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		msg := iface.NewMsgPackagewithUrl("/ping", []byte("okkkk ooo"))
		msg.SetMsgId(4196270080)
		data, err := datapack.NewDataPack().Pack(msg)
		if err != nil {
			return
		}
		var re int
		re, err = conn.Write(data)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}
		fmt.Println("write ok :", re)

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error ")
			return
		}

		fmt.Printf(" server call back : %s, cnt = %d\n", buf, cnt)

		time.Sleep(1 * time.Second)
	}
}

func main() {
	ClientTest()

	for {

	}
}
