package main

import "fmt"

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
