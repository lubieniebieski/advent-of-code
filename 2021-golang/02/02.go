package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lubieniebieski/advent-of-code/2022/tools"
)

type Command struct {
	direction string
	value     int
}

func (c Command) Horizontal() int {
	switch c.direction {
	case "forward":
		return int(c.value)
	case "back":
		return -int(c.value)
	}
	return 0

}

func (c Command) Vertical() int {
	switch c.direction {
	case "up":
		return -int(c.value)
	case "down":
		return int(c.value)
	}
	return 0

}
func PartOne(commands []Command) (int, error) {
	horizontalPosition := 0
	depth := 0
	for _, c := range commands {
		horizontalPosition += c.Horizontal()
		depth += c.Vertical()
	}
	return horizontalPosition * depth, nil
}

func ConvertToCommand(input string) (command Command) {
	split := strings.Split(input, " ")
	value, _ := strconv.Atoi(split[1])
	return Command{split[0], value}
}

func ConvertStringsToCommand(stringsList []string) (commands []Command) {
	for _, s := range stringsList {
		if len(s) == 0 {
			continue
		}
		commands = append(commands, ConvertToCommand(s))
	}
	return commands
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	strings := tools.ExtractStringsFromString(string(input))
	commands := ConvertStringsToCommand(strings)
	answer, err := PartOne(commands)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Part One: %d \n", answer)

}
