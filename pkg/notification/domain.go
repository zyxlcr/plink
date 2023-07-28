package notification

type NobodyReq struct {
}
type BaseRes struct {
	Data any    `json:"data" form:"data" binding:"required"`
	Msg  string `json:"msg" form:"msg" binding:"required"`
	Code int8   `json:"code" form:"code" binding:"required"`
}

type SendReq struct {
	Content string `json:"content" form:"content" binding:"required"`
	Img     string `json:"img" form:"img"`
}

type SendRes struct {
	Mid int64  `json:"mid" form:"mid" binding:"required"`
	Do  string `json:"do" form:"do" binding:"required"`
	Url string `json:"url" form:"url" binding:"required"`
}

type DoAddFriendReq struct {
	Mid int64  `json:"mid" form:"mid" binding:"required"`
	Do  string `json:"do" form:"do" binding:"required"`
}
