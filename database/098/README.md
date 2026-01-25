# Exercise 98: API Data Retrieval with Parameters

## URL Parameters in Fiber

Route parameters allow dynamic URLs where parts of the path are variables:
- **Definition**: `/user/:name` defines a parameter named "name"
- **Access**: `c.Params("name")` retrieves parameter value
- **Flexible**: Multiple parameters per route supported

```go
app.Get("/user/:id/posts/:postid", handler)
// Access: c.Params("id"), c.Params("postid")
```

## Database Queries with Parameters

Using route parameters in database queries:
1. Extract parameter from URL
2. Create database filter using parameter
3. Execute query with filter
4. Return results as JSON

## Error Handling Patterns

Proper API error handling:
- **404 Not Found**: Resource doesn't exist
- **500 Internal Server Error**: Database/system errors
- **400 Bad Request**: Invalid parameters
- **Consistent Format**: Standard error response structure

## JSON Response Structure

Structured API responses provide consistency:
```go
type UserResponse struct {
    Status  int        `json:"status"`
    Message string     `json:"message"`  
    Data    *fiber.Map `json:"data"`
}
```

## Case Sensitivity in Databases

MongoDB queries are case-sensitive by default:
- "John" ≠ "john" 
- Use exact matching for string fields
- Consider case-insensitive queries for search features
- Consistent data input validation important

## Your Task

Look at the main.go file and complete the exercise by:
1. Extracting name parameter from URL
2. Creating BSON filter for database query
3. Executing FindOne operation and decoding result
4. Implementing proper error handling and responses

This exercise demonstrates parameterized API routes and database integration with proper error handling.

```go
// REST API - CRUD operations
// In this exercises, we are going to create an API with the gofiber library.
// https://github.com/gofiber/fiber
// You need to go to the root of this repository and run `go get -u github.com/gofiber/fiber/v2`
package main

import (
	"context"
	"log"
	"os"
	"encoding/json"
	"net/http"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gofiber/fiber/v2"
)

// IMPORTANT FOR THIS EXERCISE, READ ALL THE FILE BEFORE BEGINNING!!! 
// Our goal is to create a new endpoint (app.Get(/users/:id)), so when we try to find /users/John we find one, but when we try to find /users/john it doesn't
// (case sensitivity for this case :) )

type User struct {
	Age  int    `bson:"age"`
	Name string `bson:"name"`
}

// Create a UserResponse type, it will have 3 elements:
// 1- Status, integer
// 2- Message, string
// 3- Data, a *fiber.Map element containing our actual payload. *fiber.Map is a shortcut for map[string]interface{}, useful for JSON returns.
type UserResponse struct {
    Status  int        `json:"status"`
    Message string     `json:"message"`
    Data    *fiber.Map `json:"data"`
}

func initDB() *mongo.Database {
	// MongoDB connection setup
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client.Database("TestCluster")
}

func getUser(c *fiber.Ctx) error {
	// Connect to the DB
	collection := initDB().Collection("users")

	// Get user by name, first extract the name from the URI:
	name :=
	
	// And now let's find one user named "John" in the database, let's create a search filter with has the "name" = parameter we set before
	filter := 
	// Let's create a user variable of type User to reference it with the FindOne function later
	var 
	// Let's use the FindOne() function and reference the return value in the result variable created above!
	err := 
	// Catch the error, if the above query was unsuccesful, return a HTTP500 status message with the err.Error() as the data payload, and "error" as message.
	if err != nil {
        
    }
	
	/*
	// DEBUGGING
	output, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		panic(err)
	}
	log.Printf("%s\n", output)
	*/

	// if everything went alright, return the c.Status(200).JSON(UserResponse{}) with the userresponse status as 200, message as success and the data our user!
	
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	// We will allocate the /user/:name uri, and we will pass the name as path parameter (/user/john)
	
}

// main 
func main (){
	app := fiber.New()
	// Add the routes to the app we created above (right now we only have 2 routes)
	setupRoutes(app)
	// Listen to port 3000
    app.Listen(":3000")
}
```

<details>
<summary> Solution: </summary>

```go
// REST API - CRUD operations
// In this exercises, we are going to create an API with the gofiber library.
// https://github.com/gofiber/fiber
// You need to go to the root of this repository and run `go get -u github.com/gofiber/fiber/v2`
package main

import (
	"context"
	"log"
	"os"
	"encoding/json"
	"net/http"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gofiber/fiber/v2"
)

// IMPORTANT FOR THIS EXERCISE, READ ALL THE FILE BEFORE BEGINNING!!! 
// Our goal is to create a new endpoint (app.Get(/users/:id)), so when we try to find /users/John we find one, but when we try to find /users/john it doesn't
// (case sensitivity for this case :) )

type User struct {
	Age  int    `bson:"age"`
	Name string `bson:"name"`
}

// Create a UserResponse type, it will have 3 elements:
// 1- Status, integer
// 2- Message, string
// 3- Data, a *fiber.Map element containing our actual payload. *fiber.Map is a shortcut for map[string]interface{}, useful for JSON returns.
type UserResponse struct {
    Status  int        `json:"status"`
    Message string     `json:"message"`
    Data    *fiber.Map `json:"data"`
}

func initDB() *mongo.Database {
	// MongoDB connection setup
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client.Database("TestCluster")
}

func getUser(c *fiber.Ctx) error {
	// Connect to the DB
	collection := initDB().Collection("users")

	// Get user by name, first extract the name from the URI:
	name := c.Params("name")
	
	// And now let's find one user named "John" in the database, let's create a search filter with has the "name" = parameter we set before
	filter := bson.D{{"name",name}}
	// Let's create a user variable of type User to reference it with the FindOne function later
	var user User
	// Let's use the FindOne() function and reference the return value in the result variable created above!
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	// Catch the error, if the above query was unsuccesful, return a HTTP500 status message with the err.Error() as the data payload, and "error" as message.
	if err != nil {
        return c.Status(500).JSON(UserResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
	
	/*
	// DEBUGGING
	output, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		panic(err)
	}
	log.Printf("%s\n", output)
	*/

	// if everything went alright, return the c.Status(200).JSON(UserResponse{}) with the userresponse status as 200, message as success and the data our user!
	return c.Status(http.StatusOK).JSON(UserResponse{Status: 200, Message: "success", Data: &fiber.Map{"data": user}})
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	// We will allocate the /user/:name uri, and we will pass the name as queryParam (/user/john)
	app.Get("/user/:name", getUser)
}

// main 
func main (){
	app := fiber.New()
	// Add the routes to the app we created above (right now we only have 1 route)
	setupRoutes(app)
	// Listen to port 3000
    app.Listen(":3000")
}
```

</details>