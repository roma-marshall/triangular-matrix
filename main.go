package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("enter matrix")

	var a1 int
	var a2 int
	var a3 int
	var b1 int
	var b2 int
	var b3 int
	var c1 int
	var c2 int
	var c3 int

	fmt.Fscan(os.Stdin, &a1, &a2, &a3)
	fmt.Fscan(os.Stdin, &b1, &b2, &b3)
	fmt.Fscan(os.Stdin, &c1, &c2, &c3)

	fmt.Println("( ", a1, " | ", a2, " | ", a3, " )")
	fmt.Println("( ", b1, " | ", b2, " | ", b3, " )")
	fmt.Println("( ", c1, " | ", c2, " | ", c3, " )")
}
