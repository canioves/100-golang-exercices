# Exercise 77: ServeMux with Stateful Handlers

## Stateful HTTP Handlers

Sometimes handlers need to maintain state (counters, databases, configuration). You can create struct-based handlers that implement the `http.Handler` interface.

## Struct-Based Handlers

```go
type Counter struct {
    mu sync.Mutex
    n  int
}

func (c *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.n++
    fmt.Fprintf(w, "Count: %d", c.n)
}
```

## Thread Safety

When multiple requests access shared state, use synchronization:
- **sync.Mutex**: For exclusive access
- **sync.RWMutex**: For read-heavy workloads
- **atomic**: For simple counters

## Using with ServeMux

```go
counter := &Counter{}
mux := http.NewServeMux()
mux.Handle("/count", counter)  // Note: Handle, not HandleFunc
```

## Your Task

Create a Counter struct that:
1. Implements the `http.Handler` interface
2. Maintains an internal counter
3. Increments the counter on each request
4. Returns the current count
5. Uses proper synchronization for thread safety

This demonstrates stateful web handlers and concurrent safety.

```go
// Exercise: Let's talk about ServeMux! And receivers in functions.
package main

import "net/http"
import "log"
import "fmt"

// We will have this Counter struct
type Counter struct {
  n int
}
// Imagine we have this counter.
// Bear in mind that this first part of the function
//   (ctr *Counter) is called a receiver, this means that the function caller WILL be a Counter element
func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  ctr.n++
  fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

func main() {  
  // Create a ServeMux variable called 'mux'
  // And assign it the value of a new servemux
  mux := http.NewServeMux()

  // Register a new Counter element in a variable called 'ctr' (tip: use the new() function!)

  // Now we will use the Handle() function with the mux variable.
  // The first argument will be the "/counter" pattern
  // The second argument will be the ctr function, in this case we can see very clearly how the handler acts as middleware. 
  

  // Start the server with ListenAndServe() function, and as the second parameter use the "mux" servemux you have created!
  server := http.ListenAndServe(":8080", mux)
  if (server != nil){
    log.Print("Cannot start sever")
  }
}
```

<details>
<summary> Solution: </summary>

```go
// Exercise: Let's talk about ServeMux! And receivers in functions.
package main

import "net/http"
import "log"
import "fmt"

// We will have this Counter struct
type Counter struct {
  n int
}

// Imagine we have this counter.
func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  ctr.n++
  fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

func main() {  
  // Create a ServeMux variable called 'mux'
  // And assign it the value of a new servemux
  mux := http.NewServeMux()

  // Register a new Counter element in a variable called 'ctr'
  ctr := new(Counter) // This will initialize the value to 0!

  // Now we will use the Handle() function
  // The first argument will be the "/counter" pattern
  // The second argument will be the ctr function, in this case we can see very clearly how the handler acts as middleware. 
  mux.Handle("/counter", ctr)

  // Start the server with ListenAndServe() function, and as the second parameter use the "mux" servemux you have created!
  server := http.ListenAndServe(":8080", mux)
  if (server != nil){
    log.Print("Cannot start sever")
  }
}
```

</details>