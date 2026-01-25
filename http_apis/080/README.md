# Exercise 80: Gin POST Requests

## Handling POST Requests

POST requests are used to send data to the server (creating new resources, form submissions, etc.). Gin makes handling POST requests and JSON data very easy.

## POST Route in Gin

```go
r.POST("/albums", func(c *gin.Context) {
    // Handle POST request
})
```

## Binding JSON Data

Gin can automatically bind JSON request bodies to Go structs:

```go
type Album struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Artist string `json:"artist"`
}

func createAlbum(c *gin.Context) {
    var newAlbum Album
    
    if err := c.ShouldBindJSON(&newAlbum); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    // Process the album data
    c.JSON(201, newAlbum)
}
```

## HTTP Methods in Gin

- **GET**: `r.GET("/path", handler)`
- **POST**: `r.POST("/path", handler)`
- **PUT**: `r.PUT("/path", handler)`
- **DELETE**: `r.DELETE("/path", handler)`

## Testing POST Requests

Use curl to test:
```bash
curl -X POST http://localhost:8080/albums \
  -H "Content-Type: application/json" \
  -d '{"id":"1","title":"Blue Album","artist":"Weezer"}'
```

## Your Task

Create a Gin server that:
1. Defines an album struct with JSON tags
2. Handles POST requests to create new albums  
3. Binds JSON request data to the struct
4. Returns appropriate status codes
5. Handles binding errors gracefully

This demonstrates building REST APIs that accept data.

```go
// Exercise: Create an API with GIN framework - POST request
package main

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
  // Create a new album
  var newAlbum album

  // Call BindJSON to bind the received JSON to
  // newAlbum.
  if err := c.BindJSON(&newAlbum); err != nil {
    return
  }

  // Append the newAlbum to the albums slice

  // Indicate the creation status for the new album

}

func main() {  
  // we will registrer a handler (or router) with gin.Default()
  router := gin.Default()
  router.GET("/albums", getAlbums)

  // Add the POST request to the '/albums' path
  

  // run the server with the Run() function
  router.Run("localhost:8080")
}
```

<details>
<summary> Solution: </summary>

```go
// Exercise: Create an API with GIN framework - POST request
package main

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

// Then, let's create a handler function ((a normal function, which then we will wrap :) ))
// name it getAlbums, and it's only argument named 'c' will be a pointer type of gin.Context.
func getAlbums(c *gin.Context) {
  // use the IndentedJSON function with the gin context, and pass it an http status ok, and the albums array
  // more info: https://pkg.go.dev/github.com/gin-gonic/gin#Context.IndentedJSON
  c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums( c *gin.Context){
  // Create a new album
  var newAlbum album

  // Call BindJSON to bind the received JSON to
  // newAlbum.
  if err := c.BindJSON(&newAlbum); err != nil {
    return
  }

  // Append the newAlbum to the albums slice
  albums = append(albums,newAlbum)

  // Indicate the creation status for the new album
  c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {  
  // we will registrer a handler (or router) with gin.Default()
  router := gin.Default()
  
  // here access the router .GET http verb for the request. 
  // https://pkg.go.dev/github.com/gin-gonic/gin#readme-using-get-post-put-patch-delete-and-options
  // the first argument will be the URI (or pattern) and the second one the getAlbums handler function we defined before
  router.GET("/albums", getAlbums)

  // Add the POST request to the '/albums' path
  router.POST("/albums", postAlbums)

  // run the server with the Run() function
  router.Run("localhost:8080")
}
```

</details>