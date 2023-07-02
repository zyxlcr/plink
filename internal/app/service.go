package app

import (
	"chatcser/pkg/chat"
	"chatcser/pkg/plink/iface"
	"chatcser/pkg/plink/server"
	"chatcser/pkg/user"

	jsoniter "github.com/json-iterator/go"
)

type Service struct {
	Ser  *server.Server
	User *user.BaseUser
	Chat *chat.BaseChat
}

func NewService() Service {
	return Service{server.NewServer(), &user.BaseUser{}, &chat.BaseChat{}}
}

func (s Service) Run() {
	s.addAllRouter()
	s.Ser.Serve()
}
func (s Service) ShouldBind(req *iface.Request, st interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(req.GetBody(), st)
	return nil
}
