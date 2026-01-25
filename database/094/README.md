# Exercise 94: Bulk Insert Operations

## Bulk Operations

Bulk operations allow inserting multiple documents in a single database round-trip, providing better performance than individual insertions.

## InsertMany()

`InsertMany()` accepts a slice of documents:
- More efficient than multiple `InsertOne()` calls
- Atomic operation (all succeed or all fail)
- Returns slice of inserted IDs
- Supports ordered and unordered insertions

```go
docs := []interface{}{
    User{Name: "Alice", Age: 30},
    User{Name: "Bob", Age: 25},
}
result, err := collection.InsertMany(context.TODO(), docs)
```

## Interface{} Usage

MongoDB operations often use `interface{}` for flexibility:
- Accepts any document type
- Allows mixed document structures
- Requires type assertion for specific operations
- Common pattern in database operations

## Bulk Operation Benefits

- **Performance**: Fewer network round-trips
- **Atomicity**: All operations succeed or fail together
- **Efficiency**: Reduced connection overhead
- **Throughput**: Higher data ingestion rates

## Ordered vs Unordered

- **Ordered**: Stops on first error, maintains insertion order
- **Unordered**: Continues on errors, may reorder for performance

```go
opts := options.InsertMany().SetOrdered(false)
result, err := collection.InsertMany(ctx, docs, opts)
```

## Use Cases

- **Data Migration**: Moving large datasets
- **Batch Processing**: Processing multiple records
- **Import Operations**: Loading external data
- **Seeding**: Populating test or initial data

## Your Task

Look at the main.go file and complete the exercise by:
1. Creating a slice of user documents with interface{} type
2. Including a user named "Rose" as specified
3. Using InsertMany() for bulk insertion
4. Understanding bulk operation results

This exercise teaches efficient bulk data insertion patterns in MongoDB.

```go
// Bulk insert
// In this exercises, we are going to learn about BSON, a binary serialization format (like JSON) which is used to marshall and unmarshall data ad make remote calls in mongoDB
// First, we need to import the `go.mongodb.org/mongo-driver/bson` library.
package main

import (
	"context"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	
	// Let's add a bunch of users with the InsertMany() method
	// Create a userList of minimum 5 users, and one of them must be called "Rose"
	userList := []interface{}{
		
	}
	res, err := 
	if err != nil {
		panic(err)
	}
	// And now let's find one user named "Rose" in the database, let's create a search filter with has the "name" = "Rose"
	filter := bson.D{{"name","Rose"}}
	// Let's create a return variable of type bson.D
	var result bson.D 
	// Let's use the FindOne() method and reference the return value in the result variable created above
	err = usersCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}

	log.Printf("inserted documents with IDs %v\n", res.InsertedIDs)
	
}
```

<details>
<summary> Solution: </summary>

```go
// Bulk insert
// In this exercises, we are going to learn about BSON, a binary serialization format (like JSON) which is used to marshall and unmarshall data ad make remote calls in mongoDB
// First, we need to import the `go.mongodb.org/mongo-driver/bson` library.
package main

import (
	"context"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	
	// Let's add a bunch of users with the InsertMany() method
	// Create a userList of minimum 5 users, and one of them must be called "Rose"
	userList := []interface{}{
		User{ Name : "Edward", Age: 10},
		User{ Name : "Elise", Age: 20},
		User{ Name : "Caroline", Age: 25},
		User{ Name : "Rose", Age: 30},
		User{ Name : "Staicy", Age: 50},
	}
	res, err := usersCollection.InsertMany(context.TODO(), userList)
	if err != nil {
		panic(err)
	}
	// And now let's find one user named "John" in the database, let's create a search filter with has the "name" = "John"
	filter := bson.D{{"name","Rose"}}
	// Let's create a return variable of type bson.D
	var result bson.D 
	// Let's use the FindOne() method and reference the return value in the result variable created above
	err = usersCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	log.Printf("inserted documents with IDs %v\n", res.InsertedIDs)
	
}
```

</details>