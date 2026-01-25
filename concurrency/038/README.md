# Exercise 38: Channel Directions - Send-Only Channels

## Send-Only Channels

Continuing with channel directions, this exercise focuses on **send-only channels** - channels that can only be used to send values.

## Send-Only Channel Syntax

A send-only channel uses the syntax `chan<- type`:

```go
func sender(ch chan<- string) {  // chan<- = send-only
    ch <- "hello"     // ✓ OK - can send
    // value := <-ch  // ✗ ERROR - cannot receive
}
```

## Memory Aid for Direction

Think of the arrow as showing data flow:
- `chan<- string`: Data flows **into** the channel (send-only)
- `<-chan string`: Data flows **out of** the channel (receive-only)
- `chan string`: Data flows **both ways** (bidirectional)

## Common Pattern: Producer-Consumer

Send-only and receive-only channels are often used together in producer-consumer patterns:

```go
func producer(out chan<- string) {  // Can only send
    out <- "data"
}

func consumer(in <-chan string) {   // Can only receive
    data := <-in
    fmt.Println(data)
}

func main() {
    ch := make(chan string)
    go producer(ch)  // Converts to send-only
    go consumer(ch)  // Converts to receive-only
}
```

## Why Restrict to Send-Only?

1. **Prevents mistakes**: Function cannot accidentally try to receive
2. **Clear responsibility**: This function is clearly a "producer"
3. **API design**: Makes function contracts explicit
4. **Compile-time safety**: Errors caught early, not at runtime

## Your Task

Create a function called `send` that:
1. Takes a send-only channel as its first and only argument
2. Can only send data to the channel
3. Cannot receive data from the channel (compiler will prevent this)

The channel is already created for you in `main()`. You need to:
- Implement the `send` function with proper channel direction
- Use the function to send data to the channel
- Receive the data in main to see it worked

## Expected Behavior

The program should demonstrate that the `send` function can only send to the channel, not receive from it.

## Hint

The channel direction for send-only is `chan<- string`. Remember the arrow points **into** the channel for sending.

```go
// Exercise: Channels directions (only send/tx)

// Make a goroutine with a channel for only send data.
// The function should be called "send" and the send-only channel should be it's 1st and only argument
// Receive data from that channel is prohibited / will cause compiler errors

package main

import "fmt"

func main () {
	var c chan string = make(chan string, 1)
	
}
```

<details>
<summary> Solution: </summary>

```go
// Exercise: Channels directions (only send/tx)

// Make a goroutine with a channel for only send data.
// The function should be called "send" and the send-only channel should be it's 1st and only argument
// Receive data from that channel is prohibited / will cause compiler errors

package main

import "fmt"

func send(c chan <- string){
	c <- "Only send"
}

func main () {
	var c chan string = make(chan string, 1)
	send(c)
	fmt.Println(<-c)
}
```

</details>
