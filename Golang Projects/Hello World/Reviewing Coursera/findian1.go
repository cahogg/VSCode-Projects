package main

import (
	"fmt"
	"strings"
)

func main() {
	var inText string
	fmt.Printf("Enter String: ")
	fmt.Scan(&inText)
	maxIdx := len(inText) - 1
	lowerText := strings.ToLower(inText)
	aIdx := strings.Index(lowerText, "a")

	if strings.Index(lowerText, "i") == 0 && strings.Index(lowerText, "n") == maxIdx && aIdx > 0 && aIdx < maxIdx {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
