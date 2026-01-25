# Exercise 57: Context Chaining and Immutability

## Context Chaining

Context values can be chained - you can create new contexts that inherit from parent contexts and add or modify values. This creates a chain of contexts, each building on the previous one.

## How Context Chaining Works

```go
ctx1 := context.Background()
ctx2 := context.WithValue(ctx1, "key1", "value1")  
ctx3 := context.WithValue(ctx2, "key2", "value2")
ctx4 := context.WithValue(ctx3, "key1", "newvalue") // Shadows "key1" from ctx2
```

## Context Immutability

**Important**: Contexts are immutable! When you call `context.WithValue()`, you don't modify the existing context - you create a new one.

```go
originalCtx := context.WithValue(context.Background(), "name", "John")
newCtx := context.WithValue(originalCtx, "name", "Jane")

// originalCtx still has "John"
// newCtx has "Jane"  
// They are separate contexts!
```

## Value Shadowing

When you add a value with the same key to a child context, it "shadows" (hides) the parent's value:

```go
parent := context.WithValue(context.Background(), "user", "Alice")
child := context.WithValue(parent, "user", "Bob")

fmt.Println(parent.Value("user")) // "Alice"
fmt.Println(child.Value("user"))  // "Bob" (shadows Alice)
```

## Your Task

1. **Complete the `doAnother` function** that prints the same key value
2. **In `doSomething`, create a new context** using `context.WithValue()` with the same key but different value
3. **Call `doAnother` with the new context**
4. **Observe the behavior**: Does the value change in the child context?
5. **Test immutability**: Check if the original context's value changed

## Key Learning Points

- Contexts form a tree structure
- Child contexts can shadow parent values
- Original contexts remain unchanged (immutable)
- Each `WithValue` call creates a new context

```go
// Contexts - Chaining contexts
package main

import (
	"context" 
	"fmt"
)

// Create another function called doAnother, as a context as it's only argument. Like the last exercise, it will print the same key as the doSomething function


// Modify the doSomething function
func doSomething(ctx context.Context) {
	// Here, print the ctx.Value() call with the key variable you added in the main function
	fmt.Println("doSomething: Name's value is", ctx.Value("Name"))
	// Create a varibale anotherCtx and use the context.WithValue(a,b,c) function to pass in the father context as A (ctx), the same key name as b, and a different value as c

	// call the doAnother function with the anotherCtx as it's only argument.
	// See it's behaviour: Have the values chainged? Are context mutable?

}

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx,"Name","John")

	doSomething(ctx)
}
```

<details>
<summary> Solution: </summary>

```go
// Contexts - Chaining contexts
package main

import (
	"context" 
	"fmt"
)

// Create another function called doAnother, as a context as it's only argument. Like the last exercise, it will print the same key as the doSomething function
func doAnother(ctx context.Context) {
	// Here, print the ctx.Value() call with the key variable you added in the main function
	fmt.Println("doSomething: Name's value is", ctx.Value("Name"))
}

// Modify the doSomething function
func doSomething(ctx context.Context) {
	// Here, print the ctx.Value() call with the key variable you added in the main function
	fmt.Println("doSomething: Name's value is", ctx.Value("Name"))
	// Create a varibale anotherCtx and use the context.WithValue(a,b,c) function to pass in the father context as A (ctx), the same key name as b, and a different value as c
	anotherCtx := context.WithValue(ctx,"Name","Mary")
	// call the doAnother function with the anotherCtx as it's only argument.
	// See it's behaviour: Have the values chainged? Are context mutable?
	doAnother(anotherCtx)
}

func main() {
	ctx := context.Background()

	// here we will assign to ctx the context.WithValue() method, the first argument will be the context, the second your key
	// (this key will be referenced as the "ctx.Value(key)" in the function doSomething above, and the third the value of your key (for example key = Name, value = John)
	ctx = context.WithValue(ctx,"Name","John")
	// call the function with the empty context created before as it's only argument
	doSomething(ctx)
}
```

</details>