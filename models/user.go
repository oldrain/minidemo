package models

import (
	"minidemo/dto/entity"
	"minidemo/lib/db"
)

func UserById(id int, user *entity.UserEntity) (err error) {
	err = db.GetUserMasterDb().Where("id = ?", id).First(user).Error
	return
}
