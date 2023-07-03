package chat

import (
	"chatcser/config"
	"chatcser/pkg/plink/iface"
)

func (b BaseChat) Chat(res iface.ResponseWriter, req *iface.Request) {
	config.GVA_LOG.Info("chat")
	config.GVA_LOG.Info(string(req.GetBody()))
	req.GetConnection().SendMsgWithUrl("/chat/ack", []byte("333"))

}
