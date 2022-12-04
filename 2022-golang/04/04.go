package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
	"golang.org/x/exp/slices"
)

func IntArrayIntersection(a, b []int) []int {
	result := make([]int, 0)

	for _, v := range a {
		if slices.Index(b, v) > -1 && slices.Index(result, v) == -1 {
			result = append(result, v)
		}
	}
	return result
}
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func RangeToArray(input string) []int {
	split := strings.Split(input, "-")
	start, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])

	return makeRange(start, end)
}
func FullyContains(firstArray, secondArray []int) bool {
	min, _ := tools.FindMinAndMax([]int{len(firstArray), len(secondArray)})
	if min == len(IntArrayIntersection(firstArray, secondArray)) {
		return true
	} else {
		return false
	}

}

func PartOne(stringsArray []string) int {
	var sum int
	for _, line := range stringsArray {
		a, b := parseStringToIntArrays(line)

		if FullyContains(a, b) {
			sum += 1
		}
	}
	return sum
}

func PartTwo(stringsArray []string) int {
	var sum int
	for _, line := range stringsArray {
		a, b := parseStringToIntArrays(line)
		if len(IntArrayIntersection(a, b)) > 0 {
			sum += 1
		}
	}
	return sum
}

func parseStringToIntArrays(input string) (first, last []int) {
	split := strings.Split(input, ",")
	first = RangeToArray(split[0])
	last = RangeToArray(split[1])
	return first, last
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
