package user

type LoginReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginRes struct {
	Uid      string `json:"uid" form:"uid" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Token    string `json:"token" form:"token" binding:"required"`
}

type RegisterReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Email    string `json:"email"`
}

type UserInfo struct {
	UserId string `json:"user"`
	Token  string `json:"token"`
}
