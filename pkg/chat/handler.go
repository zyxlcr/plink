package chat

import (
	"chatcser/config"
	"chatcser/pkg/user"
	"fmt"

	"github.com/tangpanqing/aorm"
	"github.com/tangpanqing/aorm/builder"
)

func (b BaseChat) ChatTest() {
	config.GVA_LOG.Info("chat")
	u := user.BaseUserAorm{}
	aorm.Store(&u)
	var personItem user.BaseUserAorm
	errFind := aorm.Db(config.GVA_AORM).Table(&u).OrderBy(&u.Email, builder.Desc).WhereEq(&u.Id, 1).GetOne(&personItem)
	if errFind != nil {
		config.GVA_LOG.Info(errFind.Error())
	}
	config.GVA_LOG.Info(u.Name.String)
	fmt.Println(u)
}
