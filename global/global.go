package global

import (
	"go-gin/config"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Service
	GVA_VP     *viper.Viper
	GVA_ZAP    *zap.Logger
	GVA_LOG    *zap.SugaredLogger
)
