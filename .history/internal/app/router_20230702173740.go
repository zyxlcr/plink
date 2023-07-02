package app

func (s Service) addAllRouter() {
	s.Ser.AddRouter("/ping", s.Ping)
	s.Ser.Route.Handler("/ping", s.Ping2)
	s.Ser.Route.Use(s.Auth2).Handler("/chat", s.ChatHandler2)
	s.Ser.Route.Handler("/login", s.Login)
	s.Ser.Route.Handler("/reg", s.Reg)

	// g := s.Ser.Router.GroupWithMore("/chat", s.Auth)
	// g.Post("/send", s.ChatHandler)

	//s.Ser.Router.Use()
}
