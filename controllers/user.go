package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"userapp/dto"
	"userapp/services"
)

type UserController struct {
	US services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) *UserController {
	return &UserController{
		US: userService,
	}
}

func (u *UserController) Save(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	r.Body.Close()

	var userReq dto.UserDTO

	json.Unmarshal(reqBody, &userReq)

	userRes, err := u.US.AddUser(&userReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userRes)
}
