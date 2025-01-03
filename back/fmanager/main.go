package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/icehardstone/fmanager/api"
	"github.com/icehardstone/fmanager/config"
	_ "github.com/icehardstone/fmanager/docs"
	"github.com/icehardstone/fmanager/store"
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// @title Swagger API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securitydefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func main() {

	// Init configuration
	err := config.Init()

	if err != nil {
		log.Fatalf("Reading configuration: %s\n", err)
	}

	// database migration

	store.AutoMigrate()

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	router.BasePath()

	router.LoadHTMLGlob(config.AppConfig.Server.Html)
	router.Static("/assets", config.AppConfig.Server.Static)

	router.Use(gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/signin-oidc.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signin-oidc.html", nil)
	})
	router.GET("/helloworld", Helloworld)

	// register file api
	api.RegisterFile(router)

	// register user api
	api.RegisterUser(router)

	// url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	if strings.EqualFold(*config.AppConfig.Environment, "development") {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	srv := &http.Server{
		Addr:    config.AppConfig.Server.Port,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")

}
