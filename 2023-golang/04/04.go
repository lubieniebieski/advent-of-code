package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/lubieniebieski/advent-of-code/2023-golang/tools"
)

type Card struct {
	ID             int
	Numbers        []int
	WinningNumbers []int
}

func AddCard(input string) (card Card) {
	cardIdRe := regexp.MustCompile(`Card\s+(\d+):`)
	card.ID, _ = strconv.Atoi(cardIdRe.FindStringSubmatch(input)[1])
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
	return card
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

func (c Card) MatchingNumbersCount() int {
	points := 0
	for _, number := range c.WinningNumbers {
		if slices.Contains(c.Numbers, number) {
			points++
		}
	}
	return points
}

func findCopies(strArray []string, maxDepth int) int {
	result := 0

	for i := 0; i < maxDepth; i++ {
		card := AddCard(strArray[i])
		result += 1
		result += findCopies(strArray[i+1:], card.MatchingNumbersCount())
	}
	return result
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
	stringsArray := tools.ExtractStringsFromString(input)

	return findCopies(stringsArray, len(stringsArray))
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
