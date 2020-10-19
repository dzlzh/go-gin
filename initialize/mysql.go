package initialize

import (
	"go-gin/global"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Mysql() {
	var err error
	conf := global.GVA_CONFIG.Mysql
	dsn := conf.Username + ":" + conf.Password + "@(" + conf.Addr + ")/" + conf.Database + "?" + conf.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}
	logMode := logger.Silent
	if conf.LogMode {
		logMode = logger.Info
	}
	gormConfig := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logMode),
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: conf.Prefix,
		},
	}
	if global.GVA_DB, err = gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		global.GVA_LOG.Error("MySQL启动失败", err)
		os.Exit(0)
	} else {
		sqlDB, _ := global.GVA_DB.DB()
		sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
		sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
	}
}
