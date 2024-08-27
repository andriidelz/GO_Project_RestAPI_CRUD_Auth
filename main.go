package main

import (
	"net/http"
	"test-vscode-go-module/models"
	"test-vscode-go-module/routes"

	"github.com/gin-gonic/gin"
)

var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbums(c *gin.Context) {
	var newAlbum models.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{"bad_request"})
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, ErrorResponse{"not_found"})
}

func DeleteAlbumById(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, a)
			return
		}
		c.IndentedJSON(http.StatusNotFound, ErrorResponse{"not_found"})
	}
}

func UpdateAlbumById(c *gin.Context) {
	id := c.Param("id")
	for i := range albums {
		if albums[i].ID == id {
			c.BindJSON(&albums[i])
			a := albums[i]
			c.IndentedJSON(http.StatusOK, a)
			return
		}
		c.IndentedJSON(http.StatusNotFound, ErrorResponse{"not_found"})
	}
}

func main() {
	router := gin.Default()
	routes.SetupRouter()
	router.GET("/albums", GetAlbums)
	router.GET("/albums/:id", GetAlbumById)
	router.DELETE("/albums/:id", DeleteAlbumById)
	router.PUT("/albums/:id", UpdateAlbumById)
	router.POST("/albums", PostAlbums)
	router.Run("localhost:8080")
}
