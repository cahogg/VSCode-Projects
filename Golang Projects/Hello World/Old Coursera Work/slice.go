package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//Technically, they wanted size three. I created a means of extending theoretically infinitely instead.
	UStore := make([]*int, 0)

	for {
		var input string

		fmt.Println("Please enter an integer to add to the slice, or 'X' to stop:")
		fmt.Scanln(&input)
		/* is the break command to stop. Note that, if you rerun the program, you will
		save the inputs in the slice. Using X to break the loop restores the slice to zero*/
		if input == "X" {
			break
		}
		//Converts string to int
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer or 'X' to stop.")
			continue
		}

		// Creates new variable for adding more integers.
		newNum := num
		// Some value appending
		UStore = append(UStore, &newNum)

		// Sorts the slice. This sucked.
		sort.Slice(UStore, func(i, j int) bool {
			return *UStore[i] < *UStore[j]
		})

		// This prints the updated slice, and seperates each on a new line for readability
		fmt.Println("Updated slice:")
		for _, ptr := range UStore {
			fmt.Println(*ptr)
		}
	}

	// This prints the final sorted slice
	fmt.Print("Final sorted slice: ")
	for _, ptr := range UStore {
		fmt.Print(*ptr)
	}
}
