# Exercise 60: Context with Deadline

## What is context.WithDeadline()?

`context.WithDeadline()` creates a context that automatically cancels at a specific time. It's perfect when you know exactly when an operation should stop.

```go
deadline := time.Now().Add(5 * time.Second)
ctx, cancel := context.WithDeadline(parentCtx, deadline)
defer cancel()
```

## Deadline vs Timeout

| WithDeadline | WithTimeout |
|-------------|-------------|
| Specific time point | Duration from now |
| `time.Now().Add(duration)` | Just the duration |
| More control | Simpler to use |

## How Deadlines Work

1. **Set absolute time**: "Cancel at exactly 3:30 PM"
2. **Automatic cancellation**: Context cancels itself when deadline passes
3. **Done channel closes**: `ctx.Done()` receives notification
4. **Error available**: `ctx.Err()` returns deadline exceeded error

## Context Errors

When a context is cancelled, `ctx.Err()` returns the reason:

```go
case <-ctx.Done():
    err := ctx.Err()
    if err == context.DeadlineExceeded {
        fmt.Println("Operation timed out")
    } else if err == context.Canceled {
        fmt.Println("Operation was cancelled")
    }
```

## Racing Against Time

This exercise demonstrates a race condition:
- **Loop tries to send** 3 numbers, sleeping 1 second between each
- **Context cancels after 1.5 seconds**
- **Which wins?** The loop (3+ seconds) or the deadline (1.5 seconds)?

## Your Task

1. **Create context with deadline** using the provided 1.5-second deadline
2. **Modify the loop** to check for context cancellation
3. **Handle both cases** in select:
   - Sending numbers to channel
   - Receiving cancellation signal
4. **Observe the behavior**: Does the loop complete or get cancelled?

## Key Learning

Deadlines are perfect for operations where you know exactly when to give up. The context will automatically cancel, making your code responsive to time constraints.

```go
// Contexts - Canceling a context (withDeadline)
package main

import (
	"context" 
	"fmt"
	"time"
)

// Modify the doSomething function
func doSomething(ctx context.Context) {
	// Deadline of 1,5 seconds
	deadline := time.Now().Add(1500 * time.Millisecond)
	// Repeat the same operation for cancelling the context, but this time instead of WithCancel() we will use WithDeadline(a,b) the a argument will be the parent context ctx
	// and the b argument will be the deadline above (1,5 seconds in this case)

	// defer cancelCtx when time has passed

	// Make a new unbuffered channel of integers and assign it to printCh
	printCh := make(chan int)
	// call the doAnother function as goroutine
	go doAnother(ctx, printCh)
	// Modify the loop, for each number, wait for one second.
	// And whenever it receives a ctx.Done(), just break the selection.
	for num := 1; num <= 3; num++ {
		select {
			// case 1 (receive a number to printCh channel)

				// then sleep for a second

			// case 2 (received ctx.done())

				// break

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
// Contexts - Canceling a context (withDeadline)
package main

import (
	"context" 
	"fmt"
	"time"
)

// Modify the doSomething function
func doSomething(ctx context.Context) {
	// Deadline of 1,5 seconds
	deadline := time.Now().Add(1500 * time.Millisecond)
	// Repeat the same operation for cancelling the context, but this time instead of WithCancel() we will use WithDeadline(a,b) the a argument will be the parent context ctx
	// and the b argument will be the deadline above (1,5 seconds in this case)
	ctx, cancelCtx := context.WithDeadline(ctx, deadline)
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