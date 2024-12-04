package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// GetIntsFromString converts a space-separated string of numbers into an int slice
func GetIntsFromString(input string) ([]int, error) {
	var numbers []int
	for _, s := range strings.Fields(input) {
		num, err := GetIntFromString(s)
		if err != nil {
			return nil, fmt.Errorf("failed to convert '%s' to integer: %v", s, err)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

// GetIntFromString converts a string into an int
func GetIntFromString(input string) (int, error) {
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("failed to convert '%s' to integer: %v", input, err)
	}
	return num, nil
}

func IntsToString(numbers []int) string {
	str := ""
	for i, n := range numbers {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprintf("%d", n)
	}
	return str
}

// StringsTo2DArray converts a slice of strings into a 2D array of strings
func StringsTo2DArray(input []string) [][]string {
	result := make([][]string, len(input))
	for i, line := range input {
		result[i] = strings.Split(line, "")
	}
	return result
}
