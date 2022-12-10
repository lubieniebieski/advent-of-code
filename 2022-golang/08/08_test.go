package main

import (
	"fmt"
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

func TestCalculateTreeScore(t *testing.T) {
	matrix := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}
	tests := []struct {
		i, j, want int
	}{
		{0, 0, 0},
		{4, 0, 0},
		{1, 2, 4},
		{3, 2, 8},
	}
	for _, test := range tests {
		name := fmt.Sprintf("i=%v,j=%v", test.i, test.j)
		t.Run(name, func(t *testing.T) {
			got := CalculateTreeScore(test.i, test.j, matrix)
			if got != test.want {
				t.Errorf("Expected %v, got %v", test.want, got)
			}
		})
	}
}

func TestRotateSquareArray(t *testing.T) {
	arr := [][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}, {4, 4, 4, 4}}
	rotated := RotateIntMatrix(arr)
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

func TestPartTwo(t *testing.T) {
	want := 8
	got := PartTwo(testData())
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
