// Exercise: Functions

// Create a function that sums 2 numbers, then nest that function in a second function, that will add a third number
// In the main program, only call the secondsum function

package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}

func secondSum(x, y, z int) int {
	return z + sum(x, y)
}

func main() {
	var result int
	// Your code goes here
	result = secondSum(1, 2, 3)
	fmt.Println(result)
}
