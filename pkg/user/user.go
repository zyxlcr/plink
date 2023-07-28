package user

import (
	"chatcser/pkg/model"

	"github.com/tangpanqing/aorm/null"
)

type BaseUser struct {
	model.BaseModel
	Name        string `json:"name" gorm:"column:username;not null;comment:用户名"`
	Password    string `json:"password"`
	Email       string `json:"email" gorm:"size:48;"`
	Tel         string `json:"telephone" gorm:"column:telephone;comment:手机号"`
	CanTempChat bool   `json:"can_temp_chat" gorm:"-" default:"false"`
	AvatarUrl   string `json:"avatar_url" gorm:"column:avatar_url;comment:头像"`
}
type BaseUserAorm struct {
	Id        null.Int  `json:"id" aorm:"primary;auto_increment"  admin:"disable"`   // 主键ID
	CreatedAt null.Time `json:"createdAt" aorm:"index;comment:创建时间" admin:"disable"` // 创建时间
	UpdatedAt null.Time `json:"updatedAt" aorm:"index;comment:更新时间" admin:"disable"` // 更新时间

	Name     null.String `json:"name" aorm:"column:username;not null;comment:用户名"`
	Password null.String `json:"password"`
	Email    null.String `json:"email" aorm:"size:48"`
	Tel      null.String `json:"telephone" aorm:"column:telephone;comment:手机号"`
	//CanTempChat bool        `json:"can_temp_chat" aorm:"-" default:"false"`
	AvatarUrl null.String `json:"avatar_url" aorm:"column:avatar_url;comment:头像"`
}

// 修改默认表名
// func (p *BaseUserAorm) TableName() string {
// 	//aorm.Store(p)

// 	t := reflect.TypeOf(p)

// 	if t.Kind() == reflect.Ptr {
// 		t = t.Elem()
// 	}

// 	config.GVA_LOG.Debug(t.Name())
// 	c := config.GVA_CONFIG.Mysql
// 	return c.Prefix + utils.ConvertString(t.Name()) //"base_user"
// }

// 可以定义该函数来设置表信息
func (p *BaseUser) TableOpinion() map[string]string {
	return map[string]string{
		"ENGINE":  "InnoDB",
		"COMMENT": "用户表",
	}
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

func (p *UserInfo) TableOpinion() map[string]string {
	return map[string]string{
		"ENGINE":  "InnoDB",
		"COMMENT": "用户信息表",
	}
}
