// Write a program which prompts the user to enter a string.
// The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
// The program should print “Found!” if the entered string starts with the character ‘i’,
// ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise.
// The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

// Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”,
// “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a string: ")
	scanner.Scan()
	myInput := scanner.Text()

	strCompare := strings.ToLower(myInput)

	cond1 := strings.HasPrefix(strCompare, "i")
	cond2 := strings.Contains(strCompare, "a")
	cond3 := strings.HasSuffix(strings.Replace(strCompare, " ", "", -1), "n")

	if cond1 && cond2 && cond3 {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}
}
