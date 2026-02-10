// Exercise: Pointers

// Declare a pointer variable of type int32
// Declare the address of the var "x"
// Save the address of the var "x" in the pointer variable

package main

import "fmt"

func main() {
	var x int32 = 5
	var p *int32
	p = &x
	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Memory address of x: %d\n", &x)
	fmt.Printf("Pointer value: %d\n", p)

}
