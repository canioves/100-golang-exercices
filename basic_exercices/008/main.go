// Create an array of 5 strings, and initialize it's 2 first values with some random names

package main

import "fmt"

func main() {
	arr := [5]string{"biba", "boba"}
	fmt.Println(arr, len(arr), cap(arr))
}
