package tools

import (
	"os"
	"strconv"
	"strings"
)

func ParseIntegersFromFile(fileName string) (result []int) {
	input, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return ExtractIntegersFromString(string(input))
}

func ExtractIntegersFromString(input string) (ans []int) {
	for _, l := range strings.Split(input, "\n") {

		number, _ := strconv.Atoi(strings.TrimSpace(l))
		if number == 0 {
			continue
		}
		ans = append(ans, int(number))
	}
	return ans
}

func ExtractStringsFromString(input string) (result []string) {

	for _, l := range strings.Split(input, "\n") {
		result = append(result, strings.TrimSpace(l))
	}
	return result
}

func FindMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func CreateBoolMatrix(size int) [][]bool {
	matrix := make([][]bool, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]bool, size)
	}
	return matrix
}
