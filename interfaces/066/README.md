# Exercise 66: Type Assertions with Empty Interfaces

## The Problem with Empty Interfaces

When you store values in `interface{}`, you lose type information. To use the values, you need **type assertions** to convert back to concrete types.

## Type Assertion Syntax

```go
// Basic type assertion
value := someInterface.(int)

// Safe type assertion with ok check  
value, ok := someInterface.(int)
if !ok {
    fmt.Println("Not an int!")
}
```

## Why Type Assertions Are Needed

```go
var x interface{} = 42
// x + 1  // ERROR: can't add to interface{}
y := x.(int)  // Type assertion
result := y + 1  // OK: now it's an int
```

## Safe vs Unsafe Assertions

**Unsafe** (panics if wrong type):
```go
age := person["age"].(int)  // Panic if not int
```

**Safe** (returns bool):
```go
if age, ok := person["age"].(int); ok {
    fmt.Println("Age:", age + 1)
} else {
    fmt.Println("Age is not an integer")
}
```

## Your Task

Work with a `human` map containing mixed types. Practice:
1. Storing different types (string, int) in the map
2. Using type assertions to retrieve and modify values
3. Handling the case where you need to increment an age value

This demonstrates both the power and complexity of working with empty interfaces.

```go
// Interfaces -  Empty interfaces, bear in mind...
package main

import (
	"fmt"
)

type human map[string]interface{}

// But there's an important thing to point out when it comes to retrieving and using a value from this map
// let's say that we want to get the "age" value and increment it by 1.
func main() {
	person := make(human)
	person["name"] = "Alice"
    person["age"] = 21
    person["height"] = 167.64

	// Try to update the person age, the new value will be the existing one plus 1
	// (When you run the program, see what happens), what's the error it arises?
	person["age"] = 

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
)

type human map[string]interface{}

// But there's an important thing to point out when it comes to retrieving and using a value from this map
// let's say that we want to get the "age" value and increment it by 1.
func main() {
	person := make(human)
	person["name"] = "Alice"
    person["age"] = 21
    person["height"] = 167.64

	// Try to update the person age, the new value will be the existing one plus 1
	// (When you run the program, see what happens), what's the error it arises?
	person["age"] = person["age"] + 1

    fmt.Printf("%+v", person)
}
```

</details>