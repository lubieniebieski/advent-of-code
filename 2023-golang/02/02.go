package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/lubieniebieski/advent-of-code/2023-golang/tools"
)

const MaxRed = 12
const MaxGreen = 13
const MaxBlue = 14

type Game struct {
	input string
	Id    int
}

func NewGame(input string) *Game {
	game := &Game{input: input}
	game.parse()
	return game
}

func (g *Game) parse() {
	idRe := regexp.MustCompile(`Game (\d+):`)
	id, _ := strconv.Atoi(idRe.FindStringSubmatch(g.input)[1])
	g.Id = id
}

func (g *Game) IsValid() bool {
	resultRe := regexp.MustCompile(`(\d+) (\w+)`)
	results := resultRe.FindAllStringSubmatch(g.input, -1)
	for _, result := range results {
		amount, _ := strconv.Atoi(result[1])
		color := result[2]
		if !ValidAmount(amount, color) {
			return false
		}
	}
	return true
}

func (g *Game) Power() int {
	resultRe := regexp.MustCompile(`(\d+) (\w+)`)
	results := resultRe.FindAllStringSubmatch(g.input, -1)
	sums := make(map[string]int)
	for _, result := range results {
		amount, _ := strconv.Atoi(result[1])
		color := result[2]
		if sums[color] < amount {
			sums[color] = amount
		}
	}
	return sums["red"] * sums["green"] * sums["blue"]
}

func ValidAmount(amount int, color string) bool {
	switch color {
	case "red":
		return amount <= MaxRed
	case "green":
		return amount <= MaxGreen
	case "blue":
		return amount <= MaxBlue
	default:
		return false
	}
}

func PartOne(input string) (result int) {
	stringsArray := tools.ExtractStringsFromString(input)
	for _, str := range stringsArray {
		game := NewGame(str)
		if game.IsValid() {
			result += game.Id
		}
	}
	return result
}

func PartTwo(input string) (result int) {
	stringsArray := tools.ExtractStringsFromString(input)
	for _, str := range stringsArray {
		game := NewGame(str)
		result += game.Power()
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

	fmt.Printf("Part One: %d \n", PartOne(data))
	fmt.Printf("Part Two: %d \n", PartTwo(data))
}
