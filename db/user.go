package db

import (
	"gorm.io/gorm"

	"userapp/models"
)

type UserDb struct {
	Conn *gorm.DB
}

type UserDbInterface interface {
	SaveUser(user *models.User) error
}

func (u *UserDb) SaveUser(user *models.User) error {
	if dbTrn := u.Conn.Create(user); dbTrn.Error != nil {
		return dbTrn.Error
	}
	return nil
}
