package friend

type SearchReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Uid      string `json:"uid" form:"uid" binding:"required"`
	Tel      string `json:"tel" form:"tel" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
}

type SearchRes struct {
	Username  string `json:"username" form:"username" binding:"required"`
	Uid       string `json:"uid" form:"uid" binding:"required"`
	AvatarUrl string `json:"avatar_url" form:"tel" binding:"required"`
	Tel       string `json:"tel" form:"tel" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required"`
}

type MyFriendRes struct {
	Data Friends `json:"data" form:"data" binding:"required"`
	Code int8    `json:"code" form:"code" binding:"required"`
}
