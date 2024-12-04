package day01

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func Solve1(input []string) int {

	left, right := createLists(input)

	totalDistance := 0
	for i := 0; i < len(left); i++ {
		totalDistance += int(math.Abs(float64(left[i] - right[i])))
	}

	return totalDistance
}

func createLists(input []string) ([]int, []int) {
	left := []int{}
	right := []int{}

	for _, line := range input {
		parts := strings.Fields(line)
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		left = append(left, num1)
		right = append(right, num2)
	}

	sort.Ints(left)
	sort.Ints(right)
	return left, right
}

func Solve2(input []string) int {
	left, right := createLists(input)

	totalSimilarityScore := 0

	for _, num := range left {
		totalSimilarityScore += num * countSortedInstances(right, num)
	}

	return totalSimilarityScore
}

func countSortedInstances(sorted []int, target int) int {
	// Find the position of the first occurrence
	first := sort.SearchInts(sorted, target)
	// If the number isn't found, return 0
	if first == len(sorted) || sorted[first] != target {
		return 0
	}
	// Find the position of the first larger number
	last := sort.SearchInts(sorted, target+1)
	return last - first
}
