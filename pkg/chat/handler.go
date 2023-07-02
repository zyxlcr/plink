package chat

import (
	"chatcser/pkg/plink/iface"
	"fmt"
)

func (b BaseChat) Chat(res iface.ResponseWriter, req *iface.Request) {
	fmt.Println("chat")
	fmt.Println(string(req.GetBody()))
	req.GetConnection().SendMsgWithUrl("/chat/ack", []byte("333"))

}
