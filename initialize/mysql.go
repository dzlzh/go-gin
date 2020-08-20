package initialize

import (
	"go-gin/global"
	"os"

	"github.com/jinzhu/gorm"
)

func Mysql() {
	conf := global.GVA_CONFIG.Mysql
	if db, err := gorm.Open("mysql", conf.Username+":"+conf.Password+"@("+conf.Addr+")/"+conf.Database+"?"+conf.Config); err != nil {
		global.GVA_LOG.Error("MySQL启动失败", err)
		os.Exit(0)
	} else {
		global.GVA_DB = db
		global.GVA_DB.DB().SetMaxIdleConns(conf.MaxIdleConns)
		global.GVA_DB.DB().SetMaxOpenConns(conf.MaxOpenConns)
		global.GVA_DB.LogMode(conf.LogMode)
		if conf.Prefix != "" {
			gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
				return conf.Prefix + defaultTableName
			}
		}
	}
}
