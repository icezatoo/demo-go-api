package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/icezatoo/demo-go-api/pkg/deliveries"
	"github.com/icezatoo/demo-go-api/pkg/repositories"
	"github.com/icezatoo/demo-go-api/pkg/usecases"
	"gorm.io/gorm"
)

func InitUserRoutes(db *gorm.DB, route *gin.Engine) {
	userRepository := repositories.NewUserRepository(db)
	userUseCase := usecases.NewUserUseCase(userRepository)
	userHandler := deliveries.NewUserHandler(userUseCase)

	groupRoute := route.Group("/api/v1")
	groupRoute.GET("/users", userHandler.GetUsers)
	groupRoute.GET("/users/:id", userHandler.GetUser)
	groupRoute.POST("/users", userHandler.CreateUser)
	groupRoute.PUT("/users/:id", userHandler.UpdateUser)
	groupRoute.DELETE("/users/:id", userHandler.DeleteUser)
}
