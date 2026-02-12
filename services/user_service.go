package services

import (
	"fmt"
	db "go_project_structure/db/repositories"
)

type UserService interface {
	// Define methods for user service
	CreateUser(username string, email string, password string) error
	GetUserById(id string) error
	GetAllUsers() error
	UpdateUser(id string, username string, email string) error
	DeleteUser(id string) error
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (us *UserServiceImpl) CreateUser(username string, email string, password string) error {
	fmt.Println("Creating user in user service.")
	err := us.userRepository.Create(username, email, password)
	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		return err
	}
	return nil
}

func (us *UserServiceImpl) GetUserById(id string) error {
	fmt.Println("Getting user by id in user service.")
	us.userRepository.GetByID(id)
	return nil
}

func (us *UserServiceImpl) GetAllUsers() error {
	fmt.Println("Getting all users in user service.")
	us.userRepository.GetAll()
	return nil
}

func (us *UserServiceImpl) UpdateUser(id string, username string, email string) error {
	fmt.Println("Updating user in user service.")
	us.userRepository.Update(id, username, email)
	return nil
}

func (us *UserServiceImpl) DeleteUser(id string) error {
	fmt.Println("Deleting user in user service.")
	us.userRepository.Delete(id)
	return nil
}
