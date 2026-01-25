// Exercise: Routines

// Document yourself and ask what a goroutine is.
// Create a go routine named `start()` and call it inside our main() block
// Notice how goroutines do not wait for the execution to finish and follows with the next code statements

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
