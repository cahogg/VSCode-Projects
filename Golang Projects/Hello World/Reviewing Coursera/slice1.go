package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

func main() {

	var slice = []int{0, 0, 0}
	var count = 0

	for true {
		var inputChar string
		fmt.Printf("Enter an integer: ")
		_, err := fmt.Scan(&inputChar)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		if strings.ToLower(inputChar) == "x" {
			fmt.Println("Bye!")
			return
		}
		x, err := strconv.Atoi(inputChar)
		if err != nil {
			fmt.Println("Not an integer. Try again")
			continue
		}

		if count < 3 {
			for i, v := range slice {
				if v == 0 {
					slice[i] = x
					break
				}
			}
		} else {
			slice = append(slice, x)
		}
		count = count + 1

		sort.Ints(slice)
		fmt.Println("Current slice: ", slice)
	}
}
