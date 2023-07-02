package app

func (s Service) addAllRouter() {
	s.Ser.AddRouter("/ping", s.Ping)
	s.Ser.Route.Methods("POST").Handler("/ping", s.Ping2)
	s.Ser.Route.Post().Use(s.Auth2).Handler("/chat", s.ChatHandler2)
	s.Ser.Route.Post().Handler("/login", s.Login)
	s.Ser.Route.Post().Handler("/reg", s.Reg)

	// g := s.Ser.Router.GroupWithMore("/chat", s.Auth)
	// g.Post("/send", s.ChatHandler)

	//s.Ser.Router.Use()
}
