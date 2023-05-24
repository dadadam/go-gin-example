package controllers

import (
	"net/http"

	dbc "github.com/dadadam/sono-backend/db"
	"github.com/dadadam/sono-backend/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type AuthorScheme struct {
	ID    uint   `json:"id" binding:"-"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type AuthorController struct{}

func (c AuthorController) Create(ctx *gin.Context) {
	payload := AuthorScheme{}

	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		log.Error(err)
	} else {
		author := models.Author{
			Name:  payload.Name,
			Email: payload.Email,
		}

		db := dbc.GetDB()
		err := db.Create(&author).Error
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			log.Error(err)
		} else {
			ctx.JSON(http.StatusOK, author)
		}
	}
}

func (c AuthorController) List(ctx *gin.Context) {
	var authors []models.Author
	db := dbc.GetDB()
	err := db.Find(&authors).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   authors,
	})
}
