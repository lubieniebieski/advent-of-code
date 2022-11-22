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
