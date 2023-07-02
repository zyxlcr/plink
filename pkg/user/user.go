package user

import "chatcser/pkg/model"

type BaseUser struct {
	model.BaseModel
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Tel      string `json:"telephone"`
}
