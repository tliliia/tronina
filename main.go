package main

import (
	"fmt"
	"math"
)

func main() {
	squares := SolutionSquareGenerator(11, 3)
	fmt.Println(squares)
}

//SquarePow использует стандартную библиотеку
func SquarePow(n float64) int {
	c := math.Pow(n, 2)
	return int(c)
}

//SolutionSquareGenerator is a func
func SolutionSquareGenerator(start int, n int) []int {
	var squaresArray = make([]int, n)
	fmt.Println(squaresArray)
	for i := range squaresArray {
		funcN := float64(start + i)
		squaresArray[i] = int(SquarePow(funcN))
	}
	return squaresArray
}

func SolutionBinaryGap(N int) int {
	return 0
}
