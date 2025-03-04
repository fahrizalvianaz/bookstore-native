package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-project-native/cmd/myapp1/model/dto"
	"simple-project-native/cmd/myapp1/service"

	"golang.org/x/crypto/bcrypt"
)

type UserHandle struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandle {
	return &UserHandle{service: service}
}

func (u *UserHandle) Register(w http.ResponseWriter, r *http.Request) {
	var req *dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password : ", err)
		return
	}
	req.Password = string(hashedPassword)
	if err := u.service.Register(r.Context(), req); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(req); err != nil {
		fmt.Println("Error encoding response:", err)
	}

}

func (u *UserHandle) Login(w http.ResponseWriter, r *http.Request) {
	var req *dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	passwordUser, err := u.service.Login(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Println(req.Password)
	fmt.Println(passwordUser)
	isMatch := bcrypt.CompareHashAndPassword([]byte(passwordUser), []byte(req.Password))
	if isMatch != nil {
		if isMatch == bcrypt.ErrMismatchedHashAndPassword {
			http.Error(w, "Incorect password or username", http.StatusBadRequest)
		} else {
			http.Error(w, "Error comparing password", http.StatusBadRequest)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var response = map[string]any{
		"code":    http.StatusOK,
		"message": "login succesfully",
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Println("Error encoding response:", err)
	}

}
