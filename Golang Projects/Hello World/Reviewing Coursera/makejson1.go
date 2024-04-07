// Write a program which prompts the user to first
// enter a name, and then enter an address.
// Your program should create a map and add the name
// and address to the map using the keys “name” and
// “address”, respectively. Your program should use
// Marshal() to create a JSON object from the map,
// and then your program should print the JSON object.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	var name, address string
	fmt.Printf("Please enter your Name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	name = scanner.Text()
	fmt.Println("captured name:", name)

	fmt.Printf("Please enter your Address: ")
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	address = scanner.Text()
	fmt.Println("captured address:", address)

	m := make(map[string]string)
	m["name"] = name
	m["address"] = address
	fmt.Println("map", m)

	jsonObj, err := json.Marshal(m)
	if err == nil {
		fmt.Println(string(jsonObj))
	} else {
		fmt.Println(err.Error())
	}
}
