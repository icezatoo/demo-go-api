package route

import (
	"github.com/gin-gonic/gin"
	"github.com/icezatoo/demo-go-api/handlers"
	"github.com/icezatoo/demo-go-api/repository"
	"github.com/icezatoo/demo-go-api/services"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Auth
	*/
	authRepository := repository.NewRepositoryAuth(db)
	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandles(authService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/login", authHandler.Login)
}
