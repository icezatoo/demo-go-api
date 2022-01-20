package route

import (
	"github.com/gin-gonic/gin"
	"github.com/icezatoo/demo-go-api/handlers"
	middleware "github.com/icezatoo/demo-go-api/middlewares"
	"github.com/icezatoo/demo-go-api/repository"
	"github.com/icezatoo/demo-go-api/services"
	"gorm.io/gorm"
)

func InitUserRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Auth
	*/
	userRepository := repository.NewRepositoryUser(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())

	groupRoute.GET("/users", userHandler.GetUserList)
	groupRoute.POST("/users", userHandler.CreateUserHandler)
	groupRoute.PUT("/users", userHandler.UpdateUserHandler)
	groupRoute.PUT("/users/:id", userHandler.DeleteUserHandler)
}
