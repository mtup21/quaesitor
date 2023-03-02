package main

import "os"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("test_file.txt")

	check(err)

}
