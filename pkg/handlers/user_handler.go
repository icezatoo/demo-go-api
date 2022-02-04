package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/icezatoo/demo-go-api/pkg/dto/common"
	"github.com/icezatoo/demo-go-api/pkg/services"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (u *UserHandler) GetUsers(c *gin.Context) {
	users, err := u.service.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseError{StatusCode: http.StatusBadRequest, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseSucessful{StatusCode: http.StatusOK, Items: users})

}
