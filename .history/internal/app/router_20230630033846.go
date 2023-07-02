package app

func (s Service) addAllRouter() {
	s.Ser.AddRouter("/ping", s.Ping)
	s.Ser.Route.Methods("POST").Handler("/ping", s.Ping2)
	s.Ser.AddRouter("/chat", s.ChatHandler)
	s.Ser.Route.Post().Use(s.Auth2).Handler("/chat", s.ChatHandler2)
	s.Ser.AddRouter("/reg", s.Reg)
	// g := s.Ser.Router.GroupWithMore("/chat", s.Auth)
	// g.Post("/send", s.ChatHandler)

	//s.Ser.Router.Use()
}
