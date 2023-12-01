package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/lubieniebieski/advent-of-code/2023-golang/tools"
)

type CalibrationValue struct {
	input string
}

func (c *CalibrationValue) FindNumbers() (result []int) {
	firstNumberRe := regexp.MustCompile(`(\d).*`)
	firstNumberMatch := firstNumberRe.FindAllStringSubmatch(c.input, -1)
	if len(firstNumberMatch) == 0 {
		return []int{}
	}
	firstNumberInt, _ := strconv.Atoi(firstNumberMatch[0][1])
	result = append(result, firstNumberInt)

	lastNumberRe := regexp.MustCompile(`.*(\d)`)
	lastNumber := lastNumberRe.FindAllStringSubmatch(c.input, -1)[0][1]
	lastNumberInt, _ := strconv.Atoi(lastNumber)
	result = append(result, lastNumberInt)
	return result
}

func (c *CalibrationValue) GetValue() int {
	if len(c.FindNumbers()) < 2 {
		return 0
	}
	s := strconv.Itoa(c.FindNumbers()[0]) + strconv.Itoa(c.FindNumbers()[1])
	number, _ := strconv.Atoi(s)
	return number
}

func PartOne(stringsArray []string) int {
	var result int
	for _, str := range stringsArray {
		calibration := CalibrationValue{input: str}
		result += calibration.GetValue()
	}
	return result
}

func PartTwo(stringsArray []string) int {
	var result int
	for _, str := range stringsArray {
		str = ReplaceTextToNumber(str)
		calibration := CalibrationValue{input: str}
		result += calibration.GetValue()
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

func parsedData() []string {
	rawData, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return tools.ExtractStringsFromString(string(rawData))
}

func main() {
	data := parsedData()

	fmt.Printf("Part One: %d \n", PartOne(data))
	fmt.Printf("Part Two: %d \n", PartTwo(data))
}
