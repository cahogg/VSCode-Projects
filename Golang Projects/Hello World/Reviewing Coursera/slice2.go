package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	numbers := make([]int, 3)
	var typedNum string

	for {
		fmt.Println("Enter an integer number or X to exit: ")
		_, err := fmt.Scan(&typedNum)
		if err != nil {
			fmt.Println("there was an error")
			continue
		}
		if typedNum == "X" {
			break
		}
		if value, err := strconv.Atoi(typedNum); err == nil {
			numbers = append(numbers, value)
			sort.Ints(numbers)
			fmt.Println(numbers)
		}
	}

}
