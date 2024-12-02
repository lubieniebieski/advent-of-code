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
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("failed to convert '%s' to integer: %v", s, err)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
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
