package routes

import (
	"github.com/Traezar/go-backend/controllers"
	"github.com/Traezar/go-backend/middleware"
	"github.com/gin-gonic/gin"
)

type BookRouteController struct {
	bookController controllers.BookController
}

func NewRouteBookController(bookController controllers.BookController) BookRouteController {
	return BookRouteController{bookController}
}

func (bc *BookRouteController) BookRoute(rg *gin.RouterGroup) {

	router := rg.Group("/books")
	router.Use(middleware.DeserializeUser())
	router.POST("/", bc.bookController.CreateBook)
	router.GET("/", bc.bookController.ListBooks)
	//router.GET("/", bc.bookController.ListBooksTwo)
	router.PUT("/:bookId", bc.bookController.UpdateBook)
	// router.GET("/:bookId", bc.bookController.FindBookById)
	// router.DELETE("/:bookId", bc.bookController.DeleteBook)
}
