# Exercise 82: Gin Query Parameters

## Query String Parameters

Query parameters are key-value pairs in the URL after the `?` symbol. They're used for filtering, pagination, and optional parameters.

## Query Parameter Examples

```
http://example.com/albums?artist=Beatles&year=1969
http://example.com/search?q=go+programming&limit=10
```

## Accessing Query Parameters in Gin

```go
r.GET("/albums", func(c *gin.Context) {
    artist := c.Query("artist")        // Required parameter
    year := c.DefaultQuery("year", "") // Optional with default
    
    if artist == "" {
        c.JSON(400, gin.H{"error": "artist parameter required"})
        return
    }
    
    // Use parameters to filter results
})
```

## Query Parameter Methods

- **c.Query("key")**: Get parameter value (empty string if not found)
- **c.DefaultQuery("key", "default")**: Get parameter with default value
- **c.GetQuery("key")**: Get parameter and boolean indicating if present

## Common Use Cases

- **Filtering**: `?category=electronics&price=100`
- **Pagination**: `?page=2&limit=20`
- **Search**: `?q=search+term`
- **Sorting**: `?sort=name&order=asc`

## Your Task

Create an endpoint that handles query parameters:
1. Accept query parameters (like artist name, year, etc.)
2. Use parameters to filter or search albums
3. Provide default values for optional parameters
4. Return filtered results based on query parameters
5. Handle missing required parameters gracefully

This demonstrates building flexible APIs that accept optional filtering criteria.

```go
// Exercise: Create an API with GIN framework - Query string parameters
package main

// In this exercise what we want to do is handle query string parameters
// Query string parameters are the ones you may find in the url like this:
// http://example.com/welcome?name=enin&surname=tolstoy
//                            <---1--->&<---2---------> .  2 parameters in this case


// We will handle them with gin, and we will try to access an album the same way in the previous exercise, but with it's query string parameters
// We will modify the first function to handle both ways (with and without query string parameters)
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

// IN THIS EXERCISE YOU HAVE TO MODIFY THIS FUNCTION!!
func getAlbums(c *gin.Context) {
  // get the query string (look up for the .Query() function in gin :) )
  // our QueryString key will be "id" and the value, the ID we want to get
  
  // if the id is null, serve all the albums (like in the first gin API exercise)
  

  // else, loop over the array to serve the matching album.
    
  // return some error message when album isn't found inside this else block!!
  
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
  var id string = c.Param("id")
  for _, album := range albums {
		if id == album.ID {
			c.IndentedJSON(http.StatusOK, album)
      return
		}
	}
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}


func main() {  
  // we will registrer a handler (or router) with gin.Default()
  router := gin.Default()

  router.GET("/albums", getAlbums)
  router.POST("/albums", postAlbums)
  router.GET("/albums/:id", getSpecificAlbum)

  // run the server with the Run() function
  router.Run("localhost:8080")
}
```

<details>
<summary> Solution: </summary>

```go
// Exercise: Create an API with GIN framework - Query string parameters
package main

// In this exercise what we want to do is handle query string parameters
// Query string parameters are the ones you may find in the url like this:
// http://example.com/welcome?name=enin&surname=tolstoy
//                            <---1--->&<---2---------> .  2 parameters in this case


// We will handle them with gin, and we will try to access an album the same way in the previous exercise, but with it's query string parameters
// We will modify the first function to handle both ways (with and without query string parameters)
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

// IN THIS EXERCISE YOU HAVE TO MODIFY THIS FUNCTION!!
func getAlbums(c *gin.Context) {

  // get the query string (look up for the .Query() function in gin :) )
  // our QueryString key will be "id" and the value, the ID we want to get
  var id string = c.Query("id")
  if (id == ""){
    c.IndentedJSON(http.StatusOK, albums)
  } else {
    // Here, loop over the albums array again to see if we have that query!
    for _, album := range albums {
	  	if id == album.ID {
	  		c.IndentedJSON(http.StatusOK, album)
        return
	  	}
	  }
    // return some error message when album isn't found!
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
  }
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

  // modify this first function!!
  router.GET("/albums", getAlbums)
  router.POST("/albums", postAlbums)
  router.GET("/albums/:id", getSpecificAlbum)

  // run the server with the Run() function
  router.Run("localhost:8080")
}
```

</details>