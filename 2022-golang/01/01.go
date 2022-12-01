package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
)

func PartOne(caloriesList []string) (int, error) {
	topElfCaloriesSum := 0
	currentElfCalories := 0
	for i := 0; i < len(caloriesList); i++ {
		if caloriesList[i] == "" {
			if currentElfCalories > topElfCaloriesSum {
			topElfCaloriesSum = currentElfCalories
			}
			currentElfCalories = 0
		continue
		} else {
			calories := 0
			calories, _ = strconv.Atoi(strings.TrimSpace(caloriesList[i]))
			currentElfCalories += calories
		}

	}
	return topElfCaloriesSum, nil
}

func PartTwo(caloriesList []string) (int, error) {
	var topElves []int
	currentElfCalories := 0
	for i := 0; i < len(caloriesList); i++ {
		if caloriesList[i] == "" {
			topElves = append(topElves,currentElfCalories)

			currentElfCalories = 0
		continue
		} else {
			calories := 0
			calories, _ = strconv.Atoi(strings.TrimSpace(caloriesList[i]))
			currentElfCalories += calories
		}

	}

	sort.Sort(sort.Reverse(sort.IntSlice(topElves)))
	topElfCaloriesSum := 0
	for i := 0; i <= 2; i++ {
		topElfCaloriesSum += topElves[i]
	}
	return topElfCaloriesSum, nil
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	nums := tools.ExtractStringsFromString(string(input))
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
