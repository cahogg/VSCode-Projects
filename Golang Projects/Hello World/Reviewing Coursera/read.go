package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	Fname string
	Lname string
}

func main() {

	fmt.Println("Enter the name of the text file:")
	var fileName string
	fmt.Scanln(&fileName)

	// os.Open allows us to open a file.
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	//Slice for name holding
	var names []Name

	//This reads the lines of the file.
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

	// More error checking
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	//This then just prints those names.
	for _, name := range names {
		fmt.Println("First Name:", name.Fname, "Last Name:", name.Lname)
	}
}
