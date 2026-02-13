package user

import (
	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *UserController
}

func NewUserRouter(_userController *UserController) *UserRouter {
	return &UserRouter{
		userController: _userController,
	}
}

func (ur *UserRouter) Register(r chi.Router) {
	r.Post("/register", ur.userController.RegisterUser)
	r.Get("/profile/{id}", ur.userController.GetUserById)
	r.Get("/profile", ur.userController.GetAllUsers)
	r.Patch("/profile/{id}", ur.userController.UpdateUser)
	r.Delete("/profile/{id}", ur.userController.DeleteUser)
}
