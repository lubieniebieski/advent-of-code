package main

import (
	"reflect"
	"testing"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
	"github.com/lubieniebieski/advent-of-code/2022-golang/tools/toolstest"
)

func TestIntArrayIntersection(t *testing.T) {
	t.Run("returns common runes between two int arrays if they are present", func(t *testing.T) {
		want := []int{1, 2}
		got := IntArrayIntersection([]int{1, 2, 3}, []int{1, 2})
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Intersection should contain %v but we got %v instead", want, got)
		}
	})
	t.Run("returns empty slice when there are no common runes", func(t *testing.T) {
		got := IntArrayIntersection([]int{1, 2}, []int{3, 4})
		if len(got) > 0 {
			t.Errorf("Intersection should be empty, but it isn't")
		}
	})

	t.Run("returns common rune only once", func(t *testing.T) {
		want := []int{1}
		got := IntArrayIntersection([]int{1, 1, 1, 1}, []int{1, 2, 3})
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Intersection should contain %v but we got %v instead", want, got)
		}
	})
}
func TestRangeToArray(t *testing.T) {
	input := "1-3"
	want := []int{1, 2, 3}
	got := RangeToArray(input)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v, got %v instead", want, got)
	}
}

func TestFullyContains(t *testing.T) {
	t.Run("is true when one array includes the other in fully", func(t *testing.T) {
		firstArray := []int{1, 2, 3, 4}
		secondArray := []int{1}
		if !FullyContains(firstArray, secondArray) {
			t.Fail()
		}

		if !FullyContains(secondArray, firstArray) {
			t.Fail()
		}
	})

	t.Run("is false when one array does not include the other one in fully", func(t *testing.T) {
		firstArray := []int{1, 2, 3, 4}
		secondArray := []int{4, 5}
		if FullyContains(firstArray, secondArray) {
			t.Fail()
		}

		if FullyContains(secondArray, firstArray) {
			t.Fail()
		}
	})
}
func TestPartOne(t *testing.T) {
	want := 2
	got := PartOne(testData())
	toolstest.CompareWithExample(t, want, got)
}

func TestPartTwo(t *testing.T) {
	want := 4
	got := PartTwo(testData())
	toolstest.CompareWithExample(t, want, got)
}

func testData() []string {
	return tools.ExtractStringsFromString(testInput())
}

func testInput() string {
	return `2-4,6-8
	2-3,4-5
	5-7,7-9
	2-8,3-7
	6-6,4-6
	2-6,4-8`
}
