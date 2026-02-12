package controllers

import (
	"go_project_structure/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	uc.UserService.CreateUser("alicebob", "alicebob@example.com", "password123")
	w.Write([]byte("User registration end point"))
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	uc.UserService.GetUserById(userId)
	w.Write([]byte("User get by id end point"))
}

func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	uc.UserService.GetAllUsers()
	w.Write([]byte("Get all users end point"))
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	uc.UserService.UpdateUser(userId, "newusername", "newemail@example.com")
	w.Write([]byte("User update end point"))
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	uc.UserService.DeleteUser(userId)
	w.Write([]byte("User delete end point"))
}
