# Exercise 81: Gin Path Parameters

## URL Path Parameters

Path parameters allow you to capture parts of the URL as variables. This is essential for REST APIs where you need to identify specific resources.

## Path Parameter Syntax

```go
// Route with parameter
r.GET("/albums/:id", getAlbumByID)

// In handler function
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")  // Get the :id parameter
    // Use id to find specific album
}
```

## REST API Patterns

Common RESTful URL patterns:
- **GET /albums**: Get all albums
- **GET /albums/:id**: Get specific album
- **POST /albums**: Create new album
- **PUT /albums/:id**: Update specific album
- **DELETE /albums/:id**: Delete specific album

## Finding Resources

Typically you'd search a database or slice:
```go
func findAlbumByID(id string) (*Album, bool) {
    for _, album := range albums {
        if album.ID == id {
            return &album, true
        }
    }
    return nil, false
}
```

## HTTP Status Codes

- **200**: Found and returning resource
- **404**: Resource not found

## Your Task

Create a GET endpoint that:
1. Accepts an ID as a path parameter (`/albums/:id`)
2. Searches for an album with that ID
3. Returns the album if found (200 status)
4. Returns 404 error if not found
5. Uses proper REST conventions

This demonstrates building resource-specific API endpoints.

```go
// Exercise: Create an API with GIN framework - GET request to specific ITEM
package main

// In this exercise what we want to do is get an specific album by it's ID
// The path we will want to go is "/album/$ID" and it will return our album with the ID (if it exists!!!)

import (
  "github.com/gin-gonic/gin"
  "net/http"
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

func postAlbums( c *gin.Context){
  var newAlbum album
  if err := c.BindJSON(&newAlbum); err != nil {
    return
  }
  albums = append(albums,newAlbum)
  c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getSpecificAlbum(c *gin.Context){
  // get the ID in a variable

  // check if it matches any of the slice ID (you have to iterate through the array)

  // return some error message when album isn't found!

}

func main() {  
  // we will registrer a handler (or router) with gin.Default()
  router := gin.Default()
  router.GET("/albums", getAlbums)
  router.POST("/albums", postAlbums)
  // create the route specific to get the /album/$ID (https://github.com/gin-gonic/gin#parameters-in-path)
  
  // run the server with the Run() function
  router.Run("localhost:8080")
}
```

<details>
<summary> Solution: </summary>

```go
// Exercise: Create an API with GIN framework - GET request to specific ITEM
package main

// In this exercise what we want to do is get an specific album by it's ID
// The path we will want to go is "/album/$ID" and it will return our album with the ID (if it exists!!!)

import (
  "github.com/gin-gonic/gin"
  "net/http"
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

func postAlbums( c *gin.Context){
  var newAlbum album
  if err := c.BindJSON(&newAlbum); err != nil {
    return
  }
  albums = append(albums,newAlbum)
  c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getSpecificAlbum(c *gin.Context){
  // get the ID in a variable
  var id string = c.Param("id")

  // check if it matches any of the slice ID (you have to iterate through the array)
  for _, album := range albums {
		if id == album.ID {
			c.IndentedJSON(http.StatusOK, album)
      return
		}
	}

  // return some error message when album isn't found!
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {  
  // we will registrer a handler (or router) with gin.Default()
  router := gin.Default()
  router.GET("/albums", getAlbums)
  router.POST("/albums", postAlbums)
  // create the route specific to get the /album/$ID (https://github.com/gin-gonic/gin#parameters-in-path)
  router.GET("/albums/:id", getSpecificAlbum)
  // run the server with the Run() function
  router.Run("localhost:8080")
}
```

</details>