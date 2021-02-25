package main

import (
	"fmt"
)

func main() {
	squares := SolutionSquareGenerator(11, 3)
	q := SolutionBinaryGap(5)
	fmt.Println(squares, q)
}
