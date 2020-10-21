package cron

import (
	"go-gin/global"

	"github.com/robfig/cron/v3"
)

func Cron() {
	c := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(cron.VerbosePrintfLogger(global.GVA_LOG)),
		cron.WithChain(cron.Recover(cron.DefaultLogger)),
	)

	c.AddJob("*/10 * * * * *", &Example{Message: "秒 分 时 日 月 周"})

	c.Start()
	select {}
}
