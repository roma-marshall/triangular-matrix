package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Enter Matrix:")

	// var matrixA [9]int
	matrixA := [9]int{2, -1, 4, 7, 2, 3, 3, -2, 1}
	var request int

	// for i := 0; i < len(matrixA); i++ {
	// 	fmt.Fscan(os.Stdin, &matrixA[i])
	// }

	fmt.Println("=== Matrix A ===")

	for i := 0; i < len(matrixA); i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("")
		}
		if matrixA[i] < 0 {
			fmt.Print(" | ", matrixA[i], " | ")
		} else {
			fmt.Print(" |  ", matrixA[i], " | ")
		}
	}

	fmt.Println("\nThe Matrix was saved!")
	fmt.Println("Enter the number of the choice: ")
	fmt.Println("1 - Get Determinant")
	fmt.Println("2 - Get Rang")
	fmt.Println("3 - something else")

	fmt.Fscan(os.Stdin, &request)

	switch request {
	case 1:
		fmt.Println(findDeterminant(matrixA[:]))
	case 2:
		findRang(matrixA[:])
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("Enter correct number of the choise")
	}

}

func findDeterminant(matrixA []int) string {
	a1 := matrixA[0]
	a2 := matrixA[1]
	a3 := matrixA[2]

	b1 := matrixA[3]
	b2 := matrixA[4]
	b3 := matrixA[5]

	c1 := matrixA[6]
	c2 := matrixA[7]
	c3 := matrixA[8]

	determinant := a1*b2*c3 + a2*c1*b3 + b1*c2*a3 - (a3 * b2 * c1) - (b3 * c2 * a1) - (a2 * b1 * c3)

	result := "Det(A) = "
	result += strconv.Itoa(determinant)

	return result
}

func findRang(matrixA []int) {
	a1 := matrixA[0]
	a2 := matrixA[1]
	a3 := matrixA[2]

	// b1 := matrixA[3]
	// b2 := matrixA[4]
	// b3 := matrixA[5]

	c1 := matrixA[6]
	c2 := matrixA[7]
	c3 := matrixA[8]

	matrixA[0] = (c1 + a1) / 4
	matrixA[1] = (c2 + a2) / 4
	matrixA[2] = (c3 + a3) / 4

	for i := 0; i < len(matrixA); i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("")
		}
		if matrixA[i] < 0 {
			fmt.Print(" | ", matrixA[i], " | ")
		} else {
			fmt.Print(" |  ", matrixA[i], " | ")
		}
	}

}
