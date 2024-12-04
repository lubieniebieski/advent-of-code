package day03

import (
	"regexp"
	"strings"

	"github.com/lubieniebieski/advent-of-code/2024-golang/utils"
)

type Mul struct {
	x int
	y int
}

func (m Mul) result() int {
	return m.x * m.y
}

type Muls []Mul

func (ms Muls) sum() int {
	result := 0
	for _, m := range ms {
		result += m.result()
	}
	return result
}

func parseMuls(input string) Muls {
	matcher := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := matcher.FindAllStringSubmatch(input, -1)
	muls := make(Muls, 0, len(matches))
	for _, m := range matches {
		x, _ := utils.GetIntFromString(m[1])
		y, _ := utils.GetIntFromString(m[2])
		muls = append(muls, Mul{x, y})
	}
	return muls
}

func parseEnabledMuls(input string) Muls {
	matcher := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	tokens := matcher.FindAllString(input, -1)
	muls := make(Muls, 0)
	enabled := true

	for _, token := range tokens {
		switch token {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				muls = append(muls, parseMuls(token)...)
			}
		}
	}
	return muls
}

func Solve1(input []string) int {
	joinedInput := strings.Join(input, "")
	return parseMuls(joinedInput).sum()
}

func Solve2(input []string) int {
	joinedInput := strings.Join(input, "")
	return parseEnabledMuls(joinedInput).sum()
}
