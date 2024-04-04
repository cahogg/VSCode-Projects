package main

import (
	"fmt"
	"strings"
)

func main() {
	var ian string

	fmt.Print("Type a word, and I'll tell you if it starts with I, contains an A, and ends in N!\n")

	//User input
	fmt.Scan(&ian)

	/* Sets input string to lowercase */
	ian = strings.ToLower(ian)

	//Checks if FirstLetter is i, if a is somewhere in the middle, and n is at the end
	if ian[0] == 'i' && strings.Contains(ian, "a") && ian[len(ian)-1] == 'n' {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}
}
