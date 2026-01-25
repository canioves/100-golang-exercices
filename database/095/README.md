# Exercise 95: MongoDB Aggregation Pipeline

## What is Aggregation?

Aggregation in MongoDB processes data records and returns computed results. It's similar to SQL's GROUP BY operations but more powerful and flexible.

## Aggregation Pipeline

The aggregation pipeline consists of stages that process documents:
- Documents pass through stages sequentially
- Each stage transforms the document stream
- Stages can filter, group, sort, and compute values

```go
pipeline := mongo.Pipeline{
    {{"$match", bson.D{{"status", "A"}}}},
    {{"$group", bson.D{{"_id", "$category"}, {"total", bson.D{{"$sum", "$amount"}}}}}},
}
```

## Common Aggregation Stages

- **$match**: Filter documents (like WHERE clause)
- **$group**: Group documents and perform calculations
- **$sort**: Sort documents
- **$project**: Select/transform fields
- **$limit**: Limit number of results
- **$skip**: Skip documents

## Group Stage Operations

The `$group` stage performs calculations:
- **$sum**: Calculate totals
- **$avg**: Calculate averages  
- **$min/$max**: Find minimum/maximum values
- **$count**: Count documents
- **$push**: Create arrays of values

## BSON Document Structure

Group stage example:
```go
groupStage := bson.D{
    {"$group", bson.D{
        {"_id", "$name"},                              // Group by name
        {"average_age", bson.D{{"$avg", "$age"}}},    // Calculate average age
        {"count", bson.D{{"$sum", 1}}},               // Count occurrences
    }},
}
```

## Your Task

Look at the main.go file and complete the exercise by:
1. Creating a group stage with proper BSON structure
2. Using aggregation operations like $avg and $sum
3. Running aggregation pipeline with Aggregate()
4. Processing aggregation results

This exercise introduces MongoDB's powerful aggregation framework for data analysis and reporting.

```go
// Aggregations
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
	
	// We are going to create a group stage, this will have the result for the average age for each named user!.
	// Create a variable named groupStage 
	// we will assign it a bson.D object with the following information:
	/*
		{"$group", bson.D{
			{"_id", "$name"},
			{"average_price", bson.D{{"$avg", "$age"}}},
			{"numTimes", bson.D{{"$sum", 1}}},
		}},
	*/
	
	
	// use the Aggregate() function with the second argument as mongo.Pipeline{groupStage}:
	cursor, err := 
	if err != nil {
		panic(err)
	}
	
	// display the results
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
	    panic(err)
	}

	for _, result := range results {
	    log.Printf("Average age of users named like %v user options: %v \n", result["_id"], result["average_price"])
	    log.Printf("Number of %v tea options: %v \n\n", result["_id"], result["numTimes"])
	}	
}
```

<details>
<summary> Solution: </summary>

```go
// Aggregations
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
	
	// We are going to create a group stage, this will have the result for the average age of all our users.
	// Create a variable named groupStage with the first element as "average_price" and the second element a bson.D variable containing "$avg" expression and "$age"
	groupStage := bson.D{
						{"$group", bson.D{
							{"_id", "$name"},
							{"average_price", bson.D{{"$avg", "$age"}}},
							{"numTimes", bson.D{{"$sum", 1}}},
						}},
	}
	
	// use the Aggregate() function with the second argument as mongo.Pipeline{groupStage}:
	cursor, err := usersCollection.Aggregate(context.TODO(), mongo.Pipeline{groupStage})
	if err != nil {
		panic(err)
	}
	
	// display the results
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
	    panic(err)
	}

	for _, result := range results {
	    log.Printf("Average age of users named like %v user options: %v \n", result["_id"], result["average_price"])
	    log.Printf("Number of %v tea options: %v \n\n", result["_id"], result["numTimes"])
	}
	
}
```

</details>