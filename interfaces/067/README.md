# Exercise 67: Type Assertions with Error Handling

## Handling Type Assertion Failures

When working with empty interfaces, type assertions can fail. It's important to handle these failures gracefully instead of letting your program panic.

## The Safe Approach

Always use the two-value form of type assertion:

```go
if value, ok := someInterface.(int); ok {
    // Successfully got an int
    fmt.Println("Value:", value)
} else {
    // Not an int, handle gracefully
    log.Println("Expected int, got something else")
}
```

## Why This Matters

In real applications, data might be:
- From external sources (APIs, files)
- User input
- Configuration files
- Dynamic at runtime

You can't always guarantee the types, so defensive programming is essential.

## Your Task

Build on the previous exercise but add proper error handling:
1. Use safe type assertions with the ok pattern
2. Handle cases where the type assertion fails
3. Use the `log` package to report errors
4. Prevent panics from invalid type assumptions

This teaches defensive programming practices essential for robust Go applications.

```go
// Interfaces -  Empty interfaces, bear in mind...
package main

import (
	"fmt"
	"log"
)

type human map[string]interface{}

// But there's an important thing to point out when it comes to retrieving and using a value from this map
// let's say that we want to get the "age" value and increment it by 1.
func main() {
	person := make(human)
	person["name"] = "Alice"
    person["age"] = 21
    person["height"] = 167.64

	// Type assertion for the integer. use .(int) for type assertion!
	
	// Check that the assertion was alright
	

	person["age"] = age + 1

    fmt.Printf("%+v", person)
}
```

<details>
<summary> Solution: </summary>

```go
// Interfaces -  Empty interfaces, bear in mind...
package main

import (
	"fmt"
	"log"
)

type human map[string]interface{}

// But there's an important thing to point out when it comes to retrieving and using a value from this map
// let's say that we want to get the "age" value and increment it by 1.
func main() {
	person := make(human)
	person["name"] = "Alice"
    person["age"] = 21
    person["height"] = 167.64

	// Type assertion for the integer. use .(int) for type assertion!
	age, ok := person["age"].(int)
	// Check that the assertion was alright
	if !ok {
		log.Fatal("could not assert value to int")
        return
	}

	person["age"] = age + 1

    fmt.Printf("%+v", person)
}
```

</details>