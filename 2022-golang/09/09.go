package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
)

type Point struct {
	x, y int
}

type Step struct {
	direction string
	value     int
}

func MakeMove(matrix *[][]bool, h *Point, t *Point, step Step) {
	((*matrix)[t.y][t.x]) = true
	switch step.direction {
	case "R":
		end := h.x + step.value
		for x := h.x; x < end; x++ {
			h.x++
			if h.x-t.x > 1 {
				t.x++
				t.y = h.y
				((*matrix)[t.y][t.x]) = true
			}
		}
	case "L":
		end := h.x - step.value
		for x := h.x; x > end; x-- {
			h.x--
			if t.x-h.x > 1 {
				t.x--
				t.y = h.y
				((*matrix)[t.y][t.x]) = true
			}
		}
	case "U":
		end := h.y - step.value
		for y := h.y; y > end; y-- {
			h.y--
			if t.y-h.y > 1 {
				t.y--
				t.x = h.x
				((*matrix)[t.y][t.x]) = true
			}
		}
	case "D":
		end := h.y + step.value
		for y := h.y; y < end; y++ {
			h.y++
			if h.y-t.y > 1 {
				t.y++
				t.x = h.x
				((*matrix)[t.y][t.x]) = true
			}
		}
	}
}

func CountTrue(grid *[][]bool) (result int) {
	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len(*grid); j++ {
			if (*grid)[i][j] {
				result += 1
			}
		}
	}
	return result
}

func PrintGrid(grid *[][]bool) {
	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len(*grid); j++ {
			if (*grid)[i][j] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n\n")
}

func PartOne(commands []string) (result int) {
	size := 700
	startX := 300
	startY := 600
	grid := tools.CreateBoolMatrix(size)
	headerPosition := Point{x: startX, y: startY}
	tailPosition := Point{x: startX, y: startY}
	for _, cmd := range commands {
		split := strings.Split(cmd, " ")
		value, _ := strconv.Atoi(split[1])

		step := Step{direction: split[0], value: value}
		MakeMove(&grid, &headerPosition, &tailPosition, step)
	}

	return CountTrue(&grid)
}

func PartTwo(stringsArray []string) (result int) {

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
