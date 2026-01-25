# Exercise 74: HTTP Handler Interface

## Understanding http.Handler Interface

The `http.Handler` interface is at the core of Go's HTTP handling. Any type that implements this interface can handle HTTP requests.

## The Handler Interface

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Any type with a `ServeHTTP` method automatically implements this interface.

## Creating Custom Handlers

**Method 1: Custom Type**
```go
type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from custom handler!")
}
```

**Method 2: Handler Function**
```go
func myHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from handler function!")
}

// Convert function to Handler using http.HandlerFunc
handler := http.HandlerFunc(myHandler)
```

## Handler vs HandlerFunc

- **http.Handler**: Interface requiring ServeHTTP method
- **http.HandlerFunc**: Type that converts functions to handlers
- **http.HandleFunc()**: Convenience function for registering handlers

## Your Task

This exercise focuses on understanding the `http.Handler` interface:
1. Create a custom type that implements `http.Handler`
2. Implement the `ServeHTTP` method
3. Understand how this interface enables flexible HTTP handling

This builds foundation knowledge for more advanced HTTP middleware and routing.

```go
// Exercise: Set up a simple HTTP Server

// To understand the http library better, we will do another server and handle it's router, one step at a time!

package main

import "net/http"
import "log"

// First of all, we are going to create an HTTP handler
  // The http.Handler is an interface, according to the docs like this>
  /*
    type Handler interface {
      ServeHTTP(ResponseWriter, *Request)
    }
    // This means that an object to be a handler MUST HAVE a ServeHTTP method with that signature [ServeHTTP(ResponseWriter, *Request)]
  */

// To put things simple, we will create handlers as functions, and them we will wrap around the ServeHTTP Method signature
// Create here a function called 'myCustomHandler' That has the same arguments as the ServeHTTP method:
func {
  
}
// Great! But how do we wrap around the serveHTTP method now?
// That will be easy, we will use the HandlerFunc: https://cs.opensource.google/go/go/+/refs/tags/go1.19.2:src/net/http/server.go;l=2105
// This wrapper is all we need in order to convert the function in an HTTP handler!! Easy isn't it? 

func main() {  
  // Create a server with the .ListenAndServe() function (more info here: https://pkg.go.dev/net/http#Server.ListenAndServe)
  // Also here is an excellent resource for http dissection: https://echorand.me/posts/golang-dissecting-listen-and-serve/#what-is-defaultservemux

  // This ListenAndServe will get 2 parameters: 1 - the uri (root in this case "/" and)
  //                                            2 - an http.Handler(). Or a custom function wrapped by the http.HandlerFunc(function)
  
  // This http.ListenAndServe should be called with port 8080 and use the function you created before as the Handler!
  server := 
  if (server != nil){
    log.Print("Cannot start sever")
  }
}
```

<details>
<summary> Solution: </summary>

```go
// Exercise: Set up a simple HTTP Server

// To understand the http library better, we will do another server and handle it's router, one step at a time!

package main

import "net/http"
import "log"

// First of all, we are going to create an HTTP handler
  // The http.Handler is an interface, according to the docs like this>
  /*
    type Handler interface {
      ServeHTTP(ResponseWriter, *Request)
    }
    // This means that an object to be a handler MUST HAVE a ServeHTTP method with that signature [ServeHTTP(ResponseWriter, *Request)]
  */

// To put things simple, we will create handlers as functions, and them we will wrap around the ServeHTTP Method signature
// Create here a function called 'myCustomHandler' That has the same arguments as the ServeHTTP method:
func myCustomHandler(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Hello Web!"))
}

// Great! But how do we wrap around the serveHTTP method now?
// That will be easy, we will use the HandlerFunc: https://cs.opensource.google/go/go/+/refs/tags/go1.19.2:src/net/http/server.go;l=2105
// This wrapper is all we need in order to convert the function in an HTTP handler!! Easy isn't it? 

func main() {  
  // Create a server with the .ListenAndServe() function (more info here: https://pkg.go.dev/net/http#Server.ListenAndServe)
  // Also here is an excellent resource for http dissection: https://echorand.me/posts/golang-dissecting-listen-and-serve/#what-is-defaultservemux

  // This ListenAndServe will get 2 parameters: 1 - the uri (root in this case "/" and)
  //                                            2 - an http.Handler(). Or a custom function wrapped by the http.HandlerFunc(function)
  
  // This http.ListenAndServe should be called with port 8080 and use the function you created before as the Handler!
  server := http.ListenAndServe(":8080", http.HandlerFunc(myCustomHandler))
  if (server != nil){
    log.Print("Cannot start sever")
  }
}
```

</details>