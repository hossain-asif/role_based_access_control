package db

// facilates dependency injection for repositories
type Storage struct { 
	UserRepository UserRepository
}

func NewStorage() *Storage {
	return &Storage{
		UserRepository: &UserRepositoryImpl{},
	}
} 