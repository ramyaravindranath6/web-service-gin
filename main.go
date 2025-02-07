package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(context *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)

}

func getAlbumById(context *gin.Context) {
	var id = context.Param("id")

	// Add the new album to the slice.
	for _, value := range albums {
		if value.ID == id {
			context.IndentedJSON(http.StatusOK, value)
			return
		}

	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
