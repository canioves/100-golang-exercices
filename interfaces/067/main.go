// Interfaces -  Empty interfaces, bear in mind...
package main

import (
	"fmt"
	"log"
)

type human map[string]interface{}

// But there's an important thing to point out when it comes to retrieving and using a value from this map
// let's say that we want to get the "age" value and increment it by 1.
func main() {
	person := make(human)
	person["name"] = "Alice"
	person["age"] = 21
	person["height"] = 167.64

	// Type assertion for the integer. use .(int) for type assertion!
	age, ok := person["age"].(string)
	// Check that the assertion was alright
	if !ok {
		log.Fatalln("cant add to not integer")
	}

	person["age"] = age + "123"

	fmt.Printf("%+v", person)
}
