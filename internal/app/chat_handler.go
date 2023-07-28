package app

import (
	"fmt"
	"time"

	"chatcser/config"
	"chatcser/pkg/chat"
	"chatcser/pkg/user"
	"chatcser/pkg/utils/jsonx"

	"github.com/spf13/cast"
)

// 服务端
func (s *Service) ChatHandler2(ctx any) {
	config.GVA_LOG.Info("ChatHandler2")

	req, h := s.GetToken(ctx)
	u := user.BaseUser{}
	u.ID = cast.ToInt64(h.From)
	info, err := u.Search()
	if err != nil {
		config.GVA_LOG.Info("err: " + err.Error())
		req.GetConnection().SendMsgWithUrl("/chat/ack/err", []byte(err.Error()))
		return
	}
	currentTime := time.Now()
	formattedTime := currentTime.Format("15:04:05") //2006-01-02 15:04:05
	res := chat.ToChatReq{
		Uid:         h.From,
		MsgType:     "one",
		IsMe:        false,
		FriendId:    h.To,
		FriendName:  info.Name,
		UpdateAt:    formattedTime,
		ContentType: "text",
		Content:     string(req.GetBody()),
		AvatarURL:   info.AvatarUrl,
	}
	b, err := jsonx.Marshal(res)
	if err != nil {
		req.GetConnection().SendMsgWithUrl("/chat/ack/err", []byte(err.Error()))
		return
	}
	req.GetConnection().SendMsgWithUrl("/chat/ack", []byte("ok"))
	// TODO: 给接收人发消息
	config.GVA_LOG.Info("给接收人发消息h.To: " + h.To)
	fmt.Printf("%v \n", s.TcpMap)
	if s.ReadTcpMap(h.To) != nil {
		err = s.ReadTcpMap(h.To).SendMsgWithUrl("/chat", b)
		if err != nil {
			config.GVA_LOG.Info(err.Error())
			req.GetConnection().SendMsgWithUrl("/chat/ack/err", []byte(err.Error()))
			return
		}
	}
}
