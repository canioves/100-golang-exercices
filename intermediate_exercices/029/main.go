// Exercise: Functions

// Create a function that returns 2 integer values
// There will be 2 arguments (int)
// The first returned value will be the sum of the arguments, and the second the substraction of them

package main

import "fmt"

func operation(x, y int) (int, int) {
	return x + y, x - y
}

func main() {
	sum, sub := operation(1, 2)
	fmt.Println(sum, sub)
}
