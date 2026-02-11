// Exercise: Tickers

// Create a goroutine with a infinite loop
// The condition for that infinite loop will be a range that goes over the "time.Tick(time.Second * 1)", in other words: a simple ticker
// And at every Tick, we will print "Tick"
// In the main function, call the go routine :)

package main

import (
	"fmt"
	"time"
)

func task() {
	for range time.Tick(time.Second) {
		fmt.Print("Tick! ")
	}
}

func main() {
	go task()
	time.Sleep(time.Second * 5)
}
