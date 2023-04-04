package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Traezar/go-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BookController struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func NewBookController(DB *gorm.DB, Redis *redis.Client) BookController {
	return BookController{DB, Redis}
}

// Create Book Handler
func (bc *BookController) CreateBook(ctx *gin.Context) {

	var payload *models.CreateBookRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newBook := models.Book{
		Title:     payload.Title,
		AuthorID:  payload.Author.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := bc.DB.Create(&newBook)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Author with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newBook})
}

// [...] Update Book Handler
func (pc *BookController) UpdateBook(ctx *gin.Context) {
	bookId := ctx.Param("bookId")

	var payload *models.UpdateBook
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedBook models.Book
	result := pc.DB.First(&updatedBook, "id = ?", bookId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No book with that title exists"})
		return
	}
	now := time.Now()
	bookToUpdate := models.Book{
		Title:     payload.Title,
		Author:    payload.Author,
		CreatedAt: updatedBook.CreatedAt,
		UpdatedAt: now,
	}

	pc.DB.Model(&updatedBook).Updates(bookToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedBook})
}

func (pc *BookController) ListBooks(ctx *gin.Context) {
	var books []models.Book
	result := pc.DB.Preload("Author").Find(&books)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": books})
}
func (pc *BookController) ListBooksTwo(ctx *gin.Context) {
	// First, check if the data is in the cache
	cachedData, err := pc.Redis.Get(ctx, "books").Result()
	if err == nil {
		// Data was found in the cache, return it
		var books []models.Book
		err = json.Unmarshal([]byte(cachedData), &books)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": books})
		return
	}

	// Data was not found in the cache, fetch it from the database
	var books []models.Book
	result := pc.DB.Preload("Author").Find(&books)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	// Store the fetched data in the cache for future requests
	booksData, err := json.Marshal(books)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	err = pc.Redis.Set(ctx, "books", string(booksData), time.Minute*5).Err()
	if err != nil {
		// If there was an error storing the data in the cache, log it but continue
		log.Println("Error storing books data in cache:", err)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": books})
}
