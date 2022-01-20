package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/icezatoo/demo-go-api/requests"
	"github.com/icezatoo/demo-go-api/services"
	"github.com/icezatoo/demo-go-api/utils"
	gpc "github.com/restuwahyu13/go-playground-converter"
	"github.com/sirupsen/logrus"
)

type AuthHandles interface {
	Login(ctx *gin.Context)
}

type authHandles struct {
	service services.AuthService
}

func NewAuthHandles(service services.AuthService) *authHandles {
	return &authHandles{service: service}
}

func (h *authHandles) Login(ctx *gin.Context) {
	var input requests.InputLogin
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
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
		},
	}

	errResponse, errCount := utils.GoValidator(&input, config.Options)
	if errCount > 0 {
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultLogin, errLogin := h.service.LoginService(&input)

	switch errLogin {

	case "LOGIN_NOT_FOUND_404":
		utils.APIResponse(ctx, "User account is not registered", http.StatusNotFound, http.MethodPost, nil)
		return

	case "LOGIN_NOT_ACTIVE_403":
		utils.APIResponse(ctx, "User account is not active", http.StatusForbidden, http.MethodPost, nil)
		return

	case "LOGIN_WRONG_PASSWORD_403":
		utils.APIResponse(ctx, "Username or password is wrong", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		accessTokenData := map[string]interface{}{"id": resultLogin.ID, "email": resultLogin.Email}
		accessToken, errToken := utils.Sign(accessTokenData, "JWT_SECRET", 24*60*1)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			utils.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		utils.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, map[string]string{"accessToken": accessToken})
	}
}
