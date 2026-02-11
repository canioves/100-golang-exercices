// Exercise: Channels - Range

// In this exercise we will use the range keyword to iterate over a buffered (async) channel.
// Create a buffered channel (type int) of a dimension of 5
// Put 5 numbers into the channel
// use the 'range' keyword to iterate over the channel elements and print them

// TIP: but buffered channels need to be closed before iterating over them!!!

package main

import "fmt"

func fillChannel(c chan<- int, number int) {
	var i int = 1
	for i <= number {
		c <- i
		i++
	}
}

func main() {
	var number int = 5
	var c chan int = make(chan int, number)
	fillChannel(c, number)
	close(c)
	for x := range c {
		fmt.Println(x)
	}

}
