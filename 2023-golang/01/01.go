package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/lubieniebieski/advent-of-code/2023-golang/tools"
)

type CalibrationData struct {
	input string
}

func (c *CalibrationData) GetNumber() (result int) {
	var stringNumber string
	firstNumberRe := regexp.MustCompile(`(\d).*`)
	firstNumberMatch := firstNumberRe.FindAllStringSubmatch(c.input, -1)
	if len(firstNumberMatch) == 0 {
		return 0
	}
	stringNumber = firstNumberMatch[0][1]

	lastNumberRe := regexp.MustCompile(`.*(\d)`)
	lastNumber := lastNumberRe.FindAllStringSubmatch(c.input, -1)[0][1]
	stringNumber += lastNumber
	result, _ = strconv.Atoi(stringNumber)

	return result
}

func PartOne(input string) (result int) {
	stringsArray := tools.ExtractStringsFromString(input)

	for _, str := range stringsArray {
		calibration := CalibrationData{input: str}
		result += calibration.GetNumber()
	}
	return result
}

func PartTwo(input string) (result int) {
	stringsArray := tools.ExtractStringsFromString(input)

	for _, str := range stringsArray {
		str = ReplaceTextToNumber(str)
		calibration := CalibrationData{input: str}
		result += calibration.GetNumber()
	}
	return result
}

func ReplaceTextToNumber(str string) string {
	substitutions := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}
	for k, v := range substitutions {
		str = strings.ReplaceAll(str, k, v)
	}
	return str
}

func parsedData() string {
	rawData, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(rawData)
}

func main() {
	data := parsedData()

	fmt.Printf("Part One: %d \n", PartOne(data))
	fmt.Printf("Part Two: %d \n", PartTwo(data))
}
