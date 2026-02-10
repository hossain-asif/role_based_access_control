package db

import (
	"fmt"
)

type UserRepository interface {
	// Define methods for user repository
	Create() error
}

type UserRepositoryImpl struct {
	// Add fields for database connection, etc.
	// db *sql.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{
		// db: db,
	}
}

func (r *UserRepositoryImpl) Create() error {
	fmt.Println("creating user in user repository.")
	return nil
}
