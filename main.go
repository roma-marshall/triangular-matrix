package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("enter matrix")

	var a1, a2, a3, b1, b2, b3, c1, c2, c3 int
	var request int

	fmt.Fscan(os.Stdin, &a1, &a2, &a3)
	fmt.Fscan(os.Stdin, &b1, &b2, &b3)
	fmt.Fscan(os.Stdin, &c1, &c2, &c3)

	fmt.Println("=== Matrix A ===")
	fmt.Println("( ", a1, " | ", a2, " | ", a3, " )")
	fmt.Println("( ", b1, " | ", b2, " | ", b3, " )")
	fmt.Println("( ", c1, " | ", c2, " | ", c3, " )")

	fmt.Println("The matrix was saved!")
	fmt.Println("Enter the number of the choice: ")
	fmt.Println("1. get determinant")
	fmt.Println("2. something else")
	fmt.Println("3. something else")

	fmt.Fscan(os.Stdin, &request)

	switch request {
	case 1:
		var determinant int = findDeterminant(a1, a2, a3, b1, b2, b3, c1, c2, c3)
		fmt.Println("Determinant: ", determinant)
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("Enter correct number of the choise")
	}

}

func findDeterminant(a1 int, a2 int, a3 int, b1 int, b2 int, b3 int, c1 int, c2 int, c3 int) int {
	var determinant int
	determinant = a1*b2*c3 + a2*c1*b3 + b1*c2*a3 - (a3 * b2 * c1) - (b3 * c2 * a1) - (a2 * b1 * c3)

	return determinant
}
