package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"time"

	"github.com/lubieniebieski/advent-of-code/2023-golang/tools"
)

func PartOne(input string) (result int) {
	return Calculate(input, 1)
}

func PartTwo(input string) (result int) {
	return Calculate(input, 1000000)
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

type Tile struct {
	value string
}

type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func (t Tile) IsEmpty() bool {
	return t.value == "."
}

func (t Tile) IsGalaxy() bool {
	return t.value == "#"
}

func (t Tile) String() string {
	return t.value
}

type Grid struct {
	tiles        [][]Tile
	EmptyRows    []int
	EmptyColumns []int
}

func (g Grid) String() string {
	var result string
	for _, row := range g.tiles {
		for _, tile := range row {
			result += tile.value
		}
		result += "\n"
	}
	return result
}

func (g Grid) ShortestPathBetweenGalaxies(source, destination Point, emptyMultiplier int) (result int) {
	multiplier := emptyMultiplier
	thingToSubtract := 0
	if multiplier > 1 {
		thingToSubtract = 1
	}
	emptyColumnsCount := g.CountEmptyColumnsBetweenPoints(source, destination)
	emptyRowsCount := g.CountEmptyRowsBetweenPoints(source, destination)
	x := int(math.Abs(float64(destination.x-source.x))) + emptyColumnsCount*multiplier - emptyColumnsCount*thingToSubtract
	y := int(math.Abs(float64(destination.y-source.y))) + emptyRowsCount*multiplier - emptyRowsCount*thingToSubtract

	return x + y
}

func GridFromString(input string) (grid Grid) {
	strArray := tools.ExtractStringsFromString(input)

	for _, line := range strArray {
		AddStringLineToGrid(line, &grid)
	}

	return grid
}

func AddStringLineToGrid(line string, grid *Grid) {
	var row []Tile
	for _, char := range line {
		tile := Tile{string(char)}
		row = append(row, tile)
	}
	grid.tiles = append(grid.tiles, row)
}

func (g Grid) CountEmptyColumnsBetweenPoints(source, target Point) (result int) {
	if source.x < target.x {
		for x := source.x + 1; x <= target.x; x++ {
			if slices.Contains(g.GetEmptyColumns(), x) {
				result++
			}
		}
	} else {
		for x := source.x - 1; x >= target.x; x-- {
			if slices.Contains(g.GetEmptyColumns(), x) {
				result++
			}
		}
	}

	return result
}

func (g Grid) CountEmptyRowsBetweenPoints(source, target Point) (result int) {
	if source.y < target.y {
		for y := source.y + 1; y <= target.y; y++ {
			if slices.Contains(g.GetEmptyRows(), y) {
				result++
			}
		}
	} else {
		for y := source.y - 1; y >= target.y; y-- {
			if slices.Contains(g.GetEmptyRows(), y) {
				result++
			}
		}
	}

	return result

}

func (g *Grid) GetEmptyRows() []int {
	if len(g.EmptyRows) > 0 {
		return g.EmptyRows
	}
	for y, row := range g.tiles {
		empty := true
		for _, tile := range row {
			if tile.IsGalaxy() {
				empty = false
				break
			}
		}
		if empty {
			g.EmptyRows = append(g.EmptyRows, y)
		}
	}
	return g.EmptyRows
}

func (g *Grid) GetEmptyColumns() []int {
	if len(g.EmptyColumns) > 0 {
		return g.EmptyColumns
	}
	for x, _ := range g.tiles[0] {
		empty := true
		for y := 0; y < len(g.tiles); y++ {
			if g.tiles[y][x].IsGalaxy() {
				empty = false
				break
			}
		}
		if empty {
			g.EmptyColumns = append(g.EmptyColumns, x)
		}
	}
	return g.EmptyColumns
}

func Calculate(input string, multiplier int) (result int) {
	grid := GridFromString(input)
	galaxyCoordinates := []Point{}
	for y, row := range grid.tiles {
		for x, tile := range row {
			if tile.IsGalaxy() {
				galaxyCoordinates = append(galaxyCoordinates, Point{x, y})
			}
		}
	}

	pairs := [][]Point{}
	for i := 0; i < len(galaxyCoordinates); i++ {
		for j := i + 1; j < len(galaxyCoordinates); j++ {
			pair := []Point{galaxyCoordinates[i], galaxyCoordinates[j]}
			pairs = append(pairs, pair)
		}
	}

	for _, pair := range pairs {
		length := grid.ShortestPathBetweenGalaxies(pair[0], pair[1], multiplier)
		if length > 0 {
			result += length
		}
	}

	return result
}
