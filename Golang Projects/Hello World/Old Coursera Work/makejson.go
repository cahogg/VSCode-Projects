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

func main() {
	userInfo := make(map[string]string)

	//NewReader helps prevent putting in a first and last name into the prompt and causing the last name to
	// populate the address field.
	fmt.Println("Please enter your name")
	fullName, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fullName = strings.TrimSpace(fullName)

	fmt.Println("Please enter your address:")
	var Address string
	fmt.Scanln(&Address)

	userInfo[fullName] = Address

	//err is used to check for errors. If err is not equal to nil, it will print the error associated with it.
	jsonData, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println("Error Marshalling JSON", err)
	}

	fmt.Println("JSON Data: ", string(jsonData))
}
