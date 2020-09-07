package usersmodel

import (
	"github.com/jinzhu/gorm"
	"minidemo/constant"
	orm "minidemo/lib/db"
)

var dbPrimary *gorm.DB

func init() {
	dbPrimary = orm.GetDb(constant.CfgMysqlUsersPrimary)
}

func TxBegin() *gorm.DB {
	return dbPrimary.Begin()
}
