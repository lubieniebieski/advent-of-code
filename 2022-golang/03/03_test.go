package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
	"github.com/lubieniebieski/advent-of-code/2022-golang/tools/toolstest"
)

func TestPriority(t *testing.T) {
	cases := []struct {
		input rune
		want  int
	}{
		{'p', 16},
		{'L', 38},
		{'P', 42},
		{'v', 22},
		{'t', 20},
		{'s', 19},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("Works correctly for '%v' character", test.input), func(t *testing.T) {
			toolstest.CompareWithExample(t, test.want, Priority(test.input))
		})
	}
}

func TestIntersection(t *testing.T) {
	t.Run("returns common runes between two strings if they are present", func(t *testing.T) {
		want := []rune{'a', 'b'}
		got := Intersection("abc", "abd")
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Intersection should contain %v but we got %v instead", want, got)
		}
	})
	t.Run("returns empty slice when there are no common runes", func(t *testing.T) {
		got := Intersection("abc", "def")
		if len(got) > 0 {
			t.Errorf("Intersection should be empty, but it isn't")
		}
	})

	t.Run("returns common rune only once", func(t *testing.T) {
		want := []rune{'a'}
		got := Intersection("aaaaaaabc", "aaaade")
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Intersection should contain %v but we got %v instead", want, got)
		}
	})
}

func TestPartOne(t *testing.T) {
	want := 157
	got := PartOne(testData())
	toolstest.CompareWithExample(t, want, got)
}

func TestPartTwo(t *testing.T) {
	want := 70
	got := PartTwo(testData())
	toolstest.CompareWithExample(t, want, got)
}

func testData() []string {
	return tools.ExtractStringsFromString(testInput())
}

func testInput() string {
	return `vJrwpWtwJgWrhcsFMMfFFhFp
	jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
	PmmdzqPrVvPwwTWBwg
	wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
	ttgJtRGJQctTZtZT
	CrZsJsPPZsGzwwsLwLmpwMDw`
}
