# Exercise 59: Context Cancellation

## What is Context Cancellation?

Context cancellation is a mechanism to signal that work should be stopped. It's essential for:
- **Resource cleanup**: Stop operations when they're no longer needed
- **Timeout handling**: Cancel slow operations
- **User cancellation**: Stop work when users cancel requests
- **Graceful shutdown**: Clean up when the application is stopping

## context.WithCancel()

`context.WithCancel()` creates a cancellable context:

```go
ctx, cancel := context.WithCancel(parentCtx)
defer cancel() // Always call cancel to free resources
```

Returns:
- **ctx**: The new cancellable context
- **cancel**: A function that cancels the context when called

## The Done Channel

Cancelled contexts signal through their `Done()` channel:

```go
select {
case <-ctx.Done():
    fmt.Println("Context cancelled!")
    return ctx.Err() // Returns cancellation reason
case result := <-workChannel:
    fmt.Println("Work completed:", result)
}
```

## Select with Context

The common pattern for cancellable operations:

```go
for {
    select {
    case <-ctx.Done():
        fmt.Println("Stopping due to cancellation")
        return
    case data := <-dataChannel:
        fmt.Println("Processing:", data)
    }
}
```

## Why Use Select with Context?

- **Responsive cancellation**: Check for cancellation between operations
- **Clean shutdown**: Exit loops and functions when cancelled
- **Resource management**: Stop goroutines that are no longer needed

## Your Task

This is a complex exercise involving:

1. **Create a cancellable context** using `context.WithCancel()`
2. **Create an unbuffered int channel** for communication
3. **Start a goroutine** (`doAnother`) that waits for cancellation or data
4. **Send data through the channel** (numbers 1-3)
5. **Cancel the context** to stop the goroutine
6. **Use select in the goroutine** to handle both cancellation and data

## Key Concepts

- `ctx.Done()` returns a channel that closes when context is cancelled
- Always call the cancel function to clean up resources
- Use select to make your goroutines responsive to cancellation

```go
// Contexts - Canceling a context
package main

import (
	"context" 
	"fmt"
	"time"
)

// Modify the doSomething function
func doSomething(ctx context.Context) {
	// use the context.WithCancel() function with the ctx as the only argument. Assign it to two values, first the ctx(this will be the parent function) and then the cancelCtx which will be a new empty function.

	// Make a new unbuffered channel of integers and assign it to printCh

	// call the doAnother function as goroutine

	// create a loop of 3 integer iterations from 1 to 3 which will be sent to the printCh channel!
	for num := 1; num <= 3; num++ {
		printCh <- num
	}
	// call the cancelCtx() function

	// sleep for 100 ms
	time.Sleep(100 * time.Millisecond)
	// print that doSomething has finished
	fmt.Printf("doSomething: finished\n")
}


// Modify the doAnother function,
// First add a second argument that will receive a int value from a channel into a variable called printCh
func doAnother(ctx context.Context, printCh <-chan int) {
	// For and select in conjuntion is a great way to operate with channels
	// The select statement lets a goroutine wait on multiple communication operations.
	// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready. 
	for {
		select {
		// first case will be reciving a ctx.Done() call, if this happens we will handle errors and abort the doAnother function.
		case :
			
		// Second case will receive a value from printCh and assign that to a variable called num, after that print the num variable
		case :
			fmt.Printf()
		}
	}
}

func main() {
	ctx := context.Background()
	doSomething(ctx)
}
```

<details>
<summary> Solution: </summary>

```go
// Contexts - Canceling a context
package main

import (
	"context" 
	"fmt"
	"time"
)

// Modify the doSomething function
func doSomething(ctx context.Context) {
	// use the context.WithCancel() function with the ctx as the only argument. Assign it to two values, first the ctx(this will be the parent function) and then the cancelCtx which will be a new empty function.

	ctx, cancelCtx := context.WithCancel(ctx)
	// Make a new unbuffered channel of integers and assign it to printCh
	printCh := make(chan int)
	// call the doAnother function as goroutine
	go doAnother(ctx, printCh)
	// create a loop of 3 integer iterations from 1 to 3 which will be sent to the printCh channel!
	for num := 1; num <= 3; num++ {
		printCh <- num
	}
	// call the cancelCtx() function
	cancelCtx() // cancel when we are done counting numbers
	// sleep for 100 ms
	time.Sleep(1000 * time.Millisecond)
	// print that doSomething has finished
	fmt.Printf("doSomething: finished\n")
}


// Modify the doAnother function,
// First add a second argument that will receive a int value from a channel into a variable called printCh
func doAnother(ctx context.Context, printCh <-chan int) {
	// For and select in conjuntion is a great way to operate with channels
	// The select statement lets a goroutine wait on multiple communication operations.
	// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready. 
	for {
		select {
		// first case will be reciving a ctx.Done() call, if this happens we will handle errors and abort the doAnother function.
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}
			fmt.Printf("doAnother: finished\n")
			return
		// Second case will receive a value from printCh and assign that to a variable called num, after that print the num variable
		case num := <-printCh:
			fmt.Printf("doAnother: %d\n", num)
		}
	}
}

func main() {
	ctx := context.Background()
	doSomething(ctx)
}
```

</details>