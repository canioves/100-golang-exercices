# Exercise 32: Introduction to Goroutines

## What are Goroutines?

Goroutines are lightweight threads managed by the Go runtime. They are one of Go's most powerful features for concurrent programming. Unlike operating system threads, goroutines are much more lightweight - you can create thousands or millions of them with minimal overhead.

## Key Characteristics of Goroutines

- **Lightweight**: Start with small stack (around 2KB) that grows as needed
- **Managed**: Go runtime handles scheduling onto OS threads
- **Concurrent**: Multiple goroutines can run simultaneously
- **Easy to create**: Just use the `go` keyword before a function call

## Creating Goroutines

To create a goroutine, simply prefix any function call with the `go` keyword:

```go
go functionName() // This function now runs concurrently
```

Let’s understand the difference between running a normal function and running a function as a goroutine.

### Running a normal function

```sh
statment1
start()
statement2
```

First, `statement1` will be executed.
Then `start()` function will be called.
Once the `start()` function finishes then `statement2` will be executed.

### Running a function as a goroutine

```sh
statment1
go start()
statement2
```

First, `statement1` will be executed.
Then function `start()` will be called as a goroutine which will execute asynchronously.
`statement2` will be executed immediately. It will not wait for `start()` function to complete. The start function will be executed concurrently as a goroutine while the rest of the program continues its execution.

When calling a function as a goroutine, the execution flow will not stop for the function to finish.
The program will continue from the next line while the goroutine will be executed asynchronously in the background.
It's important to note that any return value from the goroutine will be ignored.

## Important Concepts

**Main Goroutine**: Every Go program starts with one goroutine running the `main()` function. When the main goroutine terminates, the entire program terminates, even if other goroutines are still running.

**Scheduling**: The Go runtime automatically manages when and where goroutines run. You don't control which OS thread they run on.

**Why Use time.Sleep()**: In early goroutine examples, you'll often see `time.Sleep()` to give goroutines time to execute before the main function ends. This is a temporary solution - later you'll learn better synchronization methods like channels and WaitGroups.

## Your Task

Look at the `main.go` file and complete the exercise.
Create a function, which will be our goroutine, that prints "In Goroutine".
Then sleep for one second and print a statement before and after our go routine.

The goal is to understand how to:

1. Create goroutines using the `go` keyword
2. Observe concurrent execution
3. Understand why we need to wait for goroutines to complete

## Expected Behavior

When you run your completed program, you should see output from your go routine different from the usual execution flow. The go routine will not wait for the instructions to finish and go will execute immediately the instructions after our go routine.

## Hint

Remember that the `go` keyword turns any function call into a goroutine. The function signature doesn't change - you just add `go` before calling it.

```go
package main

import (
  "fmt"
  "time"
)

func main() {
  // Your go routine goes here

  fmt.Println("Started")
  // Your time.Sleep goes here

  fmt.Println("Finished")
}

func start() {

}

```

<details>
<summary> Solution: </summary>

```go
package main

import (
  "fmt"
  "time"
)

func main() {
  go start()
  fmt.Println("Started")
  time.Sleep(1 * time.Second)
  fmt.Println("Finished")
}

func start() {
  fmt.Println("In Goroutine")
}
```

</details>
