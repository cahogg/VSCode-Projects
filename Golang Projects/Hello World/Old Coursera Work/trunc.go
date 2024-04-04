package main

import (
	"fmt"
	"math"
)

var Decimal float64

func main() {

	fmt.Println("Enter a float: ")
	fmt.Scanln(&Decimal)

	Decimal = math.Trunc(Decimal)

	fmt.Printf("%.0f", Decimal)
}
