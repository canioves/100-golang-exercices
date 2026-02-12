// Exercise: Create an API with GIN framework - POST request
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	// Create a new album
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Append the newAlbum to the albums slice
	albums = append(albums, newAlbum)
	// Indicate the creation status for the new album
	c.JSON(201, newAlbum)
}

func main() {
	// we will registrer a handler (or router) with gin.Default()
	router := gin.Default()
	router.GET("/albums", getAlbums)

	// Add the POST request to the '/albums' path
	router.POST("/albums", postAlbums)
	// run the server with the Run() function
	router.Run("localhost:8080")
}
