// Exercise: Rename a file from name1 to name2
// Tip: use the "os" package

package main

import (
	"os"
)

func main() {
	var src, dest string
	src = "file_management/019/file1.txt"
	dest = "file_management/019/file2.txt"

	os.Rename(src, dest)
}
