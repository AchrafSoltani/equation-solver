package main

import (
	"fmt"

	"achrafsoltani.com/equation-solver/lib"
)

func main() {
	var op string
	fmt.Println("Welcome to the equation solver 1.0")
	fmt.Println("Please select an operation:")
	fmt.Println("(1) Linear equations")
	fmt.Println("(2) Quadratic equations")
	fmt.Println("(3) Cubic equations")
	fmt.Println("(4) Bi-quadratic equations")
	fmt.Print(">> ")
	fmt.Scanf("%s", &op)

	switch op {
	case "1":
		lib.Monomial()
	case "2":
		lib.Quadratic()
	case "3":
		lib.Cubic()
	default:
		fmt.Println("Please select one of the possible operations")
	}

}
