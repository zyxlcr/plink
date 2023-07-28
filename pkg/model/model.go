package model

import (
	"chatcser/pkg/utils"

	"github.com/tangpanqing/aorm/null"
)

type Model interface {
	Get() int64
}

type BaseModel struct {
	ID        int64          `json:"id" gorm:"primarykey" aorm:"primary;auto_increment"  admin:"disable"`     // 主键ID
	CreatedAt utils.JsonTime `json:"createdAt" gorm:"index;autoCreateTime:true;comment:创建时间" admin:"disable"` // 创建时间
	UpdatedAt utils.JsonTime `json:"updatedAt" gorm:"index;autoCreateTime:true;comment:更新时间" admin:"disable"` // 更新时间
}

type BaseModelAorm struct {
	ID        null.Int  `json:"id" aorm:"primary;auto_increment"  admin:"disable"`   // 主键ID
	CreatedAt null.Time `json:"createdAt" aorm:"index;comment:创建时间" admin:"disable"` // 创建时间
	UpdatedAt null.Time `json:"updatedAt" aorm:"index;comment:更新时间" admin:"disable"` // 更新时间
}

func (m BaseModel) Get() int64 {
	return m.ID
}

func (m BaseModelAorm) Get() null.Int {
	return m.ID
}
