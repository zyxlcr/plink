package user

import (
	"chatcser/config"
	"chatcser/pkg/model"

	"github.com/pkg/errors"
)

func (b BaseUser) Ping() {
	config.GVA_LOG.Info("ping")

}

func (b BaseUser) Search() (BaseUser, error) {
	config.GVA_LOG.Info("Search")
	mapper := model.NewMapper(b, nil)
	user, err := mapper.SelectOne()
	if err != nil {
		config.GVA_LOG.Info("用户不存在")
		return b, errors.Errorf("用户不存在: %s", b.Name)
	}
	config.GVA_LOG.Info(user.Name)
	return user, nil

}
