package chat

type ToChatReq struct {
	Uid         string `json:"uid" form:"uid" binding:"required"`
	MsgType     string `json:"msg_type" form:"msg_type" binding:"required"`
	ContentType string `json:"content_type" form:"content_type" binding:"required"`
	IsMe        bool   `json:"is_me" form:"is_me" binding:"required"`
	FriendId    string `json:"friend_id" form:"friend_id" binding:"required"`
	FriendName  string `json:"friend_name" form:"friend_name" binding:"required"`
	UpdateAt    string `json:"update_at" form:"update_at" binding:"required"`
	Content     string `json:"content" form:"content" binding:"required"`
	AvatarURL   string `json:"avatar_url" form:"avatar_url" binding:"required"`
}
