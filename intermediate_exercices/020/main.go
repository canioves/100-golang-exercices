// Exercise: STRUCT
// Create a Hotel structure with:
// numRooms int32
// streetName string
// hasPool bool

// Then assign a value to each of those attributes

package main

import "fmt"

type Hotel struct {
	numRooms   int32
	streetName string
	hasPool    bool
}

func main() {
	var myHotel Hotel
	myHotel.numRooms = 10
	myHotel.streetName = "biba street"
	myHotel.hasPool = true

	fmt.Printf("%+v", myHotel)
}
