package global

import (
	"go-gin/config"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Service
	GVA_VP     *viper.Viper
	GVA_LOG    *logrus.Logger
)
