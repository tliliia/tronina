package main

import (
	"fmt"
	"math"
)

//SolutionSquareGenerator is a func that uses math/pow
func SolutionSquareGenerator(start int, n int) []int {
	var squaresArray = make([]int, n)
	fmt.Println(squaresArray)
	for i := range squaresArray {
		floatSquare := math.Pow(float64(start+i), 2)
		squaresArray[i] = int(floatSquare)
	}
	return squaresArray
}
