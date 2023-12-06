package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/lubieniebieski/advent-of-code/2023-golang/tools"
)

type SeedMapType struct {
	Source      int
	Destination int
	Range       int
}

func SeedMapTypeFromString(input string) (smt SeedMapType, err error) {
	seedMapTypeArray := strings.Split(input, " ")

	if len(seedMapTypeArray) != 3 {
		return smt, fmt.Errorf("invalid seedMapTypeArray: %v", seedMapTypeArray)
	}
	smt.Destination, _ = strconv.Atoi(seedMapTypeArray[0])
	smt.Source, _ = strconv.Atoi(seedMapTypeArray[1])
	smt.Range, _ = strconv.Atoi(seedMapTypeArray[2])
	return smt, nil
}

type Seed struct {
	SeedId int
}

func (s Seed) SeedToSomething(target string, seedMapsMap map[string][]SeedMapType) (seed int) {
	transformations := []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	lastIndex := slices.Index(transformations, target)
	seed = s.SeedId
	for i := 1; i <= lastIndex; i++ {
		seed = s.SomethingToSomething(seed, transformations[i-1], transformations[i], seedMapsMap)
	}
	return seed
}

func (s Seed) SomethingToSomething(value int, source string, target string, seedMapsMap map[string][]SeedMapType) int {
	mySlice := seedMapsMap[source+"-to-"+target]

	for _, seedMapType := range mySlice {
		if value >= seedMapType.Source && value < seedMapType.Source+seedMapType.Range {
			return (value - seedMapType.Source) + seedMapType.Destination
		}
	}
	return value
}

func SeedsFromString(input string) (seeds []Seed, seedsMapsMap map[string][]SeedMapType) {
	stringsArray := tools.ExtractStringsFromString(input)
	seedMapsMap := make(map[string][]SeedMapType)
	currentMapping := ""
	if len(stringsArray) == 0 || len(stringsArray[0]) < 7 {
		return seeds, seedMapsMap
	}

	for _, line := range stringsArray[1:] {
		if strings.HasSuffix(line, ":") {
			currentMapping = strings.TrimSuffix(line, " map:")
		} else {
			seedMapType, err := SeedMapTypeFromString(line)
			if err != nil {
				continue
			}
			seedMapsMap[currentMapping] = append(seedMapsMap[currentMapping], seedMapType)
		}
	}

	return SimpleSeedsFromString(stringsArray[0][7:]), seedMapsMap
}

func SimpleSeedsFromString(input string) (seeds []Seed) {
	for _, seed := range strings.Split(input, " ") {
		id, _ := strconv.Atoi(seed)
		seeds = append(seeds, Seed{SeedId: id})
	}
	return seeds
}

func PartOne(input string) (result int) {
	seeds, seedMapsMap := SeedsFromString(input)
	if len(seeds) == 0 {
		return result
	}
	sort.Slice(seeds, func(i, j int) bool {
		return seeds[i].SeedToSomething("location", seedMapsMap) < seeds[j].SeedToSomething("location", seedMapsMap)
	})

	return seeds[0].SeedToSomething("location", seedMapsMap)
}

func PartTwo(input string) (result int) {
	seeds, seedMapsMap := SeedsFromString(input)
	if len(seeds) == 0 {
		return result
	}
	min := 1000000000000000000
	for i := 0; i < len(seeds); i += 2 {
		times := seeds[i+1].SeedId
		maxJ := seeds[i].SeedId + times
		for j := seeds[i].SeedId; j < maxJ; j++ {
			seed := Seed{
				SeedId: j,
			}
			if seed.SeedToSomething("location", seedMapsMap) < min {
				min = seed.SeedToSomething("location", seedMapsMap)
			}
		}
	}

	return min
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
