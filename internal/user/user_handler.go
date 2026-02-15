package user

import (
	utils "go_project_structure/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	UserService UserService
}

func NewUserController(_userService UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {

	RequestPayload := r.Context().Value("registration_payload").(registerUserRequest)

	err := uc.UserService.CreateUser(
		RequestPayload.Name,
		RequestPayload.Email, 
		RequestPayload.Password,
	)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "User registration failed.", err)
		return
	}
	responsePayload := registerUserResponse{
		Name:  RequestPayload.Name,
		Email: RequestPayload.Email,
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User registration suucessful", responsePayload)
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload = LoginUserRequest{}
	err := utils.ReadJsonBody(r, &requestPayload)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request payload", err)
		return
	}
	
	token, err := uc.UserService.LoginUser(requestPayload.Email, requestPayload.Password)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusUnauthorized, "Login failed", err)
		return
	}
	responsePayload := LoginUserResponse{
		Token: token,
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Login successful", responsePayload)
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	uc.UserService.GetUserById(userId)

	response := map[string]interface{}{
		"success": true,
		"message": "Get user by id end point",
		"data":    nil,
		"error":   nil,
	}
	utils.WriteJSONResponse(w, http.StatusOK, response)

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
