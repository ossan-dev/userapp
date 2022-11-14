package dto

import "time"

type UserDTO struct {
	Id      string    `json:"id"`
	AddedOn time.Time `json:"added_on"`
}

func NewUserDTO(id string) *UserDTO {
	return &UserDTO{Id: id}
}
