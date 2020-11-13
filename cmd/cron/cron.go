package cron

import (
	"go-gin/global"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func Cron() {
	c := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(cron.VerbosePrintfLogger(zap.NewStdLog(global.GVA_ZAP))),
		cron.WithChain(cron.Recover(cron.DefaultLogger)),
	)

	c.AddJob("*/10 * * * * *", &Example{Message: "秒 分 时 日 月 周"})

	c.Start()
	select {}
}
