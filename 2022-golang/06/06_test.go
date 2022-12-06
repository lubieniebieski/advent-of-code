package main

import (
	"strings"
	"testing"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools/toolstest"
)

func TestUniqueCharacters(t *testing.T) {
	t.Run("returns true when all characters are unique", func(t *testing.T) {
		if !UniqueCharacters("abcd") {
			t.Fatal()
		}
	})

	t.Run("returns false when all characters are not unique", func(t *testing.T) {
		if UniqueCharacters("abca") {
			t.Fatal()
		}
	})
}

func TestMarkerPosition(t *testing.T) {
	cases := []struct {
		inputText  string
		want       int
		windowSize int
	}{
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5, 4},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6, 4},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10, 4},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11, 4},
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19, 14},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23, 14},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23, 14},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29, 14},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26, 14},
	}

	for _, test := range cases {
		got := MarkerPosition(test.inputText, test.windowSize)
		name := "Example " + test.inputText[0:3] + "..."
		t.Run(name, func(t *testing.T) {
			if got != test.want {
				t.Errorf("Expected %d, got %d", test.want, got)
			}
		})

	}
}
func TestPartOne(t *testing.T) {
	want := 7
	got := PartOne(testData())
	toolstest.CompareWithExample(t, want, got)
}

func TestPartTwo(t *testing.T) {
	want := 19
	got := PartTwo(testData())
	toolstest.CompareWithExample(t, want, got)
}

func testData() []string {
	return strings.Split(testInput(), "\n")
}

func testInput() string {
	return `mjqjpqmgbljsphdztnvjfqwrcgsmlb`
}
