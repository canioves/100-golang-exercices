# Exercise 61: Context with Timeout

## What is context.WithTimeout()?

`context.WithTimeout()` creates a context that automatically cancels after a specified duration. It's the most commonly used timeout mechanism in Go.

```go
ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
defer cancel()
```

## WithTimeout vs WithDeadline

```go
// These are equivalent:
deadline := time.Now().Add(5 * time.Second)
ctx1, cancel1 := context.WithDeadline(parentCtx, deadline)

ctx2, cancel2 := context.WithTimeout(parentCtx, 5*time.Second)
```

`WithTimeout` is just a convenience function that calculates the deadline for you.

## Common Timeout Patterns

**HTTP requests:**
```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
req = req.WithContext(ctx)
```

**Database queries:**
```go
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
rows, err := db.QueryContext(ctx, "SELECT * FROM users")
```

**API calls:**
```go
ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
defer cancel()
response, err := client.GetWithContext(ctx, url)
```

## Why Always defer cancel()?

Even if the timeout hasn't been reached, you should always call `cancel()` to:
1. **Free resources**: Clean up the timeout timer
2. **Signal completion**: Tell child contexts to stop
3. **Best practice**: Ensures proper cleanup regardless of how function exits

## Timeout Behavior

1. **Normal completion**: Operation finishes before timeout
2. **Timeout**: Context cancels itself after duration
3. **Manual cancel**: You can still call cancel() early
4. **Error reporting**: `ctx.Err()` returns `context.DeadlineExceeded`

## Your Task

1. **Create a context with 1.5-second timeout**
2. **Complete the select statement** to handle both sending and timeout
3. **Observe the race**: Will the 3-second loop complete, or will the 1.5-second timeout win?
4. **Handle the timeout gracefully** by breaking out of the loop

## Expected Behavior

Since the timeout (1.5s) is shorter than the time needed to complete the loop (3+ seconds), the context should cancel before all numbers are sent.

```go
// Contexts - Canceling a context (withTimeout)
package main

import (
	"context" 
	"fmt"
	"time"
)

// Modify the doSomething function
func doSomething(ctx context.Context) {
	// Repeat the same operation for cancelling the context, but this time instead of WithDeadline(a,b) we will use WithTimeout(ctx,time.Time)
	// We will assign it a time of 1,5 seconds
	
	// defer cancelCtx when time has passed
	defer cancelCtx()
	// Make a new unbuffered channel of integers and assign it to printCh
	printCh := make(chan int)
	// call the doAnother function as goroutine
	go doAnother(ctx, printCh)
	// Modify the loop, for each number, wait for one second.
	// And whenever it receives a ctx.Done(), just break the selection.
	for num := 1; num <= 3; num++ {
		select {
			// case 1 (receive a number to printCh channel)
			case printCh <- num:
				// then sleep for a second
				time.Sleep(1 * time.Second)
			// case 2 (received ctx.done())
			case <-ctx.Done():
				// break
				break
		}
	}
	
	// sleep for 100 ms
	time.Sleep(1000 * time.Millisecond)
	// print that doSomething has finished
	fmt.Printf("doSomething: finished\n")
}



func doAnother(ctx context.Context, printCh <-chan int) {
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

<details>
<summary> Solution: </summary>

```go
// Contexts - Canceling a context (withTimeout)
package main

import (
	"context" 
	"fmt"
	"time"
)

// Modify the doSomething function
func doSomething(ctx context.Context) {
	// Repeat the same operation for cancelling the context, but this time instead of WithDeadline(a,b) we will use WithTimeout(ctx,time.Time)
	// We will assign it a time of 1,5 seconds
	ctx, cancelCtx := context.WithTimeout(ctx, 1500*time.Millisecond)
	// defer cancelCtx when time has passed
	defer cancelCtx()
	// Make a new unbuffered channel of integers and assign it to printCh
	printCh := make(chan int)
	// call the doAnother function as goroutine
	go doAnother(ctx, printCh)
	// Modify the loop, for each number, wait for one second.
	// And whenever it receives a ctx.Done(), just break the selection.
	for num := 1; num <= 3; num++ {
		select {
			// case 1 (receive a number to printCh channel)
			case printCh <- num:
				// then sleep for a second
				time.Sleep(1 * time.Second)
			// case 2 (received ctx.done())
			case <-ctx.Done():
				// break
				break
		}
	}
	
	// sleep for 100 ms
	time.Sleep(1000 * time.Millisecond)
	// print that doSomething has finished
	fmt.Printf("doSomething: finished\n")
}



func doAnother(ctx context.Context, printCh <-chan int) {
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