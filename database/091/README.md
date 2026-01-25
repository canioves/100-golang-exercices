# Exercise 91: MongoDB Collections

## Understanding MongoDB Collections

Collections in MongoDB are equivalent to tables in relational databases. They group related documents together and can be created explicitly or automatically when first accessed.

## Collection Operations

MongoDB provides several collection management operations:
- **CreateCollection()**: Explicitly create collections with options
- **Collection()**: Get reference to existing or new collection
- **ListCollections()**: View all collections in database
- **Drop()**: Remove collections

## Capped Collections

Capped collections have special characteristics:
- **Fixed Size**: Maximum storage size defined at creation
- **FIFO Behavior**: Oldest documents removed when size limit reached
- **High Performance**: Optimized for high-throughput operations
- **No Updates**: Documents cannot be updated to larger sizes

```go
opts := options.CreateCollection()
opts.SetCapped(true)
opts.SetSizeInBytes(1048576) // 1MB
```

## Collection Options

Common collection creation options:
- **Capped**: Create fixed-size collection
- **Size**: Maximum size in bytes
- **Max**: Maximum number of documents
- **Validator**: Document validation rules

## Use Cases for Capped Collections

- **Logging**: High-volume log data with automatic rotation
- **Metrics**: Time-series data with size constraints
- **Cache**: Fixed-size cache with automatic eviction
- **Event Streams**: Real-time event processing

## Your Task

Look at the main.go file and complete the exercise by:
1. Setting up collection creation options
2. Creating a capped collection with size limits
3. Getting collection references for operations
4. Understanding collection configuration

This exercise introduces you to MongoDB collection management and specialized collection types.

```go
package main

import (
	"context"
	"log"
	"os"
	"github.com/joho/godotenv"
//	"go.mongodb.org/mongo-driver/bson"
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
	
	// Create a collection called "users" into your Database.
	// To do so, use the client.Database() call, and call the .CreateCollection() method to create a collection if it doesn't exist!
	// The third and subsequent arguemnts (4th,5th,6th... will be options) 
	// We will set the options SetCapped to true and SetSizeInBytes to 1048576.
	// for better readness, we will declare a opts object as follows:
	opts := options.CreateCollection()
	// And here create the database collection below this line with the context.TODO() as first argument, the collection name as the second, and the 2 options as the third and fourth

	
	// Here, we will return the name of the collection, to check everything went allright :)
	usersCollection := client.Database("TestCluster").Collection("users")
	log.Println(usersCollection.Name())
	log.Println("You got connected!")
	
}
```

<details>
<summary> Solution: </summary>

```go
package main

import (
	"context"
	"log"
	"os"
	"github.com/joho/godotenv"
//	"go.mongodb.org/mongo-driver/bson"
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
	
	// Create a collection called "users" into your Database.
	// To do so, use the client.Database() call, and call the .CreateCollection() method to create a collection if it doesn't exist!
	// The third and subsequent arguemnts (4th,5th,6th... will be options) 
	// We will set the options SetCapped to true and SetSizeInBytes to 1048576.
	// for better readness, we will declare a opts object as follows:
	opts := options.CreateCollection()
	// And here create the database collection below this line with the context.TODO() as first argument, the collection name as the second, and the 2 options as the third and fourth
	client.Database("TestCluster").CreateCollection(context.TODO(), "users", opts.SetCapped(true), opts.SetSizeInBytes(1048576))
	
	// Here, we will return the name of the collection, to check everything went allright :)
	usersCollection := client.Database("TestCluster").Collection("users")
	log.Println(usersCollection.Name())
	log.Println("You got connected!")
	
}
```

</details>