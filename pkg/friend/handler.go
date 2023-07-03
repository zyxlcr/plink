package friend

import (
	"chatcser/config"
	"chatcser/pkg/model"
	"chatcser/pkg/user"

	"github.com/pkg/errors"
)

func (f Friend) Search(u user.BaseUser) (user.BaseUser, error) {
	config.GVA_LOG.Info("Search")
	mapper := model.NewMapper(u, nil)
	user, err := mapper.SelectOne()
	if err != nil {
		config.GVA_LOG.Info("用户不存在")
		return u, errors.Errorf("用户不存在: %s", u.Name)
	}
	config.GVA_LOG.Info(user.Name)
	return user, nil

}

func (f Friend) MyFriends() (fs Friends, err error) {
	config.GVA_LOG.Info("Search")
	err = config.GVA_DB.Model(&f).Preload("base_user").Select(&fs).Error

	return

}
