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
	return r.LastWinningIndex() - r.FirstWinningIndex() + 1
}

func (r Race) FirstWinningIndex() (result int) {
	low := 1
	high := r.Time - 1

	for low <= high {
		mid := (low + high) / 2

		if r.DistanceCoveredWithSpeed(mid) > r.Distance {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}

func (r Race) LastWinningIndex() (result int) {
	low := 1
	high := r.Time - 1

	for low <= high {
		mid := (low + high) / 2

		if r.DistanceCoveredWithSpeed(mid) > r.Distance {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high
}

func (r Race) DistanceCoveredWithSpeed(speed int) int {
	return speed * (r.Time - speed)
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
	lines := strings.Split(input, "\n")
	timeStr := strings.Fields(lines[0])
	distanceStr := strings.Fields(lines[1])
	var resultTimeStr string
	var resultDistanceStr string
	for i := 1; i < len(timeStr); i++ {
		resultTimeStr += timeStr[i]
		resultDistanceStr += distanceStr[i]
	}

	timeInt, _ := strconv.Atoi(resultTimeStr)
	distanceInt, _ := strconv.Atoi(resultDistanceStr)
	race := Race{Time: timeInt, Distance: distanceInt}
	return race.WinningOptionsCount()
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
