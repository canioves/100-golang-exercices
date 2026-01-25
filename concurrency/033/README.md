# Exercise 33: Basic Channels

## What are Channels?

Channels are Go's way of allowing goroutines to communicate and synchronize. They follow Go's concurrency principle: **"Don't communicate by sharing memory; share memory by communicating."**

Think of channels as pipes through which you can send and receive values between goroutines.

## Channel Basics

**Creating channels:**

```go
ch := make(chan string)  // Channel that carries strings
ch := make(chan int)     // Channel that carries integers
```

**Channel operations:**
- **Send**: `ch <- value` (send value into channel)
- **Receive**: `value := <-ch` (receive value from channel)

## Unbuffered Channels (Synchronous)

By default, channels are unbuffered, which means:

- **Sends block** until another goroutine is ready to receive
- **Receives block** until another goroutine is ready to send
- This creates a **synchronization point** between goroutines

## Communication Pattern

The typical pattern is:

1. Create a channel
2. Start one or more goroutines that will send/receive on the channel
3. Use the channel operations to coordinate between goroutines

Instead of guessing how long to wait with `time.Sleep()`, channels provide proper synchronization
because they wait exactly as long as needed, race conditions are avoided and there's a clear communication intent.

## Your Task

Look at the `main.go` file and complete the exercise. You need to:

1. Understand how to send messages through channels
2. Understand how to receive messages from channels
3. Set up communication between two goroutines

## Expected Behavior

When completed, you should see `f2` printing a message that includes the message sent from `f1`. The program will coordinate properly without needing to guess timing.

## Key Points to Remember

- Channels are typed - a `chan string` only carries strings
- Unbuffered channel operations are synchronous
- The arrow `<-` shows the direction of data flow
- Channel operations will block until both sides are ready

```go
// It's important to get this concept, document yourself first! :)
// There are two concurrent routines (f1 and f2)
// Send a message "Hello from f1" from f1 function
// Receive the message from f1 into the f2 function, and print "I am f2 and ..." + the message from f1
// This should be done with a channel

package main

import "fmt"
import "time"

func f1 (c chan string){

}

func f2 (c chan string){

}
func main () {

  // this sleep is in order to not exit the program sooner than the routine lifetime :)
  time.Sleep(1 * time.Second)
}
```

<details>
<summary> Solution: </summary>

```go
// Exercise: Channels

// It's important to get this concept, document yourself first! :)
// There are two concurrent routines (f1 and f2)
// Send a message "Hello from f1" from f1 function

package main

import "fmt"
import "time"

func f1 (c chan string){
  c <- "Hello from f1"
}

func f2 (c chan string){
  msg := <- c
  fmt.Println("I am f2 and...", msg)
}
func main () {
  // Your code goes here
  var c chan string = make(chan string)
  go f1(c)
  go f2(c)

  // this sleep is in order to not exit the program sooner than the routine lifetime :)
  time.Sleep(1 * time.Second)
}
```

</details>
