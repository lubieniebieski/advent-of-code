package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
	"golang.org/x/exp/slices"
)

func ConvertStringToMove(str string) (howMany, from, to int) {
	var re = regexp.MustCompile(`(?m)move (\d+) from (\d+) to (\d+)`)
	match := re.FindAllStringSubmatch(str, -1)[0]
	howMany, _ = strconv.Atoi(match[1])
	from, _ = strconv.Atoi(match[2])
	to, _ = strconv.Atoi(match[3])
	return howMany, from, to
}

func MoveSubstrings(stacks []string, inputMovement string) []string {
	howMany, from, to := ConvertStringToMove(inputMovement)
	fromStack := stacks[from-1]

	for i := 1; i <= howMany; i++ {
		cuttingPoint := len(fromStack) - 1
		stacks[to-1] += string(fromStack[cuttingPoint])
		fromStack = fromStack[0:cuttingPoint]
	}
	stacks[from-1] = fromStack

	return stacks
}

func MoveSubstringsV2(stacks []string, inputMovement string) []string {
	howMany, from, to := ConvertStringToMove(inputMovement)
	fromStack := stacks[from-1]
	cuttingPoint := len(fromStack) - howMany
	stacks[to-1] += string(fromStack[cuttingPoint:])
	fromStack = fromStack[0:cuttingPoint]

	stacks[from-1] = fromStack

	return stacks
}

func CratesToString(input string) (output string) {
	if len(input) > 4 {
		output += string(input[1])
		output += CratesToString(input[4:])

	} else if len(input) == 3 || len(input) == 4 {
		output += string(input[1])
	}
	return output
}
func StringsToStacks(stringsArray []string) (output []string) {
	for i, j := 0, len(stringsArray)-1; i < j; i, j = i+1, j-1 {
		stringsArray[i], stringsArray[j] = stringsArray[j], stringsArray[i]
	}

	for i := 0; i < len(stringsArray); i++ {
		output = append(output, "")
	}
	for i := 0; i < len(stringsArray); i++ {
		for j := 0; j < len(stringsArray[i]); j++ {
			value := string(stringsArray[i][j])
			if value != " " {
				for len(output) <= j {
					output = append(output, "")
				}
				output[j] += value
			}
		}
	}
	slices.CompactFunc(output, func(a, b string) bool {
		return true
	})
	var tmpOutput []string
	for _, v := range output {
		if v != "" {
			tmpOutput = append(tmpOutput, v)
		}
	}

	return tmpOutput
}
func InitialStack(input []string) (stacksData, remainingStrings []string) {
	cutoffIndex := slices.Index(input, "") - 1
	for _, v := range input[0:cutoffIndex] {
		stacksData = append(stacksData, CratesToString(v))
	}
	return StringsToStacks(stacksData), input[cutoffIndex+2:]
}

func PartOne(stringsArray []string) string {
	stack, moves := InitialStack(stringsArray)
	for _, move := range moves {
		stack = MoveSubstrings(stack, move)
	}
	var result string
	for _, s := range stack {
		result += string(s[len(s)-1])
	}
	return result
}

func PartTwo(stringsArray []string) string {
	stack, moves := InitialStack(stringsArray)
	for _, move := range moves {
		stack = MoveSubstringsV2(stack, move)
	}
	var result string
	for _, s := range stack {
		result += string(s[len(s)-1])
	}
	return result
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

	fmt.Printf("Part One: %v \n", PartOne(data))
	fmt.Printf("Part Two: %v \n", PartTwo(data))
}
