package main

import (
	"fmt"
	"os"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
)

func UniqueCharacters(input string) bool {
	tmp := make(map[rune]bool)
	for _, v := range input {
		if tmp[v] {
			return false
		}
		tmp[v] = true
	}
	return true
}

func MarkerPosition(input string, windowSize int) (result int) {
	for i := 1; i < len(input)-windowSize; i++ {
		start := i - 1
		end := start + windowSize
		windowContents := input[start:end]
		if UniqueCharacters(windowContents) {
			return end
		}
	}
	return -1
}

func PartOne(stringsArray []string) (result int) {
	return MarkerPosition(stringsArray[0], 4)
}

func PartTwo(stringsArray []string) (result int) {
	return MarkerPosition(stringsArray[0], 14)
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

	fmt.Printf("Part One: %v \n", PartOne(data))
	fmt.Printf("Part Two: %v \n", PartTwo(data))
}
