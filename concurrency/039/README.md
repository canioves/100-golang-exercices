# Exercise 39: Select Statement with Multiple Channels

## The Select Statement

The `select` statement is like a `switch` statement for channels. It allows a goroutine to wait on multiple channel operations simultaneously and execute whichever case becomes ready first.

## How Select Works

1. **Evaluation**: All channel expressions are evaluated
2. **Selection**: If multiple cases are ready, one is chosen **randomly**
3. **Blocking**: If no cases are ready, select blocks until one becomes ready
4. **Execution**: The chosen case executes, then select completes

## Basic Select Syntax

```go
select {
case msg1 := <-channel1:
    fmt.Println("Received from channel1:", msg1)
case msg2 := <-channel2:
    fmt.Println("Received from channel2:", msg2)
case channel3 <- "hello":
    fmt.Println("Sent to channel3")
}
```

## Key Properties

- **Non-deterministic**: If multiple channels are ready, Go randomly chooses one
- **Blocking**: Will wait until at least one case can proceed
- **One execution**: Only one case runs per select statement
- **Fair scheduling**: Random selection prevents one fast channel from starving others

## Infinite Select Loop

Common pattern for continuously monitoring multiple channels:

```go
for {
    select {
    case msg := <-ch1:
        handleMessage1(msg)
    case msg := <-ch2:
        handleMessage2(msg)
    }
}
```

## Why Use Select?

- **Multiplexing**: Handle multiple channels in one place
- **Non-blocking patterns**: With default case (not in this exercise)
- **Timeouts**: Using `time.After()` (you'll see this in exercise 40)
- **Fan-in**: Combine multiple input channels into one handler

## Your Task

Complete the exercise to create a system with:
1. Three channels (`c1`, `c2`, `c3`)
2. A `name` function that sends names to a channel in an infinite loop
3. Three goroutines running the `name` function with different names
4. A select statement that receives from all three channels and prints the names

## Expected Behavior

You should see names from all three goroutines appearing in random order, demonstrating that select fairly chooses between ready channels.

## Important Notes

- The goroutines run forever (infinite loops)
- Select will randomly pick among ready channels
- You'll see interleaved output from all three goroutines
- The program runs forever (use Ctrl+C to stop)

```go
// Exercise: Multiple different channels

// We will make use of the 'select' statement
// Create three different channels "c1,c2,c3" for a goroutine called "name"
// The go routine will have 2 arguments, first the channel (type string) and then a name (type string)
// Inside the goroutine (function) will be:
// 1- An infinite loop:
// 1.1- The second argument "name" will be ingested into the channel
// 1.2- Wait for 2 seconds

// In the main program
// call 3 times the 'name' goroutine with a different name each
// make an infinite loop containing the select statement
// Inside that select, put 3 'case' keywords. In each case print the name for each different channel.

package main

import "fmt"
import "time"

func name(c chan string, name string){
	
}

func main () {
	var c1 chan string = make(chan string)
	var c2 chan string = make(chan string)
	var c3 chan string = make(chan string)

	for {
		select {
		
		}
	}
	fmt.Println("Goroutines finished.") // You shouldn't see this message as the goroutines run forever!
}
```

<details>
<summary> Solution: </summary>

```go
// Exercise: Multiple different channels

// We will make use of the 'select' statement
// Create three different channels "c1,c2,c3" for a goroutine called "name"
// The go routine will have 2 arguments, first the channel (type string) and then a name (type string)
// Inside the goroutine (function) will be:
// 1- An infinite loop:
// 1.1- The second argument "name" will be ingested into the channel
// 1.2- Wait for 2 seconds

// In the main program
// Create three different channels "c1,c2,c3" of type string
// call 3 times the 'name' goroutine with a different name each
// make an infinite loop containing the select statement
// Inside that select, put 3 'case' keywords. In each case print the name for each different channel.

package main

import "fmt"
import "time"

func name(c chan string, name string){
	for {
		c <- name
		time.Sleep(2 * time.Second)
	}
}

func main () {
	var c1 chan string = make(chan string)
	var c2 chan string = make(chan string)
	var c3 chan string = make(chan string)

	go name(c1, "John")
	go name(c2, "Mary")
	go name(c3, "Fritz")

	for {
		select {
		case name1 := <-c1:
			fmt.Println(name1)
		case name2 := <-c2:
			fmt.Println(name2)
		case name3 := <-c3:
			fmt.Println(name3)
		}
	}
	fmt.Println("Goroutines finished.") // You shouldn't see this message as the goroutines run forever!
}
```

</details>