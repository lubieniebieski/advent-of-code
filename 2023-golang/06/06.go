package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Race struct {
	Time     int
	Distance int
}

func (r Race) WinningOptionsCount() (result int) {
	for speed := 1; speed < r.Time; speed++ {
		distance := speed * (r.Time - speed)
		if distance > r.Distance {
			result++
		}
	}
	return result
}

func (r Race) HowFarWithSpeed(speed, distance int) int {
	return speed * distance
}

func ParseRaces(input string) (result []Race) {
	lines := strings.Split(input, "\n")
	timeStr := strings.Fields(lines[0])
	distanceStr := strings.Fields(lines[1])

	for i := 1; i < len(timeStr); i++ {
		timeInt, _ := strconv.Atoi(timeStr[i])
		distanceInt, _ := strconv.Atoi(distanceStr[i])
		result = append(result, Race{Time: timeInt, Distance: distanceInt})
	}
	return result
}

func PartOne(input string) (result int) {
	races := ParseRaces(input)
	result = 1
	for _, race := range races {
		result *= race.WinningOptionsCount()
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
