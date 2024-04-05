package main

import (
	"fmt"
)

func main() {
	UStore := make([]int, 3)

	//i is used to designate a location for the user input to be placed.
	for i := 0; i < len(UStore); i++ {
		var UInt int

		fmt.Println("Please enter an integer to slice.", i+1)
		fmt.Scanln(&UInt)

		UStore[i] = UInt
	}

	fmt.Println("Your slice!", UStore)
}
