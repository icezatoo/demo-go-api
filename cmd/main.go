package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()

	// router := SetupRouter()

	// log.Fatal(router.Run(":" + utils.GetConfigByKey("GO_PORT")))
}

// func SetupRouter() *gin.Engine {

// 	// utils.LoadConfig()

// 	// db := db.Connection()

// 	// router := gin.Default()

// 	// if utils.GetConfigByKey("GO_ENV") != "production" && utils.GetConfigByKey("GO_ENV") != "test" {
// 	// 	gin.SetMode(gin.DebugMode)
// 	// } else if utils.GetConfigByKey("GO_ENV") == "test" {
// 	// 	gin.SetMode(gin.TestMode)
// 	// } else {
// 	// 	gin.SetMode(gin.ReleaseMode)
// 	// }

// 	// router.Use(cors.New(cors.Config{
// 	// 	AllowOrigins:  []string{"*"},
// 	// 	AllowMethods:  []string{"*"},
// 	// 	AllowHeaders:  []string{"*"},
// 	// 	AllowWildcard: true,
// 	// }))

// 	// router.Use(helmet.Default())

// 	// routers.InitUserRoutes(db, router)

// 	return router
// }
