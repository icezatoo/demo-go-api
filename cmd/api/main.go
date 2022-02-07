package main

import (
	"log"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/icezatoo/demo-go-api/pkg/config"
	"github.com/icezatoo/demo-go-api/pkg/routers"
	"github.com/icezatoo/demo-go-api/pkg/utils"
)

func main() {

	router := SetupRouter()
	log.Fatal(router.Run(":" + utils.GetConfigByKey("GO_PORT")))
}

func SetupRouter() *gin.Engine {

	utils.LoadConfig()

	db := config.Connection()

	router := gin.Default()

	if utils.GetConfigByKey("GO_ENV") != "production" && utils.GetConfigByKey("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if utils.GetConfigByKey("GO_ENV") == "test" {
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

	routers.InitUserRoutes(db, router)

	return router
}
