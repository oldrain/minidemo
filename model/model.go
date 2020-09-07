package model

import (
	"github.com/jinzhu/gorm"
	"minidemo/constant"
	"minidemo/lib"
	"time"
)

type BaseModel struct {
	//gorm.Model

	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	PartnerId string
	AppId string
	Deleted int
	CreatedBy string
	UpdatedBy string
}

func CloseDB(db *gorm.DB) {
	defer db.Close()
}

func MeshWhere(db *gorm.DB) *gorm.DB {
	return db.Where("partner_id = ? AND app_id = ?", lib.CfgMeshPartnerId(), lib.CfgMeshAppId())
}

func Delete(db *gorm.DB) *gorm.DB {
	return db.UpdateColumn("deleted", constant.Yes)
}

func Restore(db *gorm.DB) *gorm.DB {
	return db.UpdateColumn("deleted", constant.No)
}
