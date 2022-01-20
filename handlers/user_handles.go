package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/icezatoo/demo-go-api/requests"
	"github.com/icezatoo/demo-go-api/services"
	"github.com/icezatoo/demo-go-api/utils"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type UserHandles interface {
	CreateUserHandler(ctx *gin.Context)
	UpdateUserHandler(ctx *gin.Context)
	GetUserList(ctx *gin.Context)
	DeleteUserHandler(ctx *gin.Context)
}

type userHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *userHandler {
	return &userHandler{service: service}
}

func (h *userHandler) GetUserList(ctx *gin.Context) {
	result, errResultsUser := h.service.GetUserList()
	switch errResultsUser {

	case "RESULTS_USER_NOT_FOUND_404":
		utils.APIResponse(ctx, "Users data is not exists", http.StatusNotFound, http.MethodGet, nil)

	default:
		utils.APIResponse(ctx, "Results user data successfully", http.StatusOK, http.MethodGet, result)
	}
}

func (h *userHandler) CreateUserHandler(ctx *gin.Context) {
	var input requests.InputRegisterUser
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Fullname",
				Message: "fullname is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "lowercase",
				Field:   "Fullname",
				Message: "fullname must be using lowercase",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Email",
				Message: "email is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "email",
				Field:   "Email",
				Message: "email format is not valid",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Password",
				Message: "password is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "gte",
				Field:   "Password",
				Message: "password minimum must be 8 character",
			},
		},
	}

	errResponse, errCount := utils.GoValidator(input, config.Options)

	if errCount > 0 {
		utils.ValidatorErrorResponse(ctx, http.StatusUnprocessableEntity, http.MethodPost, errResponse)
		return
	}

	_, err := h.service.CreateUser(&input)

	switch err {
	case "USER_CONFLICT_409":
		utils.APIResponse(ctx, "Email already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "CREATE_USER_FAILED_403":
		utils.APIResponse(ctx, "Register new account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		utils.APIResponse(ctx, "Register new account successfully", http.StatusCreated, http.MethodPost, nil)
	}
}

func (h *userHandler) UpdateUserHandler(ctx *gin.Context) {
	var input requests.InputUpdateUser
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Fullname",
				Message: "fullname is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "lowercase",
				Field:   "Fullname",
				Message: "fullname must be using lowercase",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Email",
				Message: "email is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "email",
				Field:   "Email",
				Message: "email format is not valid",
			},
		},
	}

	errResponse, errCount := utils.GoValidator(input, config.Options)

	if errCount > 0 {
		utils.ValidatorErrorResponse(ctx, http.StatusUnprocessableEntity, http.MethodPut, errResponse)
		return
	}

	result, err := h.service.UpdateUser(&input)

	switch err {
	case "UPDATE_USER_NOT_FOUND_404":
		utils.APIResponse(ctx, "User data is not exist or deleted", http.StatusConflict, http.MethodPut, nil)
		return

	case "UPDATE_USER_FAILED_403":
		utils.APIResponse(ctx, "Update student data failed", http.StatusForbidden, http.MethodPut, nil)
		return

	default:
		utils.APIResponse(ctx, "Update user data successfully", http.StatusCreated, http.MethodPut, result)
	}
}

func (h *userHandler) DeleteUserHandler(ctx *gin.Context) {
	var input requests.InputDeleteUser
	input.ID = ctx.Param("id")

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "ID",
				Message: "id is required on param",
			},
			gpc.ErrorMetaConfig{
				Tag:     "uuid",
				Field:   "ID",
				Message: "params must be uuid format",
			},
		},
	}

	errResponse, errCount := utils.GoValidator(&input, config.Options)

	if errCount > 0 {
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodDelete, errResponse)
		return
	}

	_, errDeleteUser := h.service.DeleteUser(&input)

	switch errDeleteUser {

	case "DELETE_USER_NOT_FOUND_404":
		utils.APIResponse(ctx, "User data is not exist or deleted", http.StatusForbidden, http.MethodDelete, nil)
		return

	case "DELETE_USER_FAILED_403":
		utils.APIResponse(ctx, "Delete user data failed", http.StatusForbidden, http.MethodDelete, nil)
		return

	default:
		utils.APIResponse(ctx, "Delete user data successfully", http.StatusNoContent, http.MethodDelete, nil)
	}
}
