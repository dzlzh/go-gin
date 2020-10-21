package cron

import (
	"go-gin/global"
)

type Example struct {
	Message string
}

func (e *Example) Run() {
	global.GVA_LOG.Info(e.Message)
}
