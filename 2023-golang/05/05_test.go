package main

import (
	"reflect"
	"testing"

	"github.com/lubieniebieski/advent-of-code/2023-golang/tools/toolstest"
)

// Input for tests from Readme
var testCases = []struct {
	input string
	want  int
}{
	{
		"First is for fun!",
		1337,
	},
	{
		`seeds: 79 14 55 13

		seed-to-soil map:
		50 98 2
		52 50 48
		
		soil-to-fertilizer map:
		0 15 37
		37 52 2
		39 0 15
		
		fertilizer-to-water map:
		49 53 8
		0 11 42
		42 0 7
		57 7 4
		
		water-to-light map:
		88 18 7
		18 25 70
		
		light-to-temperature map:
		45 77 23
		81 45 19
		68 64 13
		
		temperature-to-humidity map:
		0 69 1
		1 0 69
		
		humidity-to-location map:
		60 56 37
		56 93 4`,
		35,
	},
	{
		``,
		46,
	},
}

func TestPartOne(t *testing.T) {
	want := testCases[1].want
	got := PartOne(testCases[1].input)
	toolstest.CompareWithExample(t, want, got)
}
func TestPartTwo(t *testing.T) {
	want := testCases[2].want
	got := PartTwo(testCases[1].input)
	toolstest.CompareWithExample(t, want, got)
}

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartOne(testCases[1].input)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartTwo(testCases[1].input)
	}
}

// Main tests below

func TestSeedsFromString(t *testing.T) {
	t.Run("Seeds are parsed correctly", func(t *testing.T) {
		want := []Seed{
			{SeedId: 79},
			{SeedId: 14},
			{SeedId: 55},
			{SeedId: 13},
		}
		got, _ := SeedsFromString(`seeds: 79 14 55 13`)
		for i := range want {
			if got[i].SeedId != want[i].SeedId {
				t.Errorf("Wrong seed id, got: %d, want: %d.", got[i].SeedId, want[i].SeedId)
			}
		}
	})
	t.Run("SeedsMap SeedToSoil is parsed correctly", func(t *testing.T) {
		want := []SeedMapType{
			{Destination: 50, Source: 98, Range: 2},
			{Destination: 52, Source: 50, Range: 48},
		}
		_, seedsMapsMap := SeedsFromString(testCases[1].input)
		if !reflect.DeepEqual(seedsMapsMap["seed-to-soil"], want) {
			t.Errorf("Wrong seed map type, got: %v, want: %v.", seedsMapsMap["seed-to-soil"], want)
		}
	})

}

func TestSeed_SeedToSomething(t *testing.T) {
	_, seedMapsMap := SeedsFromString(testCases[1].input)

	t.Run("SeedToSomething returns correct values when converting to soil", func(t *testing.T) {
		testCases := []struct {
			seed int
			want int
		}{
			{2, 2},
			{13, 13},
			{14, 14},
			{50, 52},
			{51, 53},
			{55, 57},
			{79, 81},
			{96, 98},
			{97, 99},
			{98, 50},
			{99, 51},
		}
		for _, testCase := range testCases {
			got := Seed{SeedId: testCase.seed}.SeedToSomething("soil", seedMapsMap)
			if got != testCase.want {
				t.Errorf("Wrong soil for seed %d, got: %d, want: %d.", testCase.seed, got, testCase.want)
			}
		}
	})
	t.Run("SeedToSomething returns correct values when converting to fertilizer", func(t *testing.T) {
		testCases := []struct {
			seed int
			want int
		}{
			{13, 52},
			{14, 53},
			{55, 57},
			{79, 81},
		}
		for _, testCase := range testCases {
			got := Seed{SeedId: testCase.seed}.SeedToSomething("fertilizer", seedMapsMap)
			if got != testCase.want {
				t.Errorf("Wrong fertilizer for seed %d, got: %d, want: %d.", testCase.seed, got, testCase.want)
			}
		}
	})
}

func TestSeedMapTypeFromString(t *testing.T) {
	t.Run("SeedMapType is parsed correctly", func(t *testing.T) {
		want := SeedMapType{Destination: 50, Source: 98, Range: 2}
		got, err := SeedMapTypeFromString("50 98 2")
		if got != want || err != nil {
			t.Errorf("Wrong seed map type, got: %v, want: %v.", got, want)
		}
	})
}
