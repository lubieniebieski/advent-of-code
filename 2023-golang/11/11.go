package main

import (
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/lubieniebieski/advent-of-code/2023-golang/tools"
)

func PartOne(input string) (result int) {

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

type Tile struct {
	value string
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
	tiles [][]Tile
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

func GridFromString(input string) (grid Grid) {
	strArray := ModifyStringForGalaxies(input)

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

func ModifyStringForGalaxies(input string) []string {
	strArray := tools.ExtractStringsFromString(input)
	for i := 0; i < len(strArray); i++ {
		galaxyFound := false
		for _, char := range strArray[i] {
			if string(char) != "." {
				galaxyFound = true
				break
			}
		}
		if !galaxyFound {
			strArray = slices.Insert(strArray, i, strArray[i])
			i++
		}
	}

	for x := 0; x < len(strArray[0]); x++ {
		galaxyFound := false
		for y := 0; y < len(strArray); y++ {
			if string(strArray[y][x]) != "." {
				galaxyFound = true
				break
			}
		}
		if !galaxyFound {
			for y := 0; y < len(strArray); y++ {
				strArray[y] = strArray[y][:x] + "." + strArray[y][x:]
			}
			x++
		}
	}

	return strArray
}
