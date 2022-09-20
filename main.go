package main

import (
	database "eCommerce/config"
	"eCommerce/controller"
	_ "eCommerce/docs"
	"eCommerce/infrastructure/middleware"
	"eCommerce/repository"
	"eCommerce/service"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	port = 8080
)

// @title          eCommerce API
// @version        1.0
// @description    An eCommerce service API in Go using Gin framework.
// @termsOfService https://tos.santoshk.dev
// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html
// @host           localhost:8080
// @BasePath       /api
func main() {
	// Initialize log
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	//Initialize database
	db := database.Init()
	defer database.Dispose(db)

	// Initialize Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.New() // empty engine
	r.Use(cors.Default())
	r.Use(middleware.SetCorrelationID())         // adds correlation middleware
	r.Use(middleware.SetJSONContentTypeHeader()) // adds setting json content type header middleware
	r.Use(middleware.Logger())                   // adds logger middleware
	r.Use(middleware.ErrorHandler())             // adds error handler middleware
	r.Use(gin.Recovery())                        // adds the default recovery middleware

	// Routes
	r.GET("/", HealthCheck)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register all repositories
	personRepository := repository.NewPersonRepository(db)

	// Register all services
	personService := service.NewPersonService(personRepository)

	// Register all controllers
	personController := controller.NewPersonController(personService)

	// Register routers
	basePath := r.Group("/api")
	personController.RegisterRoutes(basePath)

	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Error(err)
	}
}

func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "healthy",
	}
	c.JSON(http.StatusOK, res)
}
