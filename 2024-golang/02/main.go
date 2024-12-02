package main

import (
	"github.com/lubieniebieski/advent-of-code/2024-golang/utils"
)

func safeRow(row string) bool {
	numbers, _ := utils.GetIntsFromString(row)
	isIncreasing := numbers[0] < numbers[len(numbers)-1]

	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		if isIncreasing {
			if diff <= 0 || diff > 3 {
				return false
			}
		} else {
			if diff >= 0 || diff < -3 {
				return false
			}
		}
	}

	return true
}

func solve1(input []string) int {
	safeCount := 0
	for _, line := range input {
		if safeRow(line) {
			safeCount++
		}
	}

	return safeCount
}

func solve2(input []string) int {
	// TODO: Implement solution for part 2
	return 0
}

func main() {
	utils.Run(2, utils.Solution{Part1: solve1, Part2: solve2})
}
