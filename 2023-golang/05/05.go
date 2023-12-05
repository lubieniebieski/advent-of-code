package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

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
	SeedId      int
	seedMapsMap map[string][]SeedMapType
}

func (s Seed) SeedToSomething(seed int, target string) int {
	transformations := []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	lastIndex := slices.Index(transformations, target)
	for i := 1; i <= lastIndex; i++ {
		seed = s.SomethingToSomething(seed, transformations[i-1], transformations[i])
	}
	return seed
}

func (s Seed) SomethingToSomething(value int, source string, target string) int {

	mySlice := s.seedMapsMap[source+"-to-"+target]
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].Destination < mySlice[j].Destination
	})

	for _, seedMapType := range mySlice {
		if value >= seedMapType.Source && value < seedMapType.Source+seedMapType.Range {
			return (value - seedMapType.Source) + seedMapType.Destination
		}
	}
	return value
}

func NewSeedWithSeedMapsMap(seedId int, seedMapsMap map[string][]SeedMapType) Seed {
	return Seed{SeedId: seedId, seedMapsMap: seedMapsMap}
}

func SeedsFromString(input string) (seeds []Seed) {
	stringsArray := tools.ExtractStringsFromString(input)
	seedMapsMap := make(map[string][]SeedMapType)
	currentMapping := ""
	if len(stringsArray) == 0 || len(stringsArray[0]) < 7 {
		return seeds
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

	return SimpleSeedsFromString(stringsArray[0][7:], seedMapsMap)
}

func SimpleSeedsFromString(input string, seedMapsMap map[string][]SeedMapType) (seeds []Seed) {
	for _, seed := range strings.Split(input, " ") {
		id, _ := strconv.Atoi(seed)
		seeds = append(seeds, Seed{SeedId: id, seedMapsMap: seedMapsMap})
	}
	return seeds
}

func PartOne(input string) (result int) {
	seeds := SeedsFromString(input)
	if len(seeds) == 0 {
		return result
	}
	sort.Slice(seeds, func(i, j int) bool {
		return seeds[i].SeedToSomething(seeds[i].SeedId, "location") < seeds[j].SeedToSomething(seeds[j].SeedId, "location")
	})

	return seeds[0].SeedToSomething(seeds[0].SeedId, "location")
}

func PartTwo(input string) (result int) {
	seeds := SeedsFromString(input)
	if len(seeds) == 0 {

		return result
	}
	min := 1000000000000000000
	for i := 0; i < len(seeds); i += 2 {
		times := seeds[i+1].SeedId
		maxJ := seeds[i].SeedId + times
		for j := seeds[i].SeedId; j < maxJ; j++ {
			seed := Seed{
				SeedId:      j,
				seedMapsMap: seeds[i].seedMapsMap,
			}
			if seed.SeedToSomething(j, "location") < min {
				min = seed.SeedToSomething(j, "location")
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

	fmt.Printf("Part One: %d \n", PartOne(data))
	fmt.Printf("Part Two: %d \n", PartTwo(data))
}
