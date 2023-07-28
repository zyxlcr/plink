package config

import (
	"chatcser/config/autoload"

	"github.com/spf13/viper"
	//"github.com/tangpanqing/aorm"
	red "github.com/redis/go-redis/v9"
	"github.com/tangpanqing/aorm/base"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Configuration struct {
	Domain string         `mapstructure:"domain" json:"domain" yaml:"domain"`
	DbType string         `mapstructure:"dbType" json:"dbType" yaml:"dbType"`
	Admin  autoload.Admin `mapstructure:"admin" json:"admin" yaml:"admin"`
	Mysql  autoload.Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Zap    autoload.Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	JWT    autoload.JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis  autoload.Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
}

var (
	GVA_CONFIG Configuration
	GVA_DB     *gorm.DB
	GVA_LOG    *zap.Logger
	GVA_VP     *viper.Viper
	GVA_REDIS  *red.Client
	GVA_AORM   *base.Db
)
