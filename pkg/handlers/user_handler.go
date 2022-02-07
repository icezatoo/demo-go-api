package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/icezatoo/demo-go-api/pkg/dto/common"
	request "github.com/icezatoo/demo-go-api/pkg/dto/user"
	"github.com/icezatoo/demo-go-api/pkg/services"
	"github.com/sirupsen/logrus"
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

	c.JSON(http.StatusOK, dto.ResponseSucessful{StatusCode: http.StatusOK, Data: users})

}

func (u *UserHandler) GetUser(c *gin.Context) {
	var request request.RequestGetUser
	request.ID = c.Param("id")

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseError{StatusCode: http.StatusBadRequest, Error: err.Error()})
	}

	user, err := u.service.GetUser(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseError{StatusCode: http.StatusBadRequest, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseSucessful{StatusCode: http.StatusOK, Data: user})
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var request request.CreateUserRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.ResponseError{StatusCode: http.StatusUnprocessableEntity, Error: "Invalid request"})
		return
	}

	user, err := u.service.CreateUser(&request)

	if err != nil && err.Error() == "USER_CONFLICT_409" {
		c.JSON(http.StatusConflict, dto.ResponseError{StatusCode: http.StatusConflict, Error: "Email already exists"})
		return
	} else if err != nil {
		logrus.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, dto.ResponseError{StatusCode: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ResponseSucessful{StatusCode: http.StatusCreated, Data: user})
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	var request request.RequestDeleteUser

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.ResponseError{StatusCode: http.StatusUnprocessableEntity, Error: "Invalid request"})
		return
	}

	err := u.service.DeleteUser(&request)

	if err != nil && err.Error() == "USER_NOT_FOUND_404" {
		c.JSON(http.StatusBadRequest, dto.ResponseError{StatusCode: http.StatusBadRequest, Error: "User not found"})
		return
	} else if err != nil {
		logrus.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, dto.ResponseError{StatusCode: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
