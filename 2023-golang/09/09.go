package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/lubieniebieski/advent-of-code/2023-golang/tools"
)

func PartOne(input string) (result int) {
	strArray := tools.ExtractStringsFromString(input)
	for _, line := range strArray {
		result += LineValuesFromString(line).FindNext()
	}
	return result
}

func PartTwo(input string) (result int) {
	strArray := tools.ExtractStringsFromString(input)
	for _, line := range strArray {
		result += LineValuesFromString(line).FindPrevious()
	}
	return result
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
	start := time.Now()
	fmt.Printf("Part One: %d \n", PartOne(data))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Printf("Part Two: %d \n", PartTwo(data))
	fmt.Println(time.Since(start))
}

// Main code

type LineValues struct {
	Values []int
}

func (l *LineValues) Add(val int) {
	l.Values = append(l.Values, val)
}

func (l LineValues) AllEqual() bool {
	for i, val := range l.Values {
		if i == 0 {
			continue
		}
		if val != l.Values[i-1] {
			return false
		}
	}
	return true
}

func (l LineValues) LastDifference() (result int) {
	return l.Values[len(l.Values)-1] - l.Values[len(l.Values)-2]
}

func (l LineValues) FirstDifference() (result int) {
	return l.Values[1] - l.Values[0]
}

func (l LineValues) FindNext() (result int) {
	diffs := LineValues{}
	for i, val := range l.Values {
		if i == 0 {
			continue
		}
		diffs.Add(val - l.Values[i-1])
	}

	if diffs.AllEqual() {
		return l.Values[len(l.Values)-1] + l.LastDifference()
	} else {
		return l.Values[len(l.Values)-1] + diffs.FindNext()
	}
}

func (l LineValues) FindPrevious() (result int) {
	diffs := LineValues{}
	for i, val := range l.Values {
		if i == 0 {
			continue
		}
		diffs.Add(val - l.Values[i-1])
	}

	if diffs.AllEqual() {
		return l.Values[0] - l.FirstDifference()
	} else {
		return l.Values[0] - diffs.FindPrevious()
	}
}

func LineValuesFromString(input string) (result LineValues) {
	for _, val := range strings.Split(input, " ") {
		val, _ := strconv.Atoi(val)
		result.Values = append(result.Values, val)
	}
	return result
}
