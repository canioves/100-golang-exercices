# Exercise 75: Multiple HTTP Routes with HandleFunc

## Multiple Route Handling

Real web applications need multiple routes. Go's `http.HandleFunc()` lets you register different handlers for different URL patterns.

## Registering Multiple Routes

```go
func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/contact", contactHandler)
    
    http.ListenAndServe(":8080", nil)
}
```

## Handler Function Signature

All handler functions must match this signature:
```go
func handlerName(w http.ResponseWriter, r *http.Request) {
    // Handle the request
}
```

## DefaultServeMux

When you use `http.HandleFunc()` and `http.ListenAndServe()` with `nil`, Go uses the `DefaultServeMux` (default HTTP request multiplexer) to route requests.

## Route Patterns

- **Exact match**: `/users` matches only `/users`
- **Prefix match**: `/static/` matches `/static/css/style.css`
- **Root pattern**: `/` matches everything not matched by other patterns

## Your Task

Create an HTTP server with multiple routes:
1. Create `handler_1` function that writes "Hello from HandleFunc #1"
2. Create `handler_2` function that writes "Hello from HandleFunc #2"
3. Register both handlers for different URL patterns
4. Test that each route works correctly

## Testing Multiple Routes

Visit different URLs to test:
- Each route should respond with its specific message
- Different handlers for different paths

This demonstrates building multi-page web applications.

```go
// Exercise: Set up a simple HTTP Server
package main

import "net/http"
import "log"

// In this exercise we will register the handler function for a given pattern!
// We will use this HandleFunc BEWARE IT'S NOT THE SAME THAN HandlerFunc!!!!!
// We will register these patterns (uri) in the defaultServerMux, we will talk about it later

// Create a handler function named handler_1 that will write "Hello from Handlefunc #1"


//Create a handler function named handler_2 that will write "Hello from handlefunc #2"


func main() {  
  // This ListenAndServe will get 2 parameters: 1 - Address (:8080) in this case
  //                                            2 - nil
  
  // Now, use the http.HandleFunc() to register the handler_1 function to the "/handler1" pattern
  // And use the same method to register the handler_2 function to "/handler2" pattern
  
  server := http.ListenAndServe(":8080", nil)
  if (server != nil){
    log.Print("Cannot start sever")
  }
}
```

<details>
<summary> Solution: </summary>

```go
// Exercise: Set up a simple HTTP Server
package main

import "net/http"
import "log"

// In this exercise we will register the handler function for a given pattern!
// We will use this HandleFunc BEWARE IT'S NOT THE SAME THAN HandlerFunc!!!!!
// We will register these patterns (uri) in the defaultServerMux, we will talk about it later

// Create a function named handler_1 that will write "Hello from Handlefunc #1"
func handler_1(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Hello from Handlefunc #1"))
}

//Create a function named handler_2 that will write "Hello from handlefunc #2"
func handler_2(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Hello from Handlefunc #2"))
}

func main() {  
  // This ListenAndServe will get 2 parameters: 1 - Address (:8080) in this case
  //                                            2 - nil
  
  // Now, use the http.HandleFunc() to register the handler_1 function to the "/handler1" pattern
  // And use the same method to register the handler_2 function to "/handler2" pattern
  http.HandleFunc("/handler1", handler_1)
  http.HandleFunc("/handler2", handler_2)

  server := http.ListenAndServe(":8080", nil)
  if (server != nil){
    log.Print("Cannot start sever")
  }
}
```

</details>