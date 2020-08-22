package initialize

import (
	"go-gin/global"
	"os"

	"github.com/go-redis/redis"
)

func Redis() {
	conf := global.GVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})

	if pong, err := client.Ping().Result(); err != nil {
		global.GVA_LOG.Error("Redis启动失败", err)
		os.Exit(0)
	} else {
		global.GVA_LOG.Info("redis connect ping response:", pong)
		global.GVA_REDIS = client
	}
}
