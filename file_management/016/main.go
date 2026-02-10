// Exercise: Check if a file exists
// Tip: use the "os" package

package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("file_management/016/test.txt"); err == nil {
		fmt.Println("exists")
	} else {
		fmt.Println("not exists")
	}
}
