// Exercise: JSON!! - Marshal (encode)

// We will create a json object.
// First, create a struct called 'user' which contains 2 attributes of type string called 'Username' and 'Password'
// Then, create an array of 3 users 'John', 'Anna' and 'Mike' with any random password
// Use the json.Marshal() function to encode the go array into a JSON object.

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Username string
	Password string
}

func main() {
	users := []user{
		{"John", "123456"},
		{"Anna", "42"},
		{"Mike", "33"},
	}

	// use marshal func to convert to json
	out, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}

	// print encoded json data
	fmt.Println(string(out))
}
