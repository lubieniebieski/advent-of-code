package day02

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

func saferRow(row string) bool {
	if safeRow(row) {
		return true
	}
	numbers, _ := utils.GetIntsFromString(row)
	for i := 0; i < len(numbers); i++ {
		remaining := make([]int, 0, len(numbers)-1)
		remaining = append(remaining, numbers[:i]...)
		remaining = append(remaining, numbers[i+1:]...)

		remainingStr := utils.IntsToString(remaining)
		if safeRow(remainingStr) {
			return true
		}
	}
	return false
}

func Solve1(input []string) int {
	safeCount := 0
	for _, line := range input {
		if safeRow(line) {
			safeCount++
		}
	}

	return safeCount
}

func Solve2(input []string) int {
	safeCount := 0
	for _, line := range input {
		if saferRow(line) {
			safeCount++
		}
	}

	return safeCount
}
