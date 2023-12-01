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
		`1abc2
		pqr3stu8vwx
		a1b2c3d4e5f
		treb7uchet`,
		142,
	},
	{
		`two1nine
		eightwothree
		abcone2threexyz
		xtwone3four
		4nineeightseven2
		zoneight234
		7pqrstsixteen`,
		281,
	},
}

func TestPartOne(t *testing.T) {
	want := testCases[1].want
	got := PartTwo(testCases[1].input)
	toolstest.CompareWithExample(t, want, got)
}

func TestPartTwo(t *testing.T) {
	want := testCases[2].want
	got := PartTwo(testCases[2].input)
	toolstest.CompareWithExample(t, want, got)
}

// Main tests below

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

func TestGetNumber(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{"numbers at beggining/end", "1abc2", 12},
		{"numbers in the middle", "pqr3stu8vwx", 38},
		{"no numbers", "abc", 0},
		{"one number", "a1b", 11},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			value := CalibrationData{input: c.input}
			got := value.GetNumber()
			if !reflect.DeepEqual(c.want, got) {
				t.Errorf("FindNumbers() = %v, want %v", got, c.want)
			}
		})
	}
}
