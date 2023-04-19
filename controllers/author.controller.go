package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Traezar/go-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthorController struct {
	DB *gorm.DB
}

func NewAuthorController(DB *gorm.DB) AuthorController {
	return AuthorController{DB}
}

// Create Author Handler
// [...] Create Author Handler
func (pc *AuthorController) CreateAuthor(ctx *gin.Context) {
	var payload *models.CreateAuthorRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newAuthor := models.Author{
		Name:      payload.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := pc.DB.Create(&newAuthor)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Author with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newAuthor})
}

// [...] Update Author Handler
func (pc *AuthorController) UpdateAuthor(ctx *gin.Context) {
	fmt.Printf("%+v", ctx.Params)
	AuthorId := ctx.Param("authorId")

	var payload *models.UpdateAuthor
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedAuthor models.Author
	result := pc.DB.First(&updatedAuthor, "id::text= ?", AuthorId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No author with that name exists"})
		return
	}
	now := time.Now()
	author := models.Author{
		Name:      payload.Name,
		CreatedAt: updatedAuthor.CreatedAt,
		UpdatedAt: now,
	}

	pc.DB.Model(&updatedAuthor).Updates(author)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedAuthor})
}

func (pc *AuthorController) FindAuthorById(ctx *gin.Context) {
	authorId := ctx.Param("authorId")

	var author models.Author
	result := pc.DB.First(&author, "id = ?", authorId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No author with that id exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": author})
}

func (pc *AuthorController) FindAuthorByName(ctx *gin.Context) {
	name := ctx.Param("name")

	var author models.Author
	result := pc.DB.First(&author, "name = ?", name)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No author with that id exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": author})
}

func (pc *AuthorController) FindAuthors(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var authors []models.Author
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&authors)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(authors), "data": authors})
}

func (pc *AuthorController) DeleteAuthor(ctx *gin.Context) {
	authorId := ctx.Param("authorId")

	result := pc.DB.Delete(&models.Author{}, "id = ?", authorId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No author with that name exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
