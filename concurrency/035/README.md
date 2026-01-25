# Exercise 35: Channel Synchronization

## The Problem with Goroutines and the `main()` function

When you start a goroutine, the main function doesn't automatically wait for it to complete. If the main function ends, the entire program terminates - even if goroutines are still running.

## Synchronization Patterns

There are several ways to wait for goroutines to complete:
1. **`time.Sleep()`** - crude, unreliable (what if task takes longer?)
2. **Channel synchronization** - precise, reliable
3. **sync.WaitGroup** - for multiple goroutines (you'll learn this later, also reliable)

For now we've seen mostly `time.Sleep()` in order to learn and see the go routine behaviours
in our path to become great engineers we must adapt and learn elegant and reliable solutions.

## Channel Synchronization

Using a channel to signal completion is a clean, idiomatic Go pattern:

```go
func worker(done chan bool) {
    // Do work...
    done <- true  // Signal completion
}

func main() {
    done := make(chan bool)
    go worker(done)
    <-done  // Wait for completion signal
}
```

## How It Works

1. **Create channel**: Usually `chan bool` for simple signaling
2. **Start goroutine**: Pass the channel to the goroutine function
3. **Goroutine signals**: Sends a value when work is complete
4. **Main waits**: Receives from channel (blocks until signal arrives)

## Why This is Better Than Sleep

- **Precise timing**: Waits exactly as long as needed
- **No guessing**: No need to estimate how long work takes
- **Reliable**: Works regardless of system load or work complexity
- **Clear intent**: Code clearly shows synchronization purpose

## Signal Values

Common patterns for signaling:
- `chan bool` - just send `true` when done
- `chan struct{}` - more memory efficient, send `struct{}{}`
- `chan error` - can signal completion AND return error status

## Your Task

Look at the `main.go` file. The goroutine is already set up to send a completion signal. You need to add the code that waits for this signal in the main function.

## Expected Behavior

When completed correctly:
1. The program starts the `task` goroutine
2. Main function waits for the completion signal
3. You see "running..." then "done" before the program exits
4. The program doesn't exit early

## Hint

You need exactly one line of code that receives from the `done` channel. The receive operation will block until the goroutine sends the completion signal.

```go
// Exercise: Channels synchronisation

// We have a goroutine task, learn how the channel synchronisation works and make the program wait for the asynchronous function (make it synchronous ;) )

package main

import "fmt"
import "time"

func task(done chan bool) {
    fmt.Print("running...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main () {
  var done chan bool = make(chan bool)
  go task(done)

  // Your code goes here

}
```

<details>
<summary> Solution: </summary>

```go
// Exercise: Channels synchronisation

// We have a goroutine task, learn how the channel synchronisation works and make the program wait for the asynchronous function (make it synchronous ;) )

package main

import "fmt"
import "time"

func task(done chan bool) {
    fmt.Print("running...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main () {
  var done chan bool = make(chan bool)
  go task(done)

  // What would happen if we comment out this next line "<- done"?
  // Your code goes here
  <- done
}
```

</details>
