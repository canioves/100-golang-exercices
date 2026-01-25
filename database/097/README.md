# Exercise 97: Fiber Web Framework Basics

## Introduction to Fiber

Fiber is a modern Go web framework inspired by Express.js, designed for building fast and scalable web applications and APIs.

## Key Fiber Features

- **Zero Memory Allocation**: Efficient memory usage
- **Fast HTTP Server**: Built on top of Fasthttp
- **Express-like API**: Familiar routing and middleware patterns
- **Built-in Middleware**: CORS, Logger, Recovery, and more
- **Template Engines**: Support for various template engines

## Basic Fiber Application

```go
app := fiber.New()

app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
})

app.Listen(":3000")
```

## Context Object

The `*fiber.Ctx` context provides:
- **Request Data**: Headers, parameters, body
- **Response Methods**: Send JSON, HTML, files
- **Middleware**: Next() for middleware chaining
- **Utilities**: Parsing, validation, caching

## Route Definition

Fiber supports all HTTP methods:
- **GET**: Retrieve data
- **POST**: Create resources
- **PUT**: Update resources
- **DELETE**: Remove resources
- **PATCH**: Partial updates

## Response Methods

Common response methods:
- **SendString()**: Plain text response
- **JSON()**: JSON response
- **HTML()**: HTML response
- **SendFile()**: File download
- **Redirect()**: HTTP redirects

## Your Task

Look at the main.go file and complete the exercise by:
1. Creating a new Fiber application instance
2. Implementing the helloWorld handler function
3. Setting up basic routes with setupRoutes function
4. Understanding Fiber's application structure

This exercise introduces the Fiber web framework and establishes the foundation for building web APIs.

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
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
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
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	return client.Database("TestCluster")
}

// Create a simple function to return some string to the fiber context
// this function signature will be [func helloWorld(c *fiber.Ctx) error]
// inside you will call the Send() method in that context, with the string "Hello World!" as it's only argument.
func helloWorld(c *fiber.Ctx) error {
	return 
}

// Create a function setUP routes with a signature like func setupRoutes(a *fiber.App)
// inside this function use a.GET to the root path, and the helloWorld function we created above
func setupRoutes(app *fiber.App) {
	
}

// main 
func main (){
	// Fiber 
	// Create a new App with fiber.New()
	app := 
	// Add the routes to the app we created above (right now we only have 1 route)
	
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
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
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
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	return client.Database("TestCluster")
}

// Create a simple function to return some string to the fiber context
// this function signature will be [func helloWorld(c *fiber.Ctx)]
// inside you will call the Send() method in that context, with the string "Hello World!" as it's only argument.
func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// Create a function setUP routes with a signature like func setupRoutes(a *fiber.App)
// inside this function use a.GET to the root path, and the helloWorld function we created above
func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
}

// main 
func main (){
	// Fiber 
	// Create a new App with fiber.New()
	app := fiber.New()
	// Add the routes to the app we created above (right now we only have 1 route)
	setupRoutes(app)
	// Listen to port 3000
    app.Listen(":3000")
}
```

</details>