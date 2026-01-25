# Exercise 92: BSON and Document Insertion

## What is BSON?

BSON (Binary JSON) is MongoDB's binary serialization format for storing documents. It extends JSON with additional data types like ObjectId, Date, and Binary data while maintaining JSON-like simplicity.

## Go Structs and BSON

Go structs map naturally to BSON documents:
- Struct fields become document fields
- Field tags control serialization
- Embedded structs create nested documents

```go
type User struct {
    Name string `bson:"name"`
    Age  int    `bson:"age"`
}
```

## Document Insertion

MongoDB provides several insertion methods:
- **InsertOne()**: Insert single document
- **InsertMany()**: Insert multiple documents
- **ReplaceOne()**: Replace entire document

```go
result, err := collection.InsertOne(context.TODO(), user)
insertedID := result.InsertedID
```

## BSON Tags

Common BSON struct tags:
- `bson:"fieldname"`: Map to specific field name
- `bson:",omitempty"`: Omit empty values
- `bson:"-"`: Exclude field from serialization
- `bson:",inline"`: Inline embedded struct fields

## Context Usage

MongoDB operations require context for:
- **Cancellation**: Cancel long-running operations
- **Timeouts**: Set operation deadlines
- **Metadata**: Pass request-scoped values

## Your Task

Look at the main.go file and complete the exercise by:
1. Defining User struct with proper BSON tags
2. Creating user instances with data
3. Inserting documents into MongoDB collection
4. Understanding BSON serialization in Go

This exercise teaches you document structure and basic insertion operations in MongoDB.

```go
// In this exercises, we are going to learn about BSON, a binary serialization format (like JSON) which is used to marshall and unmarshall data ad make remote calls in mongoDB
// First, we need to import the `go.mongodb.org/mongo-driver/bson` library.
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

// Here we are going to create a defined data structure called "User"
// It will have two values:
// - Name, a string
// - Age, a integer
type User struct {
	
}

// main func
func main (){
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
	log.Println("You got connected!")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	usersCollection := client.Database("TestCluster").Collection("users")
	log.Println(usersCollection.Name())

	//Here, we are going to create a User.
	name :=  
	age  := 
	created_user := 
	// Now we will use the InsertOne() function to add the created_user into the users collection.
	_, err = 

	if err != nil {
		panic(err)
	}
	
}
```

<details>
<summary> Solution: </summary>

```go
// In this exercises, we are going to learn about BSON, a binary serialization format (like JSON) which is used to marshall and unmarshall data ad make remote calls in mongoDB
// First, we need to import the `go.mongodb.org/mongo-driver/bson` library.
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

// Here we are going to create a defined data structure called "User"
// It will have two values:
// - Name, a string
// - Age, a integer
type User struct{
	Name string
	Age  int
}

func main (){
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
	usersCollection := client.Database("TestCluster").Collection("users")
	log.Println(usersCollection.Name())

	//Here, we are going to create a User.
	name := "John" 
	age  := 24
	created_user := User{name, age}
	// Now we will use the InsertOne() function to add the created_user into the users collection.
	_, err = usersCollection.InsertOne(context.TODO(), created_user)
	if err != nil {
		panic(err)
	}

	log.Println("You got connected!")
	
}
```

</details>