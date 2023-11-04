package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Base("./test.txt"))
	migrate("C:\\Users\\jtp03\\OneDrive\\Desktop", "C:\\Users\\jtp03\\OneDrive\\Documents")
}
