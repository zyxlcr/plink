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
func (n Notification) All() (ns []Notification) {
	config.GVA_LOG.Info("add")
	//mapper := model.NewMapper(n, nil)
	//var ns []Notification
	config.GVA_DB.Preload("FriendInfo").Where(n).Find(&ns)
	//err := mapper.WhereEq().Select(&n)
	return ns

}

func (n Notification) DoAddFriend(do string) error {
	config.GVA_LOG.Info("DoAddFriend")
	nc := Notification{
		IsDo: 1,
		Do:   do,
	}
	//mapper := model.NewMapper(n, nil)
	//err := mapper.WhereEq("is_do", 0).Updates(nc)
	return config.GVA_DB.Where(n).Where("is_do", 0).Updates(nc).Error

}
