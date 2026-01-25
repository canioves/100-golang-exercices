# Exercise 93: Document Queries and Unmarshaling

## Querying Documents

MongoDB uses filters to query documents. In Go, filters are typically created using `bson.D` (ordered) or `bson.M` (unordered) types:

```go
// Find by exact match
filter := bson.D{{"name", "John"}}

// Find by multiple criteria
filter := bson.D{{"name", "John"}, {"age", 25}}
```

## BSON Types in Go

The MongoDB Go driver provides several BSON types:
- **bson.D**: Ordered document (slice of key-value pairs)
- **bson.M**: Unordered document (map)
- **bson.A**: Array (slice)
- **bson.E**: Element (key-value pair)

## FindOne Operation

`FindOne()` retrieves a single document matching the filter:
- Returns first matching document
- Requires `Decode()` to unmarshal into Go struct
- Returns error if no document found

```go
var user User
err := collection.FindOne(context.TODO(), filter).Decode(&user)
```

## Unmarshaling Process

1. **Query**: MongoDB returns BSON document
2. **Decode**: Convert BSON to Go struct
3. **Type Mapping**: BSON types map to Go types
4. **Field Matching**: BSON field names match struct tags

## Error Handling

Common query errors:
- **mongo.ErrNoDocuments**: No document matches filter
- **Decode errors**: BSON to struct conversion issues
- **Connection errors**: Database connectivity problems

## Your Task

Look at the main.go file and complete the exercise by:
1. Creating query filters using bson.D
2. Executing FindOne operations
3. Decoding BSON results into Go structs
4. Handling query results and errors

This exercise demonstrates document retrieval and BSON unmarshaling in MongoDB operations.

```go
// UNMARSHALL
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
	
	// Let's add a second user!
	created_user := User{"Mary", 50}
	res , err := usersCollection.InsertOne(context.TODO(), created_user)
	if err != nil {
		panic(err)
	}
	// We can reference it!
	log.Println("User with id created:", res.InsertedID)
	// And now let's find one user named "John" in the database, let's create a search filter with has the "name","John"
	filter := bson.D{{"name","John"}}
	// Let's create a return variable of type bson.D
	
	// Let's use the FindOne(options).Decode(&result) method and reference the return value in the result variable created above
	err = 
	if err != nil {
		panic(err)
	}
	log.Println(result)
}
```

<details>
<summary> Solution: </summary>

```go
// UNMARSHALL
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
	
	// Let's add a second user!
	created_user := User{"Mary", 50}
	res , err := usersCollection.InsertOne(context.TODO(), created_user)
	if err != nil {
		panic(err)
	}
	// We can reference it!
	log.Println("User with id created:", res.InsertedID)
	// And now let's find one user named "John" in the database, let's create a search filter with has the "name" = "John"
	filter := bson.D{{"name","John"}}
	// Let's create a return variable of type bson.D
	var result bson.D 
	// Let's use the FindOne() method and reference the return value in the result variable created above
	err = usersCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	log.Println(result)
	
}
```

</details>