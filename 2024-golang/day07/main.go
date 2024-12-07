package day07 // Change this for each new day

import (
	"strings"

	"github.com/lubieniebieski/advent-of-code/2024-golang/utils"
)

type Number struct {
	Value       int
	Ingridients []int
}

func NewNumber(input string) Number {
	split := strings.Split(input, ": ")
	value, _ := utils.GetIntFromString(split[0])
	numbers, _ := utils.GetIntsFromString(split[1])
	return Number{Value: value, Ingridients: numbers}
}

func (n Number) IsTrue() bool {
	results := allPossibleResults(n.Ingridients, n.Value)
	for _, result := range results {
		if result == n.Value {
			return true
		}
	}
	return false
}

func allPossibleResults(numbers []int, max int) (results []int) {

	if len(numbers) == 1 {
		results = numbers
	} else if len(numbers) == 2 {
		add := numbers[0] + numbers[1]
		mul := numbers[0] * numbers[1]
		if add <= max {
			results = append(results, add)
		}
		if mul <= max {
			results = append(results, mul)
		}

	} else {

		withSum := make([]int, 0)
		withSum = append(withSum, numbers[0]+numbers[1])
		withSum = append(withSum, numbers[2:]...)
		results = append(results, allPossibleResults(withSum, max)...)
		withMul := make([]int, 0)
		withMul = append(withMul, numbers[0]*numbers[1])
		withMul = append(withMul, numbers[2:]...)
		results = append(results, allPossibleResults(withMul, max)...)
	}
	return results
}

func Solve1(input []string) int {
	numbers := make([]Number, 0)
	for _, line := range input {
		numbers = append(numbers, NewNumber(line))
	}

	sum := 0
	for _, number := range numbers {
		if number.IsTrue() {
			sum += number.Value
		}
	}
	return sum
}

func Solve2(input []string) int {
	numbers := make([]Number, 0)
	for _, line := range input {
		numbers = append(numbers, NewNumber(line))
	}

	sum := 0
	for _, number := range numbers {
		if number.IsTrue() {
			sum += number.Value
		}
	}
	return sum
}
