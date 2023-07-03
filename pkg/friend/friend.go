package friend

import (
	"chatcser/pkg/model"
	"chatcser/pkg/user"
)

type Friend struct {
	model.BaseModel
	Uid        uint64        `json:"uid" gorm:"index;comment:用户id" form:"uid" binding:"required"`
	FriendId   uint64        `json:"friend_id" gorm:"index;comment:好友id" form:"friend_id" binding:"required"`
	IsDel      int8          `json:"is_del" gorm:"comment:是否删除" form:"is_del" binding:"required"`
	FriendInfo user.BaseUser `gorm:"foreignKey:id;references:friend_id"`
}
type Friends []Friend
