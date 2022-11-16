package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

func main() {
	nums := ParseIntegersFromFile("input.txt")
	answer, err := PartOne(nums)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d \n", answer)
}

func ParseIntegersFromFile(fileName string) (result []int) {
	input, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return ParseIntegersFromString(string(input))
}

func ParseIntegersFromString(input string) (ans []int) {
	for _, l := range strings.Split(input, "\n") {

		number, _ := strconv.ParseInt(strings.TrimSpace(l), 0, 32)
		if number == 0 {
			continue
		}
		ans = append(ans, int(number))
	}
	return ans
}
