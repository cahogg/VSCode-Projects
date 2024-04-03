// trunc prompts the user to enter a floating point number then prints the
// integer which is a truncated version of the floating point number that was
// entered.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const prompt = "Enter a float, and I will truncate it!  Your float: "
	var uentry string

	// Prompt the user and read their response.
	fmt.Print(prompt)
	if n, err := fmt.Scanln(&uentry); err == nil {
		if n < 1 {
			fmt.Print("You did not enter a number!\n")
			os.Exit(1)
		}
	}
	if fp, err := strconv.ParseFloat(uentry, 32); err == nil {
		// Don't do math, just print the number with 0 precision.
		fmt.Printf("%.f\n", fp)
	} else {
		fmt.Printf("unexpected error %v\n", err)
		os.Exit(1)
	}
}
