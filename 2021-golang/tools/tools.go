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
