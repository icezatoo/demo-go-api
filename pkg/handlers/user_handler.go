package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	dto "github.com/icezatoo/demo-go-api/pkg/dto/common"
	request "github.com/icezatoo/demo-go-api/pkg/dto/user"
	"github.com/icezatoo/demo-go-api/pkg/services"
	"github.com/icezatoo/demo-go-api/pkg/utils"
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
		c.JSON(http.StatusBadRequest, dto.ResponseError{StatusCode: http.StatusBadRequest, Errors: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseSucessful{StatusCode: http.StatusOK, Data: users})

}

func (u *UserHandler) GetUser(c *gin.Context) {
	var request request.RequestGetUser
	request.ID = c.Param("id")

	if err := c.ShouldBind(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusUnprocessableEntity, dto.ResponseError{StatusCode: http.StatusUnprocessableEntity, Errors: utils.NewJSONFormatter().Descriptive(verr)})
			return
		}
		c.JSON(http.StatusBadRequest, dto.ResponseError{StatusCode: http.StatusBadRequest, Errors: "Bad request"})
	}

	user, err := u.service.GetUser(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseError{StatusCode: http.StatusBadRequest, Errors: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseSucessful{StatusCode: http.StatusOK, Data: user})
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var request request.CreateUserRequest

	if err := c.ShouldBind(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusUnprocessableEntity, dto.ResponseError{StatusCode: http.StatusUnprocessableEntity, Errors: utils.NewJSONFormatter().Descriptive(verr)})
			return
		}
		c.JSON(http.StatusUnprocessableEntity, dto.ResponseError{StatusCode: http.StatusUnprocessableEntity, Errors: "Bad request"})
		return
	}

	user, err := u.service.CreateUser(&request)

	if err != nil && err.Error() == "USER_CONFLICT_409" {
		c.JSON(http.StatusConflict, dto.ResponseError{StatusCode: http.StatusConflict, Errors: "Email already exists"})
		return
	} else if err != nil {
		logrus.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, dto.ResponseError{StatusCode: http.StatusInternalServerError, Errors: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ResponseSucessful{StatusCode: http.StatusCreated, Data: user})
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	var request request.UpdateUserRequest
	request.ID = c.Param("id")

	if err := c.ShouldBind(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusUnprocessableEntity, dto.ResponseError{StatusCode: http.StatusUnprocessableEntity, Errors: utils.NewJSONFormatter().Descriptive(verr)})
			return
		}
		c.JSON(http.StatusUnprocessableEntity, dto.ResponseError{StatusCode: http.StatusUnprocessableEntity, Errors: "Bad request"})
		return
	}

	user, err := u.service.UpdateUser(&request)

	if err != nil && err.Error() == "USER_NOT_FOUND_404" {
		c.JSON(http.StatusBadRequest, dto.ResponseError{StatusCode: http.StatusBadRequest, Errors: "User not found"})
		return
	} else if err != nil {
		logrus.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, dto.ResponseError{StatusCode: http.StatusInternalServerError, Errors: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseSucessful{StatusCode: http.StatusOK, Data: user})

}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	var request request.RequestDeleteUser
	request.ID = c.Param("id")
	if err := c.ShouldBind(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusUnprocessableEntity, dto.ResponseError{StatusCode: http.StatusUnprocessableEntity, Errors: utils.NewJSONFormatter().Descriptive(verr)})
			return
		}
		c.JSON(http.StatusUnprocessableEntity, dto.ResponseError{StatusCode: http.StatusUnprocessableEntity, Errors: "Bad request"})
		return
	}

	err := u.service.DeleteUser(&request)

	if err != nil && err.Error() == "USER_NOT_FOUND_404" {
		c.JSON(http.StatusBadRequest, dto.ResponseError{StatusCode: http.StatusBadRequest, Errors: "User not found"})
		return
	} else if err != nil {
		logrus.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, dto.ResponseError{StatusCode: http.StatusInternalServerError, Errors: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
