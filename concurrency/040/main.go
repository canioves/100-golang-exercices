// Exercise: Channels and timeouts

// Create a go routine that uses time.Sleep for 10 seconds, and then add the string "10 seconds passed" into the channel (of type string)
// in the main program, inside the select block, have 2 cases:
// 1- The message from the channel
// 2- A timeout with the "time.After(3 * time.Second)" statement. After timeout happens, print "Timeout!!!!"

package main

import (
	"fmt"
	"time"
)

func timeout(c chan string) {
	for {
		time.Sleep(time.Second * 10)
		c <- "10 seconds passed"
	}
}

func main() {
	var c1 chan string = make(chan string)

	go timeout(c1)
	for {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case x := <-time.After(1 * time.Second):
			fmt.Printf("timeout after %d seconds\n", x.Second())
		}
	}
	fmt.Println("Goroutines finished.") // You shouldn't see this message as the goroutines run forever!
}
