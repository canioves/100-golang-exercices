// Exercise: Random
// Generate a random number from te range [-50, +50]
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Your code goes here
	randInt := rand.Intn(101) - 50
	fmt.Println(randInt)
}
