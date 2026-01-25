# Exercise 73: Dynamic HTTP Routes

## Dynamic Route Responses

Instead of hardcoding responses, you can make your server dynamically respond based on the requested URL path.

## Accessing Request Information

The `*http.Request` parameter contains information about the request:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path        // "/some/path"
    method := r.Method        // "GET", "POST", etc.
    host := r.Host           // "localhost:8080"
}
```

## HTML Escaping

When including user input (like URL paths) in responses, use `html.EscapeString()` to prevent XSS attacks:

```go
safeOutput := html.EscapeString(userInput)
```

## Dynamic Response Example

```go
func handler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path
    response := fmt.Sprintf("Hello, %s", html.EscapeString(path))
    fmt.Fprintf(w, response)
}
```

## Your Task

Create an HTTP server where:
1. The handler responds with "Hello, [path]" 
2. The [path] part is the actual URL path requested
3. If you visit `/bar`, it responds "Hello, /bar"
4. If you visit `/newpath`, it responds "Hello, /newpath"
5. Use proper HTML escaping for security

## Testing Different Paths

- http://localhost:8080/bar → "Hello, /bar"
- http://localhost:8080/test → "Hello, /test"  
- http://localhost:8080/anything → "Hello, /anything"

This demonstrates how to create dynamic web responses based on the request.

```go
// Exercise: Set up a simple HTTP Server

// We will use the net/http library
// https://pkg.go.dev/net/http

// Like with ExpressJS, we will make a web server, and the web server will serve the "/bar" route
// and the response should be "Hello, /var". BUT don't hardcode the URI!
// Changing the http.HandleFunc 1st variable (the URI), the message served should also change.
// Example: If I set my webserver at /newpath, my response will be "Hello, /newpath"

package main

import "net/http"
import "fmt"
import "html"
import "log"

func main() {

  http.HandleFunc(/**/ , func(w http.ResponseWriter, r *http.Request) {
	  
  })

  log.Fatal(http.ListenAndServe(":8080", nil))
}
```

<details>
<summary> Solution: </summary>

```go
// Exercise: Set up a simple HTTP Server

// We will use the net/http library
// https://pkg.go.dev/net/http

// Like with ExpressJS, we will make a web server, and the web server will serve the "/bar" route
// and the response should be "Hello, /var". BUT don't hardcode the URI!
// Changing the http.HandleFunc 1st variable (the URI), the message served should also change.
// Example: If I set my webserver at /newpath, my response will be "Hello, /newpath"

package main

import "net/http"
import "fmt"
import "html"
import "log"

func main() {

  http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  log.Fatal(http.ListenAndServe(":8080", nil))
}
```

</details>