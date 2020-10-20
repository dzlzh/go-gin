package initialize

import (
	"go-gin/global"

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
		panic(err)
	} else {
		global.GVA_LOG.Info("Redis Ping :", pong)
		global.GVA_REDIS = client
	}
}
