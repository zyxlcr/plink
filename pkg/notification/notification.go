package notification

import (
	"chatcser/pkg/model"
	"chatcser/pkg/user"
)

type Notification struct {
	model.BaseModel
	Type         string        `json:"type" gorm:"size:16;"`
	From         int64         `json:"from"`
	FromUsername string        `json:"from_username"`
	To           int64         `json:"to"`
	Img          string        `json:"img"`
	Content      string        `json:"content"`
	IsDo         int8          `json:"is_do" gorm:"size:1;default:-1;"`
	Do           string        `json:"do"  gorm:"size:8;"`
	FriendInfo   user.BaseUser `json:"friend_info" gorm:"foreignKey:id;references:from"`
}
