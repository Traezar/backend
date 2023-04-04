package routes

import (
	"github.com/Traezar/go-backend/controllers"
	"github.com/Traezar/go-backend/middleware"
	"github.com/gin-gonic/gin"
)

type AuthorRouteController struct {
	authorController controllers.AuthorController
}

func NewRouteAuthorController(authorController controllers.AuthorController) AuthorRouteController {
	return AuthorRouteController{authorController}
}

func (pc *AuthorRouteController) AuthorRoute(rg *gin.RouterGroup) {

	router := rg.Group("/authors")
	router.Use(middleware.DeserializeUser())
	router.POST("/", pc.authorController.CreateAuthor)
	router.GET("/", pc.authorController.FindAuthors)
	router.PUT("/:authorId", pc.authorController.UpdateAuthor)
	router.GET("/:authorId", pc.authorController.FindAuthorById)
	router.DELETE("/:authorId", pc.authorController.DeleteAuthor)
}
