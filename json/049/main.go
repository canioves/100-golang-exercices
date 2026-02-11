// Exercise: JSON!!

// Create a struct "Human", with 2 string attributes: Name and Description

// Find out about the unmarshal function (json decode), and decode the humanJson object into a variable of type Human called "human"
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Human struct {
	Name        string
	Description string
}

func main() {
	humanJson := `{"name": "Rick","description": "has a grandson called Morty"}`
	var human Human
	err := json.Unmarshal([]byte(humanJson), &human)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(human.Name + " is old and " + human.Description)
}
