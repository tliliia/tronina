package main

import (
	"fmt"
	"math"
)

func main() {
	squares := SolutionSquareGenerator(11, 3)
	fmt.Println(squares)
}

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

//SolutionBinaryGap is a func
func SolutionBinaryGap(N int) int {
	var maxGap int
	var currentGap int
	for N != 0 {
		lastBit := N % 2
		fmt.Println(lastBit)
		if lastBit == 0 {
			currentGap++
		} else {
			currentGap = 0
		}

		if currentGap > maxGap {
			maxGap = currentGap
		}
		N = int(N / 2)
	}
	return maxGap
}
