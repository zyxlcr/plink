package app

import (
	"chatcser/config"
	"chatcser/pkg/friend"
	"chatcser/pkg/utils/jsonx"

	"github.com/spf13/cast"
)

func (s *Service) MyFriends(ctx any) {
	config.GVA_LOG.Info("MyFriends")
	req, h := s.GetToken(ctx)
	f := friend.Friend{}
	f.Uid = cast.ToUint64(h.From)
	fs, err := f.MyFriends()
	if err != nil {
		config.GVA_LOG.Info("err: " + err.Error())
		req.GetConnection().SendMsgWithUrl("/friend/myfriend/ack/err", []byte(err.Error()))
		return
	}
	res := friend.MyFriendRes{}
	res.Data = fs
	res.Code = 0
	b, err := jsonx.Marshal(res)
	if err != nil {
		req.GetConnection().SendMsgWithUrl("/friend/myfriend/ack/err", []byte(err.Error()))
		return
	}
	req.GetConnection().SendMsgWithUrl("/friend/myfriend/ack", b)
}
