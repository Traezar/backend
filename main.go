package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Traezar/go-backend/controllers"
	"github.com/Traezar/go-backend/initializers"
	"github.com/Traezar/go-backend/routes"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

var (
	server           *gin.Engine
	AuthController   controllers.AuthController
	AuthorController controllers.AuthorController
	UserController   controllers.UserController
	BookController   controllers.BookController

	// Route Controller
	AuthRouteController   routes.AuthRouteController
	UserRouteController   routes.UserRouteController
	AuthorRouteController routes.AuthorRouteController
	BookRouteController   routes.BookRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
	initializers.ConnectRedis()

	AuthController = controllers.NewAuthController(initializers.DB)
	UserController = controllers.NewUserController(initializers.DB)
	AuthorController = controllers.NewAuthorController(initializers.DB)
	BookController = controllers.NewBookController(initializers.DB, initializers.Redis)

	AuthRouteController = routes.NewAuthRouteController(AuthController)
	UserRouteController = routes.NewRouteUserController(UserController)
	AuthorRouteController = routes.NewRouteAuthorController(AuthorController)
	BookRouteController = routes.NewRouteBookController(BookController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "? Backend API is up"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	AuthorRouteController.AuthorRoute(router)
	BookRouteController.BookRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
