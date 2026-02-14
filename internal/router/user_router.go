package router

import (
	"go_project_structure/internal/user"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type UserRouter struct {
	userController *user.UserController
}

func NewUserRouter(_userController *user.UserController) *UserRouter {
	return &UserRouter{
		userController: _userController,
	}
}

func RegisterRoutes(db *gorm.DB, router chi.Router) *UserRouter {
	ur := user.NewUserRepository(db)
	us := user.NewUserService(ur)
	uc := user.NewUserController(us)
	uRouter := NewUserRouter(uc)
	return uRouter
}

func (ur *UserRouter) Register(r chi.Router) {
	r.Use(user.RequestLoggerMiddleware)
	r.With(user.UserRegisterRequestValidator).Post("/register", ur.userController.RegisterUser)
	r.Get("/profile/{id}", ur.userController.GetUserById)
	r.Get("/profile", ur.userController.GetAllUsers)
	r.With(user.UserUpdateRequestValidator).Patch("/profile/{id}", ur.userController.UpdateUser)
	r.Delete("/profile/{id}", ur.userController.DeleteUser)
}
