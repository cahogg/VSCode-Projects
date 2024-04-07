package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Name struct represents a person's first and last name
type Name struct {
	Fname string
	Lname string
}

func main() {
	// Prompt the user for the name of the text file
	fmt.Println("Enter the name of the text file:")
	var fileName string
	fmt.Scanln(&fileName)

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a slice to hold the structs
	var names []Name

	// Read lines from the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 2 {
			fmt.Println("Invalid line in file:", line)
			continue
		}
		name := Name{
			Fname: fields[0],
			Lname: fields[1],
		}
		names = append(names, name)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the names from the slice of structs
	for _, name := range names {
		fmt.Println("First Name:", name.Fname, "Last Name:", name.Lname)
	}
}
