package services

import (
	"time"

	"userapp/db"
	"userapp/dto"
	"userapp/models"
)

type UserService struct {
	DB db.UserDbInterface
}

type UserServiceInterface interface {
	AddUser(inputReq *dto.UserDTO) (*dto.UserDTO, error)
}

func NewUserService(db db.UserDbInterface) *UserService {
	return &UserService{
		DB: db,
	}
}

func (u *UserService) AddUser(inputReq *dto.UserDTO) (*dto.UserDTO, error) {
	// here you can write complex logic
	user := models.NewUser(inputReq.Id)

	// invoke db repo
	if err := u.DB.SaveUser(user); err != nil {
		return nil, err
	}

	inputReq.AddedOn = time.Now()

	return inputReq, nil
}
