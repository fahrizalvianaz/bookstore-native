package handler

import (
	"net/http"
	"simple-project-native/cmd/myapp1/service"
)

type UserHandle struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandle {
	return &UserHandle{service: service}
}

func (u *UserHandle) Register(w http.ResponseWriter, r *http.Request) {

}

func (u *UserHandle) Login(w http.ResponseWriter, r *http.Request) {

}
