package usersmodel

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
	Email string
}

func (entity User) TableName() string {
	return "user"
}

func UserGetById(id int) *User {
	var user *User
	notFound := dbPrimary.Where("id = ?", id).First(&user).RecordNotFound()
	if notFound {
		return nil
	}
	return user
}
