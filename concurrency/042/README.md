# Exercise 42: Range Over Buffered Channels

## Using Range with Channels

The `range` keyword works with channels just like with slices and maps. It's a clean way to receive all values from a channel until it's closed.

## Range Syntax with Channels

```go
for value := range channel {
    fmt.Println(value)
}
```

## Important: Channels Must Be Closed

For `range` to work properly with channels, the channel **must be closed**. Otherwise, the range loop will block forever waiting for more values.

```go
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
close(ch)  // Essential for range to work!

for value := range ch {
    fmt.Println(value)  // Prints 1, 2, 3, then exits
}
```

## What Range Does

1. **Receives values**: Gets each value from the channel
2. **Blocks when empty**: Waits for new values if channel is open
3. **Exits when closed**: Loop ends when channel is closed AND empty
4. **No index**: Unlike slices, only gives you the value, not an index

## Buffered Channels and Range

With buffered channels, you can:
1. Fill the buffer with multiple values
2. Close the channel
3. Range over all buffered values

This is perfect for producer-consumer patterns where you generate all data first, then process it.

## Your Task

Use the `range` keyword to iterate over a buffered channel containing integers. Remember that the channel must be closed before you can range over it.

```go
// Exercise: Channels - Range

// In this exercise we will use the range keyword to iterate over a buffered (async) channel.
// Create a buffered channel (type int) of a dimension of 5
// Put 5 numbers into the channel
// use the 'range' keyword to iterate over the channel elements and print them

// TIP: but buffered channels need to be closed before iterating over them!!! 

package main

import "fmt"


func main () {
	var c chan int = make(chan int,5)

}
```

<details>
<summary> Solution: </summary>

```go
// Exercise: Channels - Range

// In this exercise we will use the range keyword to iterate over a buffered (async) channel.
// Create a buffered channel (type int) of a dimension of 5
// Put 5 numbers into the channel
// use the 'range' keyword to iterate over the channel elements and print them

// TIP: but buffered channels need to be closed before iterating over them!!! 

package main

import "fmt"


func main () {
	var c chan int = make(chan int,5)

	c <- 3
	c <- 6
	c <- 8
	c <- 22
	c <- 1
	close(c)
	
	for element := range c {
		fmt.Println(element)
	}
}
```

</details>