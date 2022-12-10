package main

import (
	"reflect"
	"testing"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
	"github.com/lubieniebieski/advent-of-code/2022-golang/tools/toolstest"
)

func TestToIntMatrix(t *testing.T) {
	input := []string{"123", "456", "789"}
	want := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	got := ToIntMatrix(input)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestNthCharToInt(t *testing.T) {
	str := "0123456789"
	cases := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, n := range cases {
		t.Run("case", func(t *testing.T) {
			if result := NthCharToInt(n, str); result != n {
				t.Errorf("Expected %v to be %v", n, result)
			}
		})

	}

}

func TestRotateSquareArray(t *testing.T) {
	// Declare a square 2D array of integers
	arr := [][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}, {4, 4, 4, 4}}

	// Rotate the array
	rotated := RotateIntMatrix(arr)

	// Check if the rotated array is as expected
	expected := [][]int{{4, 3, 2, 1}, {4, 3, 2, 1}, {4, 3, 2, 1}, {4, 3, 2, 1}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if rotated[i][j] != expected[i][j] {
				t.Errorf("Test failed: expected rotated array to be %v but got %v", expected, rotated)
			}
		}
	}
}
func TestPartOne(t *testing.T) {
	want := 21
	got := PartOne(testData())
	toolstest.CompareWithExample(t, want, got)
}

func testData() []string {
	return tools.ExtractStringsFromString(testInput())
}

func testInput() string {
	return `30373
25512
65332
33549
35390`
}
