package main

import "fmt"

var Decimal float32
var Whole int

func main() {

	fmt.Println("Enter a float: ")
	fmt.Scanln(&Decimal)

	fmt.Printf("%.0f", Decimal)
}
