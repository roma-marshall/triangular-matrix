package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// var matrixA [9]int
	// var matrixB [9]int
	matrixA := [9]int{2, -1, 4, 7, 2, 3, 3, -2, 1}
	matrixB := [9]int{2, -1, 4, 7, 2, 3, 3, -2, 1}
	var request int
	var multiplier int

	// fmt.Println("Enter matrixA(3x3):")
	// for i := 0; i < len(matrixA); i++ {
	// 	fmt.Fscan(os.Stdin, &matrixA[i])
	// }

	// fmt.Println("Enter matrixB(3x3):")
	// for i := 0; i < len(matrixB); i++ {
	// 	fmt.Fscan(os.Stdin, &matrixB[i])
	// }

	fmt.Println("      === MatrixA ===")

	for i := 0; i < len(matrixA); i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("")
		}
		if matrixA[i] < 0 {
			fmt.Print(" | ", matrixA[i], "  | ")
		} else {
			fmt.Print(" |  ", matrixA[i], "  | ")
		}
	}

	fmt.Println("\n\n      === MatrixB ===")

	for i := 0; i < len(matrixB); i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("")
		}
		if matrixB[i] < 0 {
			fmt.Print(" | ", matrixB[i], "  | ")
		} else {
			fmt.Print(" |  ", matrixB[i], "  | ")
		}
	}

	fmt.Println("\nThe Matrix was saved!")
	fmt.Println("Enter the number of the choice: ")
	fmt.Println("1 - Get Determinant")
	fmt.Println("2 - Get Sum")
	fmt.Println("3 - Multiply by")

	fmt.Print("\nYour choise: ")
	fmt.Fscan(os.Stdin, &request)

	switch request {
	case 1:
		fmt.Println(findDeterminant(matrixA[:]))
	case 2:
		findSum(matrixA[:], matrixB[:])
	case 3:
		fmt.Print("\nEnter multiplier: ")
		fmt.Fscan(os.Stdin, &multiplier)
		multiplyMatrix(matrixA[:], matrixB[:], multiplier)
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

func findSum(matrixA []int, matrixB []int) {
	matrixR := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < len(matrixR); i++ {
		matrixR[i] = matrixA[i] + matrixB[i]
	}

	for i := 0; i < len(matrixR); i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("")
		}
		if matrixA[i] < 0 {
			fmt.Print(" | ", matrixR[i], "  | ")
		} else {
			fmt.Print(" |  ", matrixR[i], "  | ")
		}
	}
}

func multiplyMatrix(matrixA []int, matrixB []int, multiplier int) {
	for i := 0; i < len(matrixA); i++ {
		matrixA[i] = matrixA[i] * multiplier
		matrixB[i] = matrixB[i] * multiplier
	}

	fmt.Println("      === MULTIPLY A ===")

	for i := 0; i < len(matrixA); i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("")
		}
		if matrixA[i] < 0 {
			fmt.Print(" | ", matrixA[i], "  | ")
		} else {
			fmt.Print(" |  ", matrixA[i], "  | ")
		}
	}

	fmt.Println("\n\n      === MULTIPLY B ===")

	for i := 0; i < len(matrixB); i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("")
		}
		if matrixB[i] < 0 {
			fmt.Print(" | ", matrixB[i], "  | ")
		} else {
			fmt.Print(" |  ", matrixB[i], "  | ")
		}
	}
}
