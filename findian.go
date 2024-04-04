package main

import (
	"fmt"
	"strings"
)

func main() {
	var ian string

	fmt.Print("Type a word, and I'll tell you if it starts with I, A, or N!\n")
	fmt.Scan(&ian)
	fmt.Println(ian)

	FirstLetter := strings.ToLower(ian)[0]

	if FirstLetter == "i" || FirstLetter == "a" || FirstLetter == "n" {
		fmt.Print("Found!", ian, FirstLetter)
	} else {
		fmt.Print("Not Found!")
	}
}
