// Exercise: Read a file

// beware: You should run this code where the read file is, or reference it!
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("file_management/017/read.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
