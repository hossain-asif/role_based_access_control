package user

type registerUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type registerUserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type updateUserRequest struct {
	Name  string `json:"username"`
	Email string `json:"email"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}