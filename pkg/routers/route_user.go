package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/icezatoo/demo-go-api/pkg/handler"
	"github.com/icezatoo/demo-go-api/pkg/repository"
	"github.com/icezatoo/demo-go-api/pkg/service"
	"gorm.io/gorm"
)

func InitUserRoutes(db *gorm.DB, route *gin.Engine) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	groupRoute := route.Group("/api/v1")
	groupRoute.GET("/users", userHandler.GetUsers)
	groupRoute.GET("/users/:id", userHandler.GetUser)
	groupRoute.POST("/users", userHandler.CreateUser)
	groupRoute.PUT("/users/:id", userHandler.UpdateUser)
	groupRoute.DELETE("/users/:id", userHandler.DeleteUser)
}
