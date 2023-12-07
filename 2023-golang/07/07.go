package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/lubieniebieski/advent-of-code/2023-golang/tools"
)

func PartOne(input string) (result int) {

	hands := SortHandsByStrength(HandsFromString(input))
	for i, hand := range hands {
		result += hand.Rank * (i + 1)
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
	start := time.Now()
	fmt.Printf("Part One: %d \n", PartOne(data))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Printf("Part Two: %d \n", PartTwo(data))
	fmt.Println(time.Since(start))
}

// Main code

const (
	HighCard     = 0
	OnePair      = 1
	TwoPairs     = 2
	ThreeOfAKind = 3
	FullHouse    = 4
	FourOfAKind  = 5
	FiveOfAKind  = 6
)

type Hand struct {
	Rank               int
	StrValue           string
	CalculatedStrength int
}

var cardValues = map[string]int{
	"2":  0,
	"3":  1,
	"4":  2,
	"5":  3,
	"6":  4,
	"7":  5,
	"8":  6,
	"9":  7,
	"10": 8,
	"J":  9,
	"Q":  10,
	"K":  11,
	"A":  12,
}

func HandFromString(str string) (hand Hand) {
	values := strings.Split(str, " ")
	rank, _ := strconv.Atoi(values[1])
	hand.Rank = int(rank)
	hand.StrValue = values[0]
	hand.CalculatedStrength = hand.Strength()
	return hand
}

func (h Hand) Strength() (result int) {

	cards := strings.Split(h.StrValue, "")
	cardCount := make(map[string]int)

	for _, card := range cards {
		cardCount[card]++
	}

	switch len(cardCount) {
	case 1:
		return FiveOfAKind
	case 2:
		for _, count := range cardCount {
			switch count {
			case 4:
				return FourOfAKind
			case 3:
				return FullHouse
			}
		}
	case 3:
		for _, count := range cardCount {
			switch count {
			case 3:
				return ThreeOfAKind
			case 2:
				return TwoPairs
			}
		}
	case 4:
		return OnePair
	}
	return result
}

func SortHandsByStrength(hands []Hand) (result []Hand) {
	sort.Slice(hands, func(a, b int) bool {
		if hands[a].CalculatedStrength < hands[b].CalculatedStrength {
			return true
		} else if hands[a].CalculatedStrength > hands[b].CalculatedStrength {
			return false
		} else {
			return !IsFirstHandHigher(hands[a], hands[b])
		}
	})
	return hands
}

func IsFirstHandHigher(a, b Hand) bool {
	for i := 0; i < len(a.StrValue); i++ {
		if cardValues[string(a.StrValue[i])] > cardValues[string(b.StrValue[i])] {
			return true
		} else if cardValues[string(a.StrValue[i])] < cardValues[string(b.StrValue[i])] {
			return false
		}
	}
	return false
}

func HandsFromString(str string) (hands []Hand) {
	stringsArray := tools.ExtractStringsFromString(str)
	for _, line := range stringsArray {
		hands = append(hands, HandFromString(line))
	}
	return hands
}
