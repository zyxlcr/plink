package notification

import "chatcser/pkg/model"

type Notification struct {
	model.BaseModel
	Type         string `json:"type" gorm:"size:16;"`
	From         int64  `json:"from"`
	FromUsername string `json:"from_username"`
	To           int64  `json:"to"`
	Img          string `json:"img"`
	Content      string `json:"content"`
	Do           string `json:"do"  gorm:"size:8;"`
}
