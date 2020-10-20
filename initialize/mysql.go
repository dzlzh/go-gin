package initialize

import (
	"go-gin/global"

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
		panic(err)
	} else {
		AutoMigrate()
		sqlDB, _ := global.GVA_DB.DB()
		sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
		sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
	}
}

func MysqlClose() {
	db, _ := global.GVA_DB.DB()
	db.Close()
}

func AutoMigrate() {
	err := global.GVA_DB.AutoMigrate()
	if err != nil {
		global.GVA_LOG.Error("MySQL 迁移失败", err)
	}
	global.GVA_LOG.Info("MySQL 迁移成功")
}
