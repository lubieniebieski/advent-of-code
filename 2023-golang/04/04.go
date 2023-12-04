package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/lubieniebieski/advent-of-code/2023-golang/tools"
)

type Card struct {
	Numbers        []int
	WinningNumbers []int
}

func AddCard(input string) (card Card) {
	cardIdRe := regexp.MustCompile(`Card\s+\d+:`)

	input = cardIdRe.ReplaceAllString(input, "")
	numbersStr := strings.Split(input, "|")[0]
	winningNumbersStr := strings.Split(input, "|")[1]

	for _, number := range strings.Fields(numbersStr) {
		numberInt, _ := strconv.Atoi(number)
		card.Numbers = append(card.Numbers, numberInt)
	}

	for _, number := range strings.Fields(winningNumbersStr) {
		numberInt, _ := strconv.Atoi(number)
		card.WinningNumbers = append(card.WinningNumbers, numberInt)
	}
	return Card{
		Numbers:        card.Numbers,
		WinningNumbers: card.WinningNumbers,
	}
}

func (c Card) Points() int {
	points := 0
	for _, number := range c.WinningNumbers {
		if slices.Contains(c.Numbers, number) {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}
	return points
}

func PartOne(input string) (result int) {
	stringsArray := tools.ExtractStringsFromString(input)
	for _, data := range stringsArray {
		card := AddCard(data)
		result += card.Points()
	}
	return result
}

func PartTwo(input string) (result int) {

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

	fmt.Printf("Part One: %d \n", PartOne(data))
	fmt.Printf("Part Two: %d \n", PartTwo(data))
}
