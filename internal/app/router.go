package app

func (s Service) addAllRouter() {
	s.Ser.AddRouter("/ping", s.Ping)

	s.Ser.Route.Handler("/ping", s.Ping2)

	s.Ser.Route.Use(s.Auth2).Handler("/chat", s.ChatHandler2)

	s.Ser.Route.Handler("/login", s.Login)
	s.Ser.Route.Handler("/reg", s.Reg)

	s.Ser.Route.Handler("/user/search", s.SearchUser)

	s.Ser.Route.Use(s.Auth2).Handler("/notification/send", s.Send)
	s.Ser.Route.Use(s.Auth2).Handler("/notification/doAddfriend", s.DoNotification)

	s.Ser.Route.Use(s.Auth2).Handler("/friend/myfriend", s.MyFriends)

}
