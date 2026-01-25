# Exercise 99: API Data Creation with POST

## HTTP POST Method

POST requests create new resources:
- **Request Body**: Data sent in HTTP body (usually JSON)
- **Content-Type**: `application/json` for JSON data
- **Status Codes**: 201 Created for successful creation
- **Response**: Return created resource with ID

## Request Body Parsing

Fiber's `BodyParser()` method:
- Automatically parses JSON to Go structs
- Handles content-type detection
- Validates JSON structure
- Returns parsing errors

```go
var user User
err := c.BodyParser(&user)
```

## JSON Struct Tags

Proper JSON mapping with struct tags:
```go
type User struct {
    Age  int    `json:"age"`
    Name string `json:"name"`
}
```

## Database Insertion Flow

1. **Parse Request**: Extract JSON from request body
2. **Validate Data**: Check required fields and formats
3. **Insert Document**: Add to database collection
4. **Return Response**: Confirm creation with inserted ID

## Error Scenarios

Handle common POST errors:
- **400 Bad Request**: Invalid JSON or missing fields
- **409 Conflict**: Duplicate resource
- **500 Internal Server Error**: Database errors
- **503 Service Unavailable**: System issues

## Testing POST Endpoints

Tools for testing POST APIs:
- **Postman**: GUI tool for API testing
- **curl**: Command-line HTTP client
- **HTTPie**: User-friendly HTTP client
- **Automated Tests**: Go testing with HTTP requests

## Your Task

Look at the main.go file and complete the exercise by:
1. Connecting to database collection
2. Parsing JSON request body into User struct
3. Creating new user document with parsed data
4. Inserting document and returning success response

This exercise completes the CREATE operation in a REST API, handling JSON input and database insertion.

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
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Age  int    `json:"age,string"`
	Name string `json:"name"`
}

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

func createUser(c *fiber.Ctx) error {
	// In order to create a user, we are going to send the user attributes through 
	// an http body payload containing the json for the user like  "{"name":"Kyle","age":"23"}"

	// 1. Connect to the DB
	collection := 
	// 2. Create a variable named user and get the user from the request body! (You can use POSTMAN to test! )
	// use the context.BodyParser function referencing that user variable as it's argument
	
	if err != nil {
		return c.Status(503).JSON(UserResponse{Status: 503, Message: "error: Incorrect user payload", Data: &fiber.Map{"data": err.Error()}})
	}

	// 3.Create a variable newUser with the parsed data!
	
	// 4.Insert the user into the MongoDB collection
	// Use the collection.InsertOne() func with the newUser var
	

	if err != nil {
		log.Fatal(err)
	}
	// 5. If everything went fine, return http 200 with the response for the created newUser
	
}

func getUser(c *fiber.Ctx) error {
	// Connect to the DB
	collection := initDB().Collection("users")
	name := c.Params("name")
	filter := bson.D{{"name",name}}
	var user User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
        return c.Status(500).JSON(UserResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

	return c.Status(200).JSON(UserResponse{Status: 200, Message: "success", Data: &fiber.Map{"data": user}})
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/user/:name", getUser)
	// Add the .Post handler to the "/user" endpoint!
	app.Post("/user", createUser)
}

// main 
func main (){
	// Fiber 
	app := fiber.New()
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
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Age  int    `json:"age,string"`
	Name string `json:"name"`
}

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

func createUser(c *fiber.Ctx) error {
	// Connect to the DB
	collection := initDB().Collection("users")
	// Get the user from the request body! (You can use POSTMAN to test! )
	var user User
	err := c.BodyParser(&user);
	if err != nil {
		return c.Status(503).JSON(UserResponse{Status: 503, Message: "error: Incorrect user payload", Data: &fiber.Map{"data": err.Error()}})
	}

	// Create a variable newUser with the parsed data!
	newUser := User{
		Name: user.Name,
		Age : user.Age,
	}
	// Insert the user into the MongoDB collection
	res, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	
	return c.Status(200).JSON(UserResponse{Status: 200, Message: "success", Data: &fiber.Map{"data": res}})
}

func getUser(c *fiber.Ctx) error {
	// Connect to the DB
	collection := initDB().Collection("users")
	name := c.Params("name")
	filter := bson.D{{"name",name}}
	var user User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
        return c.Status(500).JSON(UserResponse{Status: 500, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

	return c.Status(200).JSON(UserResponse{Status: 200, Message: "success", Data: &fiber.Map{"data": user}})
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/user/:name", getUser)
	app.Post("/user", createUser)
}

// main 
func main (){
	// Fiber 
	app := fiber.New()
	setupRoutes(app)
	// Listen to port 3000
    app.Listen(":3000")
}
```

</details>