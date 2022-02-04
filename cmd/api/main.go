package main

import "github.com/gin-gonic/gin"

type Response struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, Response{StatusCode: 200, Message: "Test"})
	})
	r.Run()

	// EntityUsers

	// router := SetupRouter()
	// log.Fatal(router.Run(":" + utils.GodotEnv("GO_PORT")))
}

// func SetupRouter() *gin.Engine {

// 	db := config.Connection()

// 	router := gin.Default()

// 	if utils.GodotEnv("GO_ENV") != "production" && utils.GodotEnv("GO_ENV") != "test" {
// 		gin.SetMode(gin.DebugMode)
// 	} else if utils.GodotEnv("GO_ENV") == "test" {
// 		gin.SetMode(gin.TestMode)
// 	} else {
// 		gin.SetMode(gin.ReleaseMode)
// 	}

// 	router.Use(cors.New(cors.Config{
// 		AllowOrigins:  []string{"*"},
// 		AllowMethods:  []string{"*"},
// 		AllowHeaders:  []string{"*"},
// 		AllowWildcard: true,
// 	}))
// 	router.Use(helmet.Default())

// 	/**
// 	@description Init All Route
// 	*/
// 	route.InitUserRoutes(db, router)
// 	route.InitAuthRoutes(db, router)

// 	return router
// }
