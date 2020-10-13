package initialize

import (
	"fmt"
	"go-gin/global"
	"path"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func Log() {
	conf := global.GVA_CONFIG.Log
	global.GVA_LOG = logrus.New()

	var loglevel logrus.Level
	loglevel.UnmarshalText([]byte(conf.Level))
	global.GVA_LOG.SetLevel(loglevel)

	// 为不同级别设置不同的输出目的
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer(conf.Path, "debug"),
		logrus.InfoLevel:  writer(conf.Path, "info"),
		logrus.WarnLevel:  writer(conf.Path, "warn"),
		logrus.ErrorLevel: writer(conf.Path, "error"),
		logrus.FatalLevel: writer(conf.Path, "fatal"),
		logrus.PanicLevel: writer(conf.Path, "panic"),
	}, &MineFormatter{})

	global.GVA_LOG.AddHook(lfHook)

}

func writer(logPath string, level string) *rotatelogs.RotateLogs {
	logFullPath := path.Join(logPath, level)
	fileSuffix := time.Now().Format("20060102")

	writer, err := rotatelogs.New(
		logFullPath+"."+fileSuffix+".log",
		rotatelogs.WithLinkName(logFullPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	if err != nil {
		panic(err)
	}

	return writer
}

type MineFormatter struct{}

func (s *MineFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	msg := fmt.Sprintf("[%s] [%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}
