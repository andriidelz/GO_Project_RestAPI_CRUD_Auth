package routes

import (
	"test-vscode-go-module/controllers"
	"test-vscode-go-module/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/login", controllers.Login)
	// router.GET("/albums", controllers.GetAlbums)

	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/albums", controllers.GetAlbums)
		authorized.POST("/albums", controllers.PostAlbums)
	}

	return router
}
