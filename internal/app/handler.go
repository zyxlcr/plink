package app

import (
	"errors"
	"fmt"

	"chatcser/config"
	"chatcser/pkg/auth"
	"chatcser/pkg/notification"
	"chatcser/pkg/plink/iface"
	"chatcser/pkg/plink/route"
	"chatcser/pkg/user"
	"chatcser/pkg/utils"
	"chatcser/pkg/utils/jsonx"

	"github.com/spf13/cast"
)

func (s Service) Ping(res iface.ResponseWriter, req *iface.Request) {
	config.GVA_LOG.Info("ping")
	s.User.Ping()

}
func (s Service) Ping2(ctx any) {
	config.GVA_LOG.Info("ping")

}

func (s Service) Login(ctx any) {
	//h := iface.Header{}
	req, ok := ctx.(*iface.Request)
	if !ok {
		// 失败处理：myInterface 不是 myType 类型
		req.GetConnection().SendMsgWithUrl("/login/ack/err", []byte{})
		return
	}
	var u auth.LoginReq
	err := s.ShouldBind(req, &u)
	if err != nil {
		fmt.Println(err)
		req.GetConnection().SendMsgWithUrl("/login/ack/err", []byte(err.Error()))
		return
	}
	fmt.Println(u.Username)
	//var res *user.LoginRes
	res, err := auth.Login(u)
	if err != nil {
		config.GVA_LOG.Info(err.Error())
		req.GetConnection().SendMsgWithUrl("/login/ack/err", []byte(err.Error()))
		return
	}
	if err != nil {
		fmt.Println(err)
		req.GetConnection().SendMsgWithUrl("/login/ack/err", []byte(err.Error()))
		return
	}
	b, err := jsonx.Marshal(res)
	config.GVA_LOG.Info(string(b))
	req.GetConnection().SendMsgWithUrl("/login/ack", b)

}

func (s Service) Reg(ctx any) {
	config.GVA_LOG.Info("Reg")
	req, ok := ctx.(*iface.Request)
	if !ok {
		// 失败处理：myInterface 不是 myType 类型
		req.GetConnection().SendMsgWithUrl("/reg/ack/err", []byte{})
		return
	}
	var u user.BaseUser
	err := s.ShouldBind(req, &u)
	if err != nil {
		fmt.Println(err)
		req.GetConnection().SendMsgWithUrl("/reg/ack/err", []byte{})
		return
	}
	fmt.Println(u.Name)
	err = auth.Reg(u, req)
	if err != nil {
		fmt.Println(err)
		req.GetConnection().SendMsgWithUrl("/reg/ack/err", []byte{})
		return
	}
	req.GetConnection().SendMsgWithUrl("/reg/ack", []byte("333"))
}

func (s Service) ChatHandler2(ctx any) {
	config.GVA_LOG.Info("ChatHandler2")
	//s.Chat.Chat(res, req)

}

func (s Service) Send(ctx any) {
	config.GVA_LOG.Info("Send")
	req, h := s.GetToken(ctx)
	var n notification.SendReq
	err := s.ShouldBind(req, &n)
	if err != nil {
		fmt.Println(err)
		req.GetConnection().SendMsgWithUrl("/notification/send/ack/err", []byte(err.Error()))
		return
	}

	user, err := auth.GetAuth(h.Token)

	nc := notification.Notification{
		Type:         "makeFriend",
		From:         user.UID,
		To:           cast.ToInt64(h.To),
		FromUsername: user.Username,
		Content:      n.Content,
	}
	err = nc.Add()
	if err != nil {
		req.GetConnection().SendMsgWithUrl("/notification/send/ack/err", []byte(err.Error()))
		return
	}

	req.GetConnection().SendMsgWithUrl("/notification/send/ack", []byte("333"))

}

func (s Service) DoNotification(ctx any) {
	config.GVA_LOG.Info("DoNotification")
	req, _ := s.GetToken(ctx)
	var body notification.DoAddFriendReq
	err := s.ShouldBind(req, &body)
	if err != nil {
		fmt.Println(err)
		req.GetConnection().SendMsgWithUrl("/notification/doAddfriend/ack/err", []byte(err.Error()))
		return
	}
	nc := notification.Notification{}
	nc.ID = body.Mid
	err = nc.DoAddFriend(body.Do)
	if err != nil {
		req.GetConnection().SendMsgWithUrl("/notification/doAddfriend/ack/err", []byte(err.Error()))
		return

	}
	req.GetConnection().SendMsgWithUrl("/notification/doAddfriend/ack", []byte("ok"))
}
func (s Service) GetToken(ctx any) (r *iface.Request, h iface.Header) {
	//h := iface.Header{}
	req, ok := ctx.(*iface.Request)
	if !ok {
		// 失败处理：myInterface 不是 myType 类型
		req.GetConnection().SendMsgWithUrl("/auth/ack/err", []byte{})
		return r, h
	}
	iface.FromJsonTo(req.GetHeader(), &h)
	return req, h

}

func (s Service) Auth2(next route.HandlerFun) route.HandlerFun {
	//config.GVA_LOG.Info("Auth2 s1")
	return func(ctx any) {
		req, h := s.GetToken(ctx)
		uid, err := auth.Auth(h.Token)
		utils.CheckError(err)
		if err != nil {
			req.GetConnection().SendMsgWithUrl("/auth/ack/err", []byte(err.Error()))
			return
		}
		if cast.ToString(uid) != h.From {
			//return res, req
			utils.CheckError(errors.New("token eror"))
			req.GetConnection().SendMsgWithUrl("/auth/ack/err", []byte("token eror"))
			return
		}
		//return res, req
		next(ctx)
		//config.GVA_LOG.Info("Auth2 end")
	}

}
