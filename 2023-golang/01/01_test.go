package main

import (
	"reflect"
	"testing"

	"github.com/lubieniebieski/advent-of-code/2023-golang/tools"
	"github.com/lubieniebieski/advent-of-code/2023-golang/tools/toolstest"
)

func TestPartOne(t *testing.T) {
	want := 142
	got := PartOne(testData())
	toolstest.CompareWithExample(t, want, got)
}

func TestPartTwo(t *testing.T) {
	t.Skip("TBD")
	want := 45000
	got := PartTwo(testData())
	toolstest.CompareWithExample(t, want, got)
}

func TestFindNumbers(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  []int
	}{
		{"numbers at beggining/end", "1abc2", []int{1, 2}},
		{"numbers in the middle", "pqr3stu8vwx", []int{3, 8}},
		{"no numbers", "abc", []int{}},
		{"one number", "a1b", []int{1, 1}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			value := CalibrationValue{input: c.input}
			got := value.FindNumbers()
			if !reflect.DeepEqual(c.want, got) {
				t.Errorf("FindNumbers() = %v, want %v", got, c.want)
			}
		})
	}
}

func TestGetValue(t *testing.T) {
	c := CalibrationValue{input: "1abc2"}
	want := 12
	got := c.GetValue()
	if want != got {
		t.Errorf("GetValue() = %v, want %v", got, want)
	}
}

func testData() []string {
	return tools.ExtractStringsFromString(testInput())
}

func testInput() string {
	return `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet
  `
}
