package app

import (
	"chatcser/config"
	"chatcser/pkg/friend"
	"chatcser/pkg/plink/iface"
	"chatcser/pkg/user"
	"chatcser/pkg/utils/jsonx"

	"github.com/spf13/cast"
)

func (s *Service) SearchUser(ctx any) {
	config.GVA_LOG.Info("SearchUser")
	req, ok := ctx.(*iface.Request)
	if !ok {
		// 失败处理：myInterface 不是 myType 类型
		req.GetConnection().SendMsgWithUrl("/user/search/ack/err", []byte{})
		return
	}
	var f friend.SearchReq
	err := s.ShouldBind(req, &f)
	if err != nil {
		config.GVA_LOG.Info(err.Error())
		req.GetConnection().SendMsgWithUrl("/user/search/ack/err", []byte(err.Error()))
		return
	}
	u := user.BaseUser{
		Name: f.Username,
		//ID: cast.ToInt64(f.Uid),
		Email: f.Email,
		Tel:   f.Tel,
	}
	u, err = u.Search()
	if err != nil {
		config.GVA_LOG.Info(err.Error())
		req.GetConnection().SendMsgWithUrl("/user/search/ack/err", []byte(err.Error()))
		return
	}

	var res = friend.SearchRes{
		Username:  u.Name,
		AvatarUrl: u.AvatarUrl,
		Uid:       cast.ToString(u.ID),
		Tel:       u.Tel,
		Email:     u.Email,
	}

	b, err := jsonx.Marshal(res)
	if err != nil {
		config.GVA_LOG.Info(err.Error())
		req.GetConnection().SendMsgWithUrl("/user/search/ack/err", []byte(err.Error()))
		return
	}

	req.GetConnection().SendMsgWithUrl("/user/search/ack", b)

}
