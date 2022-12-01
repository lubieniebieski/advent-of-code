package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
)

func PartOne(stringsArray []string) int {
	var biggestSum int
	var currentSum int
	for _, v := range stringsArray {
		if v == "" {
			if currentSum > biggestSum {
				biggestSum = currentSum
			}
			currentSum = 0
			continue
		} else {
			var intValue int
			intValue, _ = strconv.Atoi(v)
			currentSum += intValue
		}
	}

	return biggestSum
}

func PartTwo(caloriesList []string) int {
	var sumList []int
	var currentSum int
	for _, v := range caloriesList {
		if v == "" {
			sumList = append(sumList, currentSum)
			currentSum = 0
			continue
		} else {
			var intValue int
			intValue, _ = strconv.Atoi(strings.TrimSpace(v))
			currentSum += intValue
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sumList)))
	var biggestSums int
	for _, v := range sumList[0:3] {
		biggestSums += v
	}
	return biggestSums
}

func parsedData() []string {
	rawData, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return tools.ExtractStringsFromString(string(rawData))
}

func main() {
	data := parsedData()

	fmt.Printf("Part One: %d \n", PartOne(data))
	fmt.Printf("Part Two: %d \n", PartTwo(data))
}
