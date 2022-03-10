package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/icezatoo/demo-go-api/pkg/config"
	"github.com/icezatoo/demo-go-api/pkg/db"
	"github.com/icezatoo/demo-go-api/pkg/routers"
	"github.com/sirupsen/logrus"
)

func main() {
	config := config.LoadConfigENV()

	router := SetupRouter(config)

	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			logrus.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Server forced to shutdown:", err)
	}

	logrus.Println("Server exiting")

}

func SetupRouter(config *config.Config) *gin.Engine {

	db := db.Connection(config)

	router := gin.Default()

	if config.Environment != "production" && config.Environment != "test" {
		gin.SetMode(gin.DebugMode)
	} else if config.Environment == "test" {
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
