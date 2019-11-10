package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"minidemo/dto"
	"minidemo/lib"
	"minidemo/constant"
	"time"
)

func GetUserMasterDb() *gorm.DB {
	cfg := new(dto.MysqlCfg)

	err := lib.CfgBindObj(constant.ConfigMysqlUserMaster, cfg)
	if err != nil {
		lib.GetLogger(constant.LogMysql).Error(err)
	}

	db, err := gorm.Open("mysql", cfg.Url)

	if err != nil {
		lib.GetLogger(constant.LogMysql).Error(err)
	}

	db.DB().SetMaxOpenConns(cfg.MaxOpenConn)
	db.DB().SetMaxIdleConns(cfg.MaxIdleConn)
	db.DB().SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifeTime * time.Minute))

	return db
}
