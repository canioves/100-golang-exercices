// Exercise: Write a list of 5 countries to a file
// Tip: use the "os" package

package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("file_management/018/create.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	arr := [6]string{"biba", "boba", "buba", "babiba", "gugu", "gaga"}
	for _, e := range arr {
		file.WriteString(e + "\n")
	}
}
