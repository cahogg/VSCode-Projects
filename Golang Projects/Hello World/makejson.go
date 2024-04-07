package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// The `json:"name"` stuff is required to let the script know that it will be turned into a json
// Also JSONs need the info from our struct to be capitalize. Name not name, for example.
type UserInfo struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	var User UserInfo

	//NewReader helps prevent putting in a first and last name into the prompt and causing the last name to
	// populate the address field.
	fmt.Println("Please enter your name")
	fullName, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fullName = strings.TrimSpace(fullName)

	/*adding the User.name makes it so that the variable "User" can access the "name" part of the struct.
	similar to accessing certain aspects of nodes in GDscript*/
	User.Name = fullName

	fmt.Println("Please enter your address:")
	fmt.Scanln(&User.Address)

	//err is used to check for errors. If err is not equal to nil, it will print the error associated with it.
	jsonData, err := json.Marshal(User)
	if err != nil {
		fmt.Println("Error Marshalling JSON", err)
	}

	fmt.Println("JSON Data: ", string(jsonData))
}
