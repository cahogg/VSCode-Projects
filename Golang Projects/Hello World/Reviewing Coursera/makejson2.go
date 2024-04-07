package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	m := make(map[string]string)
	var name string
	var address string
	fmt.Print("Enter your Name:")
	fmt.Scan(&name)
	fmt.Print("Enter your Address:")
	fmt.Scan(&address)
	m["name"] = name
	m["address"] = address
	barr, err := json.Marshal(m)
	if err == nil {
		fmt.Print("Marshalled Json:\n")
		os.Stdout.Write(barr)
	}
}
