package app

import (
	"errors"
	"fmt"

	"chatcser/pkg/plink/iface"
	"chatcser/pkg/plink/route"
	"chatcser/pkg/user"
	"chatcser/pkg/utils"
	"chatcser/pkg/utils/jsonx"

	"github.com/spf13/cast"
)

func (s Service) Ping(res iface.ResponseWriter, req *iface.Request) {
	fmt.Println("ping")
	s.User.Ping(res, req)

}
func (s Service) Ping2(ctx any) {
	fmt.Println("ping")
	//s.User.Ping(res, req)

}

func (s Service) Login(ctx any) {
	//h := iface.Header{}
	req, ok := ctx.(*iface.Request)
	if !ok {
		// 失败处理：myInterface 不是 myType 类型
		return
	}
	var u user.LoginReq
	err := s.ShouldBind(req, &u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.Username)
	var res *user.LoginRes
	res, err = s.User.Login(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := jsonx.Marshal(res)
	req.GetConnection().SendMsgWithUrl("/login/ack", b)

}

func (s Service) Reg(res iface.ResponseWriter, req *iface.Request) {
	fmt.Println("Reg")
	var u user.BaseUser
	err := s.ShouldBind(req, &u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.Name)
	err = s.User.Reg(u, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.GetConnection().SendMsgWithUrl("/reg/ack", []byte("333"))
}

func (s Service) ChatHandler(res iface.ResponseWriter, req *iface.Request) {
	fmt.Println("ping")
	//s.Chat.Chat(res, req)
	req.GetConnection().SendMsgWithUrl("/chat/ack", []byte("333"))

}
func (s Service) ChatHandler2(ctx any) {
	fmt.Println("ChatHandler2")
	//s.Chat.Chat(res, req)

}

func (s Service) Auth(iface.HandlerFunc) iface.HandlerFunc {
	fmt.Println("Auth")
	return func(res iface.ResponseWriter, req *iface.Request) {
		h := iface.Header{}
		iface.FromJsonTo(req.GetHeader(), &h)
		uid, err := s.User.Auth(h.Token)
		utils.CheckError(err)
		if cast.ToString(uid) != h.From {
			//return res, req
			utils.CheckError(errors.New("token eror"))
		}
		//return res, req
	}

}

func (s Service) Auth2(next route.HandlerFun) route.HandlerFun {
	fmt.Println("Auth2 s1")
	return func(ctx any) {
		fmt.Println("Auth2 s2")
		h := iface.Header{}
		req, ok := ctx.(*iface.Request)
		if !ok {
			// 失败处理：myInterface 不是 myType 类型
			return
		}

		iface.FromJsonTo(req.GetHeader(), &h)
		uid, err := s.User.Auth(h.Token)
		utils.CheckError(err)
		if err != nil {
			req.GetConnection().SendMsgWithUrl("/chat/ack/err", []byte(err.Error()))
			return
		}
		if cast.ToString(uid) != h.From {
			//return res, req
			utils.CheckError(errors.New("token eror"))
			req.GetConnection().SendMsgWithUrl("/chat/ack/err", []byte("token eror"))
			return
		}
		//return res, req
		next(ctx)
		fmt.Println("Auth2 end")
	}

}
