package main

import (
	"fmt"
	"os"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
	"golang.org/x/exp/slices"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Priority(char rune) int {
	return slices.Index([]rune(alphabet), char) + 1
}

func Intersection(a, b string) []rune {
	result := make([]rune, 0)

	bRunes := []rune(b)
	for _, v := range a {
		if slices.Index(bRunes, v) > -1 && slices.Index(result, v) == -1 {
			result = append(result, v)
		}
	}
	return result
}

func PartOne(stringsArray []string) int {
	var prioritiesSum int
	for _, v := range stringsArray {
		size := len(v) / 2
		first, last := v[0:size], v[size:]
		commonElements := Intersection(first, last)
		for _, e := range commonElements {
			prioritiesSum += Priority(e)
		}

	}
	return prioritiesSum
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
