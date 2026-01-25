# Exercise 36: Sequential Execution with Channels

## Controlling Execution Order

Sometimes you need to run goroutines in a specific sequence.
While goroutines normally run concurrently (at the same time), channels allow you to coordinate and control their execution order.

## Sequential vs Concurrent Execution

- **Concurrent**: Multiple goroutines running at the same time
- **Sequential**: One goroutine runs, then the next, then the next...

## Using Channels for Sequencing

You can use channel synchronization to ensure one goroutine completes before starting the next:

```go
// Pattern for sequential execution
done := make(chan bool)
go task1(done)
<-done  // Wait for task1 to complete

go task2()  // Now start task2
```

## Buffered Channel Note

Notice the exercise uses a buffered channel with capacity 1:
```go
make(chan bool, 1)
```

This prevents the goroutine from blocking when sending the completion signal, even if the main function isn't immediately ready to receive.

## Why Sequential Execution?

Sometimes you need sequential execution because:
- **Dependencies**: Task2 depends on Task1's results
- **Resource constraints**: Only one task can use a resource at a time
- **Ordered output**: Results must appear in specific order

## Your Task

You need to:
1. Start the first goroutine (`task`)
2. Wait for it to complete
3. Only then start the second goroutine (`task2`)

## Expected Behavior

The output should show:
1. "I am running in the main thread concurrently"
2. "running Task1 goroutine..."
3. "done"
4. "Task2 goroutine..."

Task2 should only run AFTER Task1 completes.

## Hint

You'll need to use channel operations to coordinate between the main function and the goroutines. Remember that receiving from a channel blocks until a value is available.

```go
// There are two goroutines now
// Call the first goroutine (named 'task') and run the second goroutine 'task2' ONLY after the first one has finished

package main

import "fmt"
import "time"

func task(done chan bool) {
  fmt.Print("running Task1 goroutine...")
  time.Sleep(time.Second)
  fmt.Println("done")
  done <- true
}

func task2() {
  fmt.Println("Task2 goroutine...")
}

func main() {
  var done chan bool = make(chan bool, 1)
  fmt.Println("I am running in the main thread concurrently")
  // Your code goes here

}


```

<details>
<summary> Solution: </summary>

```go
// Exercise: Channels synchronisation

// There are two goroutines now
// Call the first goroutine (named 'task') and run the second goroutine 'task2' ONLY after the first one has finished

package main

import "fmt"
import "time"

func task(done chan bool) {
  fmt.Print("running Task1 goroutine...")
  time.Sleep(time.Second)
  fmt.Println("done")
  done <- true
}

func task2() {
  fmt.Println("Task2 goroutine...")
}

func main() {
  var done chan bool = make(chan bool, 1)
  fmt.Println("I am running in the main thread concurrently")
  // Your code goes here
  go task(done)
  if <-done {
    go task2()
    time.Sleep(time.Second)
  }
}
```

</details>
