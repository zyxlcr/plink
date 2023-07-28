package app

import (
	"sync"

	"chatcser/pkg/chat"
	"chatcser/pkg/plink/iface"
	"chatcser/pkg/plink/server"
	"chatcser/pkg/user"

	jsoniter "github.com/json-iterator/go"
)

type Service struct {
	Ser    *server.Server
	User   *user.BaseUser
	Chat   *chat.BaseChat
	mu     sync.RWMutex
	TcpMap map[string]iface.IConnection
}

func NewService() Service {
	return Service{Ser: server.NewServer(), User: &user.BaseUser{}, Chat: &chat.BaseChat{}, TcpMap: make(map[string]iface.IConnection)}
}

func (s *Service) ReadTcpMap(key string) iface.IConnection {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.TcpMap[key]
}

func (s *Service) WriteTcpMap(key string, value iface.IConnection) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.TcpMap[key] = value
}

func (s *Service) Run() {
	s.addAllRouter()
	s.Ser.Serve()
}
func (s *Service) ShouldBind(req *iface.Request, st interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(req.GetBody(), st)
	return nil
}
