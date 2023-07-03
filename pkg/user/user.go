package user

import "chatcser/pkg/model"

type BaseUser struct {
	model.BaseModel
	Name        string `json:"name" gorm:"column:username;not null;comment:用户名"`
	Password    string `json:"password"`
	Email       string `json:"email" gorm:"size:48;"`
	Tel         string `json:"telephone" gorm:"column:telephone;comment:手机号"`
	CanTempChat bool   `json:"can_temp_chat" gorm:"-" default:"false"`
	AvatarUrl   string `json:"avatar_url" gorm:"column:avatar_url;comment:头像"`
}

type UserInfo struct {
	model.BaseModel
	Uid       int64  `json:"uid" gorm:"index;not null;"`
	Gender    uint8  `json:"gender" gorm:"column:gender;not null;default:0;comment:性别"`
	Nickname  string `json:"nickname"`
	Markname  string `json:"markname" gorm:"size:48;"`
	City      uint16 `json:"city" gorm:"column:city;comment:城市"`
	AvatarUrl string `json:"avatar_url" gorm:"column:avatar_url;comment:头像"`
}
