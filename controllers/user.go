package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) Retreive(c *gin.Context) {
	userData := fmt.Sprintf("User ID: %s", c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   userData,
	})
}
