package day07 // Change this for each new day

import (
	"fmt"
	"strings"

	"github.com/lubieniebieski/advent-of-code/2024-golang/utils"
)

type Number struct {
	Value       int
	Ingridients []int
}

type Version1 struct{}
type Version2 struct{}

type Version interface {
	V2() bool
}

func (Version1) V2() bool {
	return false
}

func (Version2) V2() bool {
	return true
}

func NewNumber(input string) Number {
	split := strings.Split(input, ": ")
	value, _ := utils.GetIntFromString(split[0])
	numbers, _ := utils.GetIntsFromString(split[1])
	return Number{Value: value, Ingridients: numbers}
}

func (n Number) IsTrue(version Version) bool {
	return canReachTarget(version, n.Ingridients, n.Value)

}

func isSuccessfulOperation(version Version, a, b, target int) bool {
	add := a + b
	mul := a * b
	if add == target || mul == target {
		return true
	} else if version.V2() {
		concat := fmt.Sprintf("%d%d", a, b)
		concatInt, _ := utils.GetIntFromString(concat)
		return concatInt == target
	}

	return false
}

func canReachTarget(version Version, numbers []int, target int) bool {
	if len(numbers) == 1 {
		return numbers[0] == target
	}

	if len(numbers) == 2 {
		return isSuccessfulOperation(version, numbers[0], numbers[1], target)
	}

	// Try combining first two numbers with each operation
	add := numbers[0] + numbers[1]
	mul := numbers[0] * numbers[1]

	if canReachTarget(version, append([]int{add}, numbers[2:]...), target) ||
		canReachTarget(version, append([]int{mul}, numbers[2:]...), target) {
		return true
	} else if version.V2() {
		concat := fmt.Sprintf("%d%d", numbers[0], numbers[1])
		concatInt, _ := utils.GetIntFromString(concat)
		return canReachTarget(version, append([]int{concatInt}, numbers[2:]...), target)
	}
	return false

}

func Solve1(input []string) int {
	numbers := make([]Number, 0)
	for _, line := range input {
		numbers = append(numbers, NewNumber(line))
	}

	sum := 0
	for _, number := range numbers {
		if number.IsTrue(Version1{}) {
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
		if number.IsTrue(Version2{}) {
			sum += number.Value
		}
	}
	return sum
}
