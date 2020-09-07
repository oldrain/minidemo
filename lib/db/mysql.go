package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"minidemo/constant"
	"minidemo/lib"
	"time"
)

type MysqlCfg struct {
	Url string
	MaxOpenConn int
	MaxIdleConn int
	ConnMaxLifeTime time.Duration
}

func GetDb(connectName string) *gorm.DB {
	cfg := new(MysqlCfg)

	err := lib.CfgBindObj(connectName, cfg)
	if err != nil {
		lib.GetLogger(constant.LogMysql).Error(err)
	}

	db, err := gorm.Open("mysql", cfg.Url)
	if err != nil {
		lib.GetLogger(constant.LogMysql).Error(err)
	}

	db.LogMode(true)

	db.DB().SetMaxOpenConns(cfg.MaxOpenConn)
	db.DB().SetMaxIdleConns(cfg.MaxIdleConn)
	db.DB().SetConnMaxLifetime(cfg.ConnMaxLifeTime * time.Minute)

	return db
}
