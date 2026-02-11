// Exercise: Channels

// It's important to get this concept, document yourself first! :)
// There are two concurrent routines (f1 and f2)
// Send a message "Hello from f1" from f1 function
// Receive the message from f1 into the f2 function, and print "I am f2 and ..." + the message from f1
// This should be done with a channel

package main

import (
	"fmt"
	"time"
)

func f1(c chan string) {
	s := "Hello from f1"
	c <- s
}

func f2(c chan string) {
	s1 := <-c
	s2 := "I am f2 and"
	fmt.Println(s2, s1)
}

func main() {
	channel := make(chan string)
	go f1(channel)
	go f2(channel)

	// this sleep is in order to not exit the program sooner than the routine lifetime :)
	time.Sleep(1 * time.Second)
}
