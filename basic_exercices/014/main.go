// Exercise: User input
// Using only the fmt package, ask a user for it's name and then for it's surname
// Store it in 2 variables called "name" and "surname"
// After user has entered the data, print it out

// Tip: https://pkg.go.dev/fmt#hdr-Scanning

package main

import "fmt"

func main() {
	var name, surname string
	fmt.Printf("Name? ")
	fmt.Scan(&name)
	fmt.Printf("Surname? ")
	fmt.Scan(&surname)

	fmt.Println(name, surname)
}
