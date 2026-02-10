// Exercise: MAP
// Create a map of ints to strings
// 1 should resolve A
// 2 should resolve B
// 3 should resolve C

package main

import "fmt"

func main() {
	var mapp map[int]string = map[int]string{4: "A", 2: "B", 3: "C"}
	for k, v := range mapp {
		fmt.Println(k, v)
	}

}
