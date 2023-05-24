package server

import (
	"github.com/dadadam/sono-backend/controllers"
	"github.com/dadadam/sono-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.ResponseWrapperMiddleware())

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	api := router.Group("/api")

	v1 := api.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/:id", user.Retreive)
		}

		authorGroup := v1.Group("author")
		{
			author := new(controllers.AuthorController)
			authorGroup.POST("", author.Create)
			authorGroup.GET("", author.List)
		}

		bookGroup := v1.Group("book")
		{
			book := new(controllers.BookController)
			bookGroup.POST("", book.Create)
			bookGroup.GET("", book.List)
		}
	}

	return router
}
