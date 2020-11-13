package cron

import (
	"go-gin/global"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func Cron() {
	stdLog := zap.NewStdLog(global.GVA_ZAP)
	c := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(cron.PrintfLogger(stdLog)),
		cron.WithChain(cron.Recover(cron.VerbosePrintfLogger(stdLog))),
	)

	// c.AddJob("*/10 * * * * *", &Example{Message: "秒 分 时 日 月 周"})

	c.Start()
	select {}
}
