# Exercise 62: Introduction to Interfaces

## What are Interfaces?

An interface in Go is a type that defines a set of method signatures. It specifies what methods a type must have, but not how those methods are implemented.

Think of interfaces as contracts - they define what behavior a type must provide.

## Interface Syntax

```go
type InterfaceName interface {
    Method1() ReturnType
    Method2(param Type) ReturnType
    // ... more method signatures
}
```

## Why Use Interfaces?

1. **Abstraction**: Define behavior without implementation details
2. **Flexibility**: Different types can implement the same interface
3. **Testing**: Easy to create mock implementations
4. **Decoupling**: Code depends on behavior, not concrete types

## Example

```go
type Writer interface {
    Write([]byte) (int, error)
}

// Any type with a Write method implements Writer
// Files, network connections, buffers, etc.
```

## Interface Design Principles

- **Keep them small**: Single-method interfaces are common and powerful
- **Focus on behavior**: What the type can do, not what it is
- **Name by capability**: Often end in "-er" (Writer, Reader, Stringer)

## Your Task

Create your first interface:

1. **Define an interface called "Geometry"** (note the capitalization)
2. **Add a method signature `area()`** that returns `float64`
3. **Understand**: This interface defines that any type implementing it must have an `area()` method

## Key Concepts

- Interface defines the "what" (method signatures)
- Types provide the "how" (method implementations)  
- This is just the interface definition - no implementation yet!

```go
// Interfaces - Intro

// An interface type in Go is kind of like a definition. It defines and describes the exact methods that some other type must have.
// Interfaces are a concept that at first may need some time to be understood, take all your time.

package main

import (
	"fmt"
	"math"
)

// This first exercise we are going to create an interface, no execution!
// Creating our first interface:
// declare an interface called "Geometry"

	// create a function signature area(), it's return type will be float64
	


// A rectangle struct
type rect struct {
    width, height float64
}

func main() {
	ctx := context.Background()
	doSomething(ctx)
}
```

<details>
<summary> Solution: </summary>

```go
// Interfaces - Intro

// An interface type in Go is kind of like a definition. It defines and describes the exact methods that some other type must have.
// Interfaces are a concept that at first may need some time to be understood, take all your time.

package main

import (
	"fmt"
	"math"
)

// This first exercise we are going to create an interface, no execution!
// Creating our first interface:
// declare an interface called "Geometry"
type geometry interface{
	// create a function signature area(), it's return type will be float64
	area() float64
}

// A rectangle struct
type rect struct {
    width, height float64
}

func main() {
	ctx := context.Background()
	doSomething(ctx)
}
```

</details>