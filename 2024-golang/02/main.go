package main

import (
	"github.com/lubieniebieski/advent-of-code/2024-golang/utils"
)

func safeRow(row string) bool {
	numbers, _ := utils.GetIntsFromString(row)
	if numbers[0] < numbers[1] {
		for i := 0; i < len(numbers)-1; i++ {
			diff := numbers[i+1] - numbers[i]
			if diff <= 0 || diff > 3 {
				return false
			}
		}
	} else {
		for i := 0; i < len(numbers)-1; i++ {
			diff := numbers[i] - numbers[i+1]
			if diff <= 0 || diff > 3 {
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
