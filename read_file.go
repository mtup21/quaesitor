package main

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.ReadFile("test_file.txt")
	check(err)
	fmt.Print(string(file))
}
