package main

import (
	"fmt"

	"github.com/lubieniebieski/advent-of-code/2022/tools"
)

func PartOne(nums []int) (int, error) {
	increasedCounter := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			increasedCounter++
		}
	}
	return increasedCounter, nil
}

func PartTwo(nums []int) (int, error) {
	increasedCounter := 0
	i := 2
	previousWindowSum := nums[i-2] + nums[i-1] + nums[i]
	for i := 3; i < len(nums); i++ {

		windowSum := nums[i-2] + nums[i-1] + nums[i]
		if windowSum > previousWindowSum {
			increasedCounter++
		}
		previousWindowSum = windowSum
	}
	return increasedCounter, nil
}

func main() {
	nums := tools.ParseIntegersFromFile("input.txt")
	answer, err := PartOne(nums)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d \n", answer)

	answer, err = PartTwo(nums)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Part Two: %d \n", answer)
}
