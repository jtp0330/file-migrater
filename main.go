package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Base("./test.txt"))
	migrate("SOURCE PATH", "DESTINATION PATH")
}
