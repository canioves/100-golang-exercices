# Exercise 63: Implicit Interface Implementation

## Go's Implicit Interface Implementation

Go uses **implicit interface implementation** - you don't explicitly declare that a type implements an interface. If a type has all the methods that an interface requires, it automatically implements that interface.

## No "implements" Keyword

Unlike other languages, Go has no `implements` keyword:

```go
// Java/C# style (NOT Go):
// class Rectangle implements Geometry { ... }

// Go style - just implement the methods:
type Rectangle struct { ... }
func (r Rectangle) area() float64 { ... }  // Now Rectangle implements Geometry!
```

## How Implicit Implementation Works

1. **Interface defines contract**: "Types must have these methods"
2. **Type provides methods**: Methods with exact same signatures
3. **Go compiler checks**: "Does this type satisfy the interface?"
4. **Automatic implementation**: No explicit declaration needed

## Method Receivers

When implementing interface methods, you typically use method receivers:

```go
type rect struct {
    width, height float64
}

// Method with receiver - implements interface method
func (r rect) area() float64 {
    return r.width * r.height
}
```

## Benefits of Implicit Implementation

- **Flexibility**: Existing types can implement new interfaces
- **Decoupling**: Interface and implementation can be in different packages
- **Simplicity**: Less boilerplate code
- **Retroactive implementation**: Add interface support to existing types

## Your Task

1. **Implement the `area()` method** for the `rect` type
2. **Use a method receiver**: `func (r rect) area() float64`
3. **Calculate rectangle area**: width × height
4. **Create and test** a rectangle instance

## Key Learning

By implementing the `area()` method, your `rect` type automatically satisfies the `geometry` interface - no explicit declaration needed!

```go
// Interfaces - Intro - Implicit implementation

// In golang, we don't use the "implements" keyword when using an interface.
// We don't have to say that explicitly, instead we need some type to satisfy the area() method.
// We say something implements an interface if it has a method with the exact signature.
// Inside our geometry interface, we have a signature called "area()"
// In this exercise, we are going to create a function with that signature!

package main

import (
	"fmt"
)

type geometry interface{
	area() float64
}

// A rectangle struct
type rect struct {
    width, height float64
}

// We are going to create a function with the area() signature.
// The function will be a receiver function (the receiver will be a rectangle "r")
// The function will be named area, without any arguments and will return a float64 value

	// The return value will be the area of the rectangle



func main() {
	// Create a new rectangle and store it in a variable
	
	// Print the variable
	fmt.Println(r.area())
}
```

<details>
<summary> Solution: </summary>

```go
// Interfaces - Intro - Implicit implementation

// In golang, we don't use the "implements" keyword when using an interface.
// We don't have to say that explicitly, instead we need some type to satisfy the area() method.
// We say something implements an interface if it has a method with the exact signature.
// Inside our geometry interface, we have a signature called "area()"
// In this exercise, we are going to create a function with that signature!

package main

import (
	"fmt"
)

type geometry interface{
	area() float64
}

// A rectangle struct
type rect struct {
    width, height float64
}

// We are going to create a function with the area() signature.
// The function will be a receiver function (the receiver will be a rectangle "r")
// The function will be named area, without any arguments and will return a float64 value
func (r rect) area() float64 {
	// The return value will be the area of the rectangle
    return r.width * r.height
}

func main() {
	// Create a new rectangle and store it in a variable
	r := rect{width: 3, height: 4}
	fmt.Println(r.area())
}
```

</details>