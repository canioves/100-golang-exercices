# Exercise 78: Creating a JSON API

## Building REST APIs

APIs (Application Programming Interfaces) allow different systems to communicate. JSON APIs are the most common format for modern web services.

## JSON API Basics

A JSON API typically:
1. **Accepts** JSON requests
2. **Processes** data  
3. **Returns** JSON responses
4. **Uses** HTTP methods (GET, POST, PUT, DELETE)

## Setting Response Headers

For JSON responses, set the content type:
```go
w.Header().Set("Content-Type", "application/json")
```

## Returning JSON Data

```go
type Response struct {
    Message string `json:"message"`
    Data    Animal `json:"data"`
}

response := Response{Message: "success", Data: animal}
json.NewEncoder(w).Encode(response)
```

## HTTP Status Codes

- **200**: OK
- **201**: Created  
- **400**: Bad Request
- **404**: Not Found
- **500**: Internal Server Error

## Your Task

Create a JSON API with an Animal struct:
1. Define an Animal struct with Name and Type fields
2. Create an endpoint that returns JSON data
3. Set proper content type headers
4. Return structured JSON responses

This is the foundation for building REST APIs that other applications can consume.

```go
// Exercise: Create an API
package main

import (
  "encoding/json"
  "net/http"
  "log"
)
// Let's create a struct 'Animal' with a variable 'Name' and another called 'Type' both type string
type Animal struct {
  Name  string `json:"Name"`
  Type  string `json:"Type"`
}

// Then, let's create a handler function ((a normal function, which then we will wrap :) ))
func AnimalsHandler(w http.ResponseWriter, r *http.Request) {
  // Create an array of 3 animals called "animals" and put 3 animals in there with it's name and type:
  
  // use the json NewEncoder() function, to encode the animals array within the responsewriter element
  
}

func main() {  
  // Register the handler into the defaultservemux (remember which function was used!) at the root path "/"
  
  // Start the server with the DefaultServeMux!
  server := http.ListenAndServe(":8080", )
  if (server != nil){
    log.Print("Cannot start sever")
  }
}
```

<details>
<summary> Solution: </summary>

```go
// Exercise: Create an API
package main

import (
  "encoding/json"
  "net/http"
  "log"
)
// Let's create a struct 'Animal' with a variable 'Name' and another called 'Type' both type string
type Animal struct {
  Name  string `json:"Name"`
  Type  string `json:"Type"`
}

// Then, let's create a handler function ((a normal function, which then we will wrap :) ))
func AnimalsHandler(w http.ResponseWriter, r *http.Request) {
  // Create an array of 3 animals called "animals" and put 3 animals in there with it's name and type:
  animals := []Animal{
      Animal{"Alice","Cat"},
      Animal{"Bob","Cat"},
      Animal{"Trinity","Dog"},
  }
  // use the json NewEncoder() function, to encode the animals array within the responsewriter element
  json.NewEncoder(w).Encode(animals)
}

func main() {  
  // Register the handler into the defaultservemux (remember which function was used!) at the root path "/"
  http.HandleFunc("/", AnimalsHandler)
  // Start the server with the DefaultServeMux!
  server := http.ListenAndServe(":8080", nil)
  if (server != nil){
    log.Print("Cannot start sever")
  }
}
```

</details>