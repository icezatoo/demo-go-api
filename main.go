package main

import (
	"log"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	config "github.com/icezatoo/demo-go-api/configs"
	route "github.com/icezatoo/demo-go-api/routes"
	"github.com/icezatoo/demo-go-api/utils"
)

func main() {

	router := SetupRouter()

	log.Fatal(router.Run(":" + utils.GodotEnv("GO_PORT")))
}

func SetupRouter() *gin.Engine {

	db := config.Connection()

	router := gin.Default()

	if utils.GodotEnv("GO_ENV") != "production" && utils.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if utils.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	router.Use(helmet.Default())

	/**
	@description Init All Route
	*/
	route.InitUserRoutes(db, router)
	route.InitAuthRoutes(db, router)

	return router
}
