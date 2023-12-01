package main

import (
	"reflect"
	"testing"

	"github.com/lubieniebieski/advent-of-code/2023-golang/tools"
	"github.com/lubieniebieski/advent-of-code/2023-golang/tools/toolstest"
)

func TestPartOne(t *testing.T) {
	want := 142
	got := PartOne(testData(testInput()))
	toolstest.CompareWithExample(t, want, got)
}

func TestPartTwo(t *testing.T) {
	want := 281
	got := PartTwo(testData(testInput2()))
	toolstest.CompareWithExample(t, want, got)
}

func TestReplaceTextToNumber(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"one", "aoneb", "ao1eb"},
		{"two", "atwo", "at2o"},
		{"three", "athree", "at3e"},
		{"four", "afour", "af4r"},
		{"five", "afive", "af5e"},
		{"six", "asix", "as6x"},
		{"seven", "aseven", "as7n"},
		{"eight", "aeight89", "ae8t89"},
		{"nine", "anine", "an9e"},
		{"eightwo combo", "eightwo3", "e8t2o3"},
		{"nineight combo", "nineight1", "n9e8t1"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := ReplaceTextToNumber(c.input)
			if c.want != got {
				t.Errorf("replaceTextToNumber() = %v, want %v", got, c.want)
			}
		})
	}
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

func testData(str string) []string {
	return tools.ExtractStringsFromString(str)
}

func testInput() string {
	return `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet
  `
}
func testInput2() string {
	return `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen
  `
}
