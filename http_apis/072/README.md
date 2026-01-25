# Exercise 72: Basic HTTP Server

## Introduction to net/http

Go's `net/http` package makes it incredibly easy to create web servers. You can build a complete HTTP server with just a few lines of code.

## Simple HTTP Server

```go
package main

import (
    "net/http"
    "io"
)

func handler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello, Web!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

## Key Components

**Handler Function:**
- Takes `http.ResponseWriter` and `*http.Request`
- Writes response back to client

**http.HandleFunc():**
- Registers a handler for a URL pattern

**http.ListenAndServe():**
- Starts the server on specified port

## Testing Your Server

Once running, test with:
- **Browser**: http://localhost:8080
- **curl**: `curl http://localhost:8080`
- **wget**: `wget -qO- http://localhost:8080`

## Your Task

Create a simple HTTP server that:
1. Responds with "Hello web World" 
2. Uses the `net/http` package
3. Runs on any port you choose
4. Can be tested with a browser or curl

This is the foundation of web development in Go!

```go
// Exercise: Set up a simple HTTP Server

// We will use the net/http library
// https://pkg.go.dev/net/http

// There are multiple ways of creating a http server in GO
// In this task, we are going to crete a http in your preferred way, so run free with the documentation :)
// I will provide the simplest solution, the only requirement for this exercise is that the server must 
// reply with a "Hello web World" response.

package main

import "net/http"
import "io"
import "log"

func main() {
   
}
```

<details>
<summary> Solution: </summary>

```go
// Exercise: Set up a simple HTTP Server

// We will use the net/http library
// https://pkg.go.dev/net/http

// There are multiple ways of creating a http server in GO
// In this task, we are going to crete a http in your preferred way, so run free with the documentation :)
// I will provide the simplest solution, the only requirement for this exercise is that the server must 
// reply with a "Hello web World" response.

package main

import "net/http"
import "io"
import "log"

func main() {
   // Set routing rules
   http.HandleFunc("/", WebFunc)

   //Use the default DefaultServeMux.
   err := http.ListenAndServe(":8080", nil)
   if err != nil {
           log.Fatal(err)
   }
}

func WebFunc(w http.ResponseWriter, r *http.Request) {
   io.WriteString(w, "Hello web World!")
}
  
```

</details>