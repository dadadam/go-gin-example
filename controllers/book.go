package controllers

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	dbc "github.com/dadadam/sono-backend/db"
	"github.com/dadadam/sono-backend/models"
	"github.com/gin-gonic/gin"
)

type BookCreateScheme struct {
	Title       string `json:"title" binding:"required"`
	AuthorID    uint   `json:"authorId" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type BookController struct{}

func (c BookController) Create(ctx *gin.Context) {
	payload := BookCreateScheme{}

	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		log.Error(err)
	} else {
		book := models.Book{
			Title:       payload.Title,
			Description: payload.Description,
			AuthorID:    payload.AuthorID,
		}

		db := dbc.GetDB()
		err := db.Create(&book).Error
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			log.Error(err)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"data":   book,
			})
		}
	}
}

func (c BookController) List(ctx *gin.Context) {
	var books []models.Book
	db := dbc.GetDB()
	err := db.Model(&models.Book{}).Preload("Author").Find(&books).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, books)
	}
}
