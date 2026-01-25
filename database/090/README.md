# Exercise 90: MongoDB Connection Setup

## Introduction to MongoDB

MongoDB is a popular NoSQL document database that stores data in flexible, JSON-like documents. It's designed for scalability, performance, and ease of development with modern applications.

## MongoDB Atlas

MongoDB Atlas is MongoDB's cloud database service offering:
- **Free Tier**: Perfect for learning and development
- **Managed Service**: No infrastructure management needed
- **Global Clusters**: Deploy worldwide
- **Built-in Security**: Authentication and encryption

## Go MongoDB Driver

The official MongoDB Go driver provides:
- Native Go API for MongoDB operations
- Connection pooling and management
- BSON encoding/decoding
- Comprehensive error handling

```go
import "go.mongodb.org/mongo-driver/mongo"
import "go.mongodb.org/mongo-driver/mongo/options"
```

## Environment Variables

Using `.env` files for configuration:
- Keeps sensitive data out of source code
- Different environments (dev, staging, prod)
- Easy configuration management

```go
import "github.com/joho/godotenv"

err := godotenv.Load()
uri := os.Getenv("MONGODB_URI")
```

## Connection Lifecycle

1. **Load Environment**: Read connection string from .env
2. **Create Client**: Initialize MongoDB client with URI
3. **Connect**: Establish connection to database
4. **Use Database**: Perform operations
5. **Disconnect**: Clean up resources

## Your Task

Look at the main.go file and complete the exercise by:
1. Installing required packages (mongo-driver and godotenv)
2. Loading environment variables from .env file
3. Establishing connection to MongoDB Atlas
4. Implementing proper connection cleanup

This exercise establishes the foundation for database operations in Go applications.

```go
package main

// In this exercise we will Stablish a connection with mongoDB - Atlas
// MongoDB is a non relational kind of db, and it's saas has a free tier. Find out more here.
// https://www.mongodb.com/docs/atlas/tutorial/deploy-free-tier-cluster/
// You can also use a docker container locally to connect, use whatever option suits you best.

// 1- You need to install these two libraries through the following commands, one is for the mongo driver and the other for a dotenv file
// go get go.mongodb.org/mongo-driver/mongo
// go get github.com/joho/godotenv
// 2- Once installed, create a file named ".env" containing your connection string. Use that env file to connect to the db.
//    This .env file, will have one variable called "MONGODB_URI=" with the value of your mongodb connection string.
import (
	"context"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main (){
	// Load the godotenv library
	err := 

	if err != nil {
		log.Fatal("Could not load .env file")
	}

	// Create a variable named "uri" containing your MONGODB_URI string connection.
	uri := os.Getenv()
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	// Thanks to the dotenv file, you can get the MONGODB_URI from the environment!
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	log.Println("You got connected!")
}
```

<details>
<summary> Solution: </summary>

```go
package main

// In this exercise we will Stablish a connection with mongoDB - Atlas
// MongoDB is a non relational kind of db, and it's saas has a free tier. Find out more here.
// https://www.mongodb.com/docs/atlas/tutorial/deploy-free-tier-cluster/
// You can also use a docker container locally to connect, use whatever option suits you best.

// 1- You need to install these two libraries through the following commands, one is for the mongo driver and the other for a dotenv file
// go get go.mongodb.org/mongo-driver/mongo
// go get github.com/joho/godotenv
// 2- Once installed, create a file named ".env" containing your 
import (
	"context"
	//"encoding/json"
	//"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main (){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}

	// Create a variable named "uri" containing your MONGODB_URI string connection.
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

	log.Println("You got connected!")
}
```

</details>