# Exercise 79: Introduction to Gin Framework

## What is Gin?

Gin is a popular HTTP web framework for Go. It provides a more convenient API than the standard `net/http` package, with features like middleware, JSON binding, and easier routing.

## Why Use Gin?

- **Faster development**: Less boilerplate code
- **Built-in features**: JSON binding, validation, middleware
- **Performance**: Very fast HTTP router
- **Familiar API**: Similar to Express.js (Node.js)

## Installing Gin

```bash
go mod init your-project
go get -u github.com/gin-gonic/gin
```

## Basic Gin Server

```go
import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello World",
        })
    })
    
    r.Run(":8080")
}
```

## Gin vs net/http

**Standard net/http**:
```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
```

**With Gin**:
```go
func handler(c *gin.Context) {
    c.JSON(200, data)
}
```

## Your Task

Create your first Gin web server:
1. Install Gin framework
2. Create a basic server with Gin
3. Set up a simple route that returns JSON
4. Compare the simplicity to standard net/http

This introduces modern Go web development with frameworks.

```go
// Exercise: Create an API with GIN framework
package main

// remember to install gin!!
// go mod init && go get -u github.com/gin-tonic/gin
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
func getAlbums() {
  // use the IndentedJSON function with the gin context, and pass it an http status ok, and the albums array
  // more info: https://pkg.go.dev/github.com/gin-gonic/gin#Context.IndentedJSON
  
}

func main() {  
  // we will registrer a handler (or router) with gin.Default()
  
  
  // here access the router .GET http verb for the request. 
  // https://pkg.go.dev/github.com/gin-gonic/gin#readme-using-get-post-put-patch-delete-and-options
  // the first argument will be the URI (or pattern) and the second one the getAlbums handler function we defined before
  

  // run the server with the Run() function
  
  
}
```

<details>
<summary> Solution: </summary>

```go
// Exercise: Create an API with GIN framework
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

func main() {  
  // we will registrer a handler (or router) with gin.Default()
  router := gin.Default()
  
  // here access the router .GET http verb for the request. 
  // https://pkg.go.dev/github.com/gin-gonic/gin#readme-using-get-post-put-patch-delete-and-options
  // the first argument will be the URI (or pattern) and the second one the getAlbums handler function we defined before
  router.GET("/albums", getAlbums)

  // run the server with the Run() function
  router.Run("localhost:8080")
  
}
```

</details>