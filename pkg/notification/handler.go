package notification

import (
	"chatcser/config"
	"chatcser/pkg/model"
)

func (b Notification) Ping() {
	config.GVA_LOG.Info("ping")

}

func (n Notification) Add() error {
	config.GVA_LOG.Info("add")
	mapper := model.NewMapper(n, nil)
	err := mapper.Insert(&n)
	return err

}

func (n Notification) DoAddFriend(do string) error {
	config.GVA_LOG.Info("add")
	mapper := model.NewMapper(n, nil)
	err := mapper.Update("do", do)
	return err

}
