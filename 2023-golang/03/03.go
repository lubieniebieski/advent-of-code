package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/lubieniebieski/advent-of-code/2023-golang/tools"
)

type Part struct {
	PartNumber     int
	x, y           int
	AdjecentSymbol Point
}

type Point struct {
	x, y int
}

func (p *Part) Length() int {
	return len(strconv.Itoa(p.PartNumber))
}

func (p *Part) AdjecentToCoordinates() []Point {
	coordinates := make(map[Point]bool)

	for i := 0; i < p.Length(); i++ {
		coordinates[Point{p.x + 1 + i, p.y}] = true
		coordinates[Point{p.x - 1 + i, p.y}] = true
		coordinates[Point{p.x + i, p.y + 1}] = true
		coordinates[Point{p.x + i, p.y - 1}] = true
		coordinates[Point{p.x + 1 + i, p.y + 1}] = true
		coordinates[Point{p.x + 1 + i, p.y - 1}] = true
		coordinates[Point{p.x - 1 + i, p.y + 1}] = true
		coordinates[Point{p.x - 1 + i, p.y - 1}] = true
	}

	result := make([]Point, 0, len(coordinates))
	for coordinate := range coordinates {
		result = append(result, coordinate)
	}

	return result
}

func (p *Part) AdjecentToSymbol(data []string, symbols []string) bool {
	for _, point := range p.AdjecentToCoordinates() {
		x, y := point.x, point.y
		if x < 0 || y < 0 {
			continue
		}
		if x >= len(data[0]) || y >= len(data) {
			continue
		}
		if IsSymbol(string(data[y][x]), symbols) {
			p.AdjecentSymbol = Point{x, y}
			return true
		}
	}
	return false
}

func (p *Part) PointsAdjecent(data []string, symbols []string) []Point {
	points := make([]Point, 0)
	for _, point := range p.AdjecentToCoordinates() {
		x, y := point.x, point.y
		if x < 0 || y < 0 {
			continue
		}
		if x >= len(data[0]) || y >= len(data) {
			continue
		}
		if IsSymbol(string(data[y][x]), symbols) {
			points = append(points, point)
		}
	}
	return points
}

func IsSymbol(str string, symbols []string) bool {
	for _, symbol := range symbols {
		if strings.Contains(str, symbol) {
			return true
		}
	}
	return false
}

func FindSymbols(data []string) []string {
	symbols := []string{}
	str := strings.Join(data, "")
	nonSymbolsRe := regexp.MustCompile(`(\d+)|\.`)
	nonSymbols := nonSymbolsRe.ReplaceAllString(str, "")
	uniques := make(map[string]bool)
	for _, symbol := range nonSymbols {
		uniques[string(symbol)] = true
	}
	for symbol := range uniques {
		symbols = append(symbols, symbol)
	}

	return symbols
}

func FindParts(input string, coordinateY int) (parts []Part) {
	partRe := regexp.MustCompile(`\d+`)
	for _, match := range partRe.FindAllStringSubmatch(input, -1) {
		partNumber, _ := strconv.Atoi(match[0])
		parts = append(parts, Part{PartNumber: partNumber})
	}

	for i, match := range partRe.FindAllStringSubmatchIndex(input, -1) {
		parts[i].x = match[0]
		parts[i].y = coordinateY
	}

	return parts
}

func PartOne(input string) (result int) {
	stringsArray := tools.ExtractStringsFromString(input)
	symbols := FindSymbols(stringsArray)
	for y, str := range stringsArray {
		parts := FindParts(str, y)
		for _, part := range parts {
			if part.AdjecentToSymbol(stringsArray, symbols) {
				result += part.PartNumber
			}
		}
	}
	return result
}

func PartTwo(input string) (result int) {
	stringsArray := tools.ExtractStringsFromString(input)
	asterisks := make(map[Point]Part)
	for y, str := range stringsArray {
		parts := FindParts(str, y)
		for _, part := range parts {
			if part.AdjecentToSymbol(stringsArray, []string{"*"}) {
				if asterisks[part.AdjecentSymbol] != (Part{}) {
					result += part.PartNumber * asterisks[part.AdjecentSymbol].PartNumber
				} else {
					asterisks[part.AdjecentSymbol] = part
				}
			}
		}
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
	start := time.Now()
	fmt.Printf("Part One: %d \n", PartOne(data))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Printf("Part Two: %d \n", PartTwo(data))
	fmt.Println(time.Since(start))
}
