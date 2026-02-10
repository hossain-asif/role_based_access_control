package services

import (
	"fmt"
	db "go_project_structure/db/repositories"
)

type UserService interface {
	// Define methods for user service
	CreateUser() error
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (us *UserServiceImpl) CreateUser() error {
	fmt.Println("Creating user in user service.")
	us.userRepository.Create()
	return nil
}
