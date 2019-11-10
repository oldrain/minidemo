package entity

import "github.com/jinzhu/gorm"

type UserEntity struct {
	gorm.Model
	Name string
	Email string
}

func (entity UserEntity) TableName() string {
	return "user"
}