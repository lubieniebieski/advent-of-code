package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
)

func RoundScore(input string) int {
	s := strings.Split(input, " ")
	move, response := s[0], s[1]
	var points int
	switch response {
	case "X":
		points = 1
		switch move {
		case "A":
			points += 3
		case "B":
			points += 0
		case "C":
			points += 6
		}
	case "Y":
		points = 2
		switch move {
		case "A":
			points += 6
		case "B":
			points += 3
		case "C":
			points += 0
		}
	case "Z":
		points = 3
		switch move {
		case "A":
			points += 0
		case "B":
			points += 6
		case "C":
			points += 3
		}
	}

	return points
}

func RoundScoreV2(input string) int {
	s := strings.Split(input, " ")
	move, expectedResult := s[0], s[1]
	var points int
	switch move {
	case "A":
		switch expectedResult {
		case "X":
			points = 3
		case "Y":
			points = 4
		case "Z":
			points = 8
		}
	case "B":
		switch expectedResult {
		case "X":
			points = 1
		case "Y":
			points = 5
		case "Z":
			points = 9
		}
	case "C":
		points = 1
		switch expectedResult {
		case "X":
			points = 2
		case "Y":
			points = 6
		case "Z":
			points = 7
		}
	}

	return points
}

func PartOne(stringsArray []string) int {
	var total int
	for _, v := range stringsArray {
		total += RoundScore(v)
	}

	return total
}

func PartTwo(stringsArray []string) int {
	var total int
	for _, v := range stringsArray {
		total += RoundScoreV2(v)
	}

	return total
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
