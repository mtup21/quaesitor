package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {

	r, _ := regexp.Compile("^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{4}|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}={2})$")
	
	file, err := os.Open("test_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	
		if r.MatchString(scanner.Text()) {
			fmt.Println(scanner.Text())
			for i :=0; i < 5; i++{
			   scanner.Scan()
				fmt.Println(scanner.Text())
			}
	
		}
	
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	  }
	}