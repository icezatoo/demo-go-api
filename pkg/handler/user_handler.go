package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	request "github.com/icezatoo/demo-go-api/pkg/dto/user"
	customError "github.com/icezatoo/demo-go-api/pkg/errors"
	"github.com/icezatoo/demo-go-api/pkg/service"
	"github.com/icezatoo/demo-go-api/pkg/utils/formatter"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (u *UserHandler) GetUsers(c *gin.Context) {
	users, err := u.userService.GetUsers()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (u *UserHandler) GetUser(c *gin.Context) {
	var request request.RequestGetUser
	request.ID = c.Param("id")

	if err := c.ShouldBind(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"errors": formatter.NewJSONFormatter().Descriptive(verr)})
			return
		}
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	user, err := u.userService.GetUser(&request)

	if err != nil {
		if ok := customError.IsNotFoundError(err); ok {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": err.Error()})
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var request request.CreateUserRequest

	if err := c.ShouldBind(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"errors": formatter.NewJSONFormatter().Descriptive(verr)})
			return
		}
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	user, err := u.userService.CreateUser(&request)

	if err != nil {
		if ok := customError.IsAlreadyExistsError(err); ok {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"errors": err.Error()})
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	c.JSON(http.StatusCreated, user)
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	var request request.UpdateUserRequest
	request.ID = c.Param("id")

	if err := c.ShouldBind(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"errors": formatter.NewJSONFormatter().Descriptive(verr)})
			return
		}
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	user, err := u.userService.UpdateUser(&request)

	if err != nil {
		if ok := customError.IsNotFoundError(err); ok {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": err.Error()})
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	var request request.RequestDeleteUser
	request.ID = c.Param("id")
	if err := c.ShouldBind(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"errors": formatter.NewJSONFormatter().Descriptive(verr)})
			return
		}
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	err := u.userService.DeleteUser(&request)

	if err != nil {
		if ok := customError.IsNotFoundError(err); ok {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": err.Error()})
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
