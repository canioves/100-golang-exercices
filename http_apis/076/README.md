# Exercise 76: Understanding ServeMux

## What is ServeMux?

`ServeMux` (HTTP request multiplexer) is Go's built-in HTTP router. It matches incoming requests to registered patterns and calls the appropriate handler.

## DefaultServeMux vs Custom ServeMux

**DefaultServeMux** (what you've been using):
```go
http.HandleFunc("/path", handler)  // Registers with DefaultServeMux
http.ListenAndServe(":8080", nil)  // nil = use DefaultServeMux
```

**Custom ServeMux**:
```go
mux := http.NewServeMux()
mux.HandleFunc("/path", handler)
http.ListenAndServe(":8080", mux)  // Use custom mux
```

## Why Use Custom ServeMux?

- **Isolation**: Multiple servers with different routes
- **Testing**: Easier to test with isolated mux
- **Control**: More explicit about routing
- **Security**: DefaultServeMux is global and can be modified by other packages

## Your Task

Create and use a custom ServeMux instead of the DefaultServeMux:
1. Create a new ServeMux with `http.NewServeMux()`
2. Register your handlers with the custom mux
3. Pass the custom mux to `ListenAndServe`

This gives you explicit control over HTTP routing.

```go

```

<details>
<summary> Solution: </summary>

```go
// Exercise: Let's talk about ServeMux!
package main

import "net/http"
import "log"

// In this exercise I want you to talk about serveMux, arm yourself with documentation first please!
// In a nutshell, servemux is an http request multiplexer (https://pkg.go.dev/net/http#ServeMux)
// In the last exercise, notice that we passed the 'nil' value to the .ListenAndServe() function. When a nil value is passed, the DefaultServeMux value is gathered.

// Let's supose we have the previous handler function:
func customHandler (w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Hello from customHandler!"))
}

func main() {  
  // Create a ServeMux variable called 'mux'
  // And assign it the value of a new servemux


  //Use the HandleFunc() function, as the 1st argument you will have to put the root path "/" and the 2nd argument will be the customHandler function
  

  // Start the server with ListenAndServe() function, and as the second parameter use the "mux" servemux you have created!
  server := 
  if (server != nil){
    log.Print("Cannot start sever")
  }
}

  //Use the HandleFunc() function, as the 1st argument you will have to put the root path "/" and the 2nd argument will be the customHandler function
  mux.HandleFunc("/", customHandler)

  // Start the server with ListenAndServe() function, and as the second parameter use the "mux" servemux you have created!
  server := http.ListenAndServe(":8080", mux)
  if (server != nil){
    log.Print("Cannot start sever")
  }
}
```

</details>