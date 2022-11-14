package models

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Id string `gorm:"primaryKey"`
}

func NewUser(id string) *User {
	return &User{Id: id}
}
