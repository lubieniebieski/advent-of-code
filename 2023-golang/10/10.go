package main

import (
	"fmt"
	"os"
	"time"

	"github.com/lubieniebieski/advent-of-code/2023-golang/tools"
)

func PartOne(input string) (result int) {
	pipeMap := ParsePipePoints(input)
	var start PipePoint
	for _, point := range pipeMap {
		for _, pipe := range point {
			if pipe.IsStart() {
				start = pipe
			}
		}
	}
	choices := start.WhereDoIStart(pipeMap)
	path := choices[0].GoToStart(start, pipeMap)

	return len(path)/2 + 1
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
type PipePoint struct {
	x, y int
	id   string
}

func (p PipePoint) IsStart() bool {
	return p.id == "S"
}

func (p PipePoint) WhereDoIStart(points [][]PipePoint) (result []PipePoint) {
	coordinatesToCheck := []struct{ x, y int }{
		{p.x, p.y + 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y},
		{p.x - 1, p.y},
	}

	for _, coordinates := range coordinatesToCheck {
		if coordinates.x >= 0 && coordinates.y >= 0 && coordinates.x < len(points) && coordinates.y < len(points[0]) {

			point := points[coordinates.y][coordinates.x]

			possibleJumps, err := point.PossibleJumps(points)
			if err != nil {
				continue
			}
			for _, jump := range possibleJumps {
				if jump.IsStart() {
					result = append(result, point)
				}
			}
		}
	}

	return result
}

func (p PipePoint) PossibleJumps(pipeMap [][]PipePoint) (result []PipePoint, err error) {
	if p.IsStart() {
		return p.WhereDoIStart(pipeMap), nil
	}
	pairs := [][]int{}
	switch p.id {
	case "|":
		pairs = append(pairs, []int{p.x, p.y + 1}, []int{p.x, p.y - 1})
	case "-":
		pairs = append(pairs, []int{p.x + 1, p.y}, []int{p.x - 1, p.y})
	case "L":
		pairs = append(pairs, []int{p.x + 1, p.y}, []int{p.x, p.y - 1})
	case "J":
		pairs = append(pairs, []int{p.x - 1, p.y}, []int{p.x, p.y - 1})
	case "7":
		pairs = append(pairs, []int{p.x - 1, p.y}, []int{p.x, p.y + 1})
	case "F":
		pairs = append(pairs, []int{p.x + 1, p.y}, []int{p.x, p.y + 1})
	}
	for _, pair := range pairs {
		if pair[0] >= 0 && pair[1] >= 0 && pair[0] < len(pipeMap[0]) && pair[1] < len(pipeMap) {
			result = append(result, pipeMap[pair[1]][pair[0]])
		}
	}
	return result, err
}

func ParsePipePoints(input string) (points [][]PipePoint) {
	stringsArr := tools.ExtractStringsFromString(input)

	for y, str := range stringsArr {
		var row []PipePoint
		for x, char := range str {
			row = append(row, PipePoint{x, y, string(char)})
		}
		points = append(points, row)
	}
	return points
}
func (p PipePoint) MoveFromPoint(origin PipePoint, pipeMap [][]PipePoint) (previous, next PipePoint) {
	possibleJumps, err := p.PossibleJumps(pipeMap)
	if err != nil {
		panic(err)
	}
	for _, jump := range possibleJumps {
		if jump != origin {
			return p, jump
		}
	}
	panic(fmt.Sprintf("No jump from %v to %v", origin, p))
}

func (p PipePoint) GoToStart(start PipePoint, pipeMap [][]PipePoint) (path []PipePoint) {
	origin, next := p.MoveFromPoint(start, pipeMap)
	for !next.IsStart() {
		path = append(path, origin)
		origin, next = next.MoveFromPoint(origin, pipeMap)
	}

	return path
}

func (p PipePoint) String() string {
	return p.id
}
