package usersmodel

import (
	"minidemo/constant"
	"minidemo/lib"
	"minidemo/model"
	"time"
)

type CustomerToken struct {
	model.BaseModel

	CustomerCode string
	Token        string
	ClientName   string
	LoginTime    time.Time
}

func (model CustomerToken) TableName() string {
	return "customer_token"
}

func CustomerTokenGetByCode(customerCode string) *CustomerToken {
	var m CustomerToken
	err := dbPrimary.Scopes(model.MeshWhere).
		Where("customer_code = ?", customerCode).First(&m)
	if err != nil {
		return nil
	}
	return &m
}

func CustomerTokenCreate(model *CustomerToken) error {
	now := time.Now()
	model.PartnerId = lib.CfgMeshPartnerId()
	model.AppId = lib.CfgMeshAppId()
	model.LoginTime = now
	model.Deleted = constant.No
	model.DeletedAt = nil
	model.CreatedAt = now
	model.CreatedBy = ""
	model.UpdatedAt = now
	model.UpdatedBy = ""
	return dbPrimary.Create(model).Error
}

func CustomerTokenUpdate(model *CustomerToken) error {
	now := time.Now()
	model.LoginTime = now
	model.UpdatedAt = now
	return dbPrimary.Model(model).Where("customer_code = ?", model.CustomerCode).Updates(model).Error
}
