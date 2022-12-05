package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools/toolstest"
)

func TestCratesToString(t *testing.T) {
	t.Run("Works for strings starting with empty crate", func(t *testing.T) {
		want := " A"
		got := CratesToString("    [A]")
		if want != got {
			t.Errorf("Expected to get %v, got %v instead", want, got)
		}
	})
	t.Run("Works for strings containing create in the middle", func(t *testing.T) {
		want := "A B"
		got := CratesToString("[A]     [B]")
		if want != got {
			t.Errorf("Expected to get %v, got %v instead", want, got)
		}
	})
	t.Run("Works for single crane", func(t *testing.T) {
		want := "A"
		got := CratesToString("[A]")
		if want != got {
			t.Errorf("Expected to get %v, got %v instead", want, got)
		}
	})

}

func TestStringsToStacks(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {

		want := []string{
			"CBA",
			"FE",
		}
		got := StringsToStacks([]string{"AE", "BF", "C"})
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Returned value is not correct, it should be %v, got %v", want, got)
		}
	})
	t.Run("example 2", func(t *testing.T) {

		want := []string{
			"ZN",
			"MCD",
			"P",
		}
		got := StringsToStacks([]string{" D", "NC", "ZMP"})
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Returned value is not correct, it should be %v, got %v", want, got)
		}
	})
	t.Run("example 3", func(t *testing.T) {

		want := []string{
			"CBA",
			"HB",
			"G",
			"J",
		}
		got := StringsToStacks([]string{"   J", "ABG", "BH", "C"})
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Returned value is not correct, it should be %v, got %v", want, got)
		}
	})

}
func TestInitialStack(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		want := []string{
			"ZN",
			"MCD",
			"P",
		}
		a := testData()
		got, _ := InitialStack(a)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Returned value is not correct, it should be %v, got %v", want, got)
		}
	})
	t.Run("example 2", func(t *testing.T) {

	})

}

func TestConvertStringToMove(t *testing.T) {
	input := "move 1 from 2 to 1"

	wantHowMany := 1
	wantFrom := 2
	wantTo := 1

	howMany, from, to := ConvertStringToMove(input)
	t.Run("How many is parsed correctly", func(t *testing.T) {
		if howMany != wantHowMany {
			t.Errorf("Expected %v, got %v", wantHowMany, howMany)
		}
	})
	t.Run("From is parsed correctly", func(t *testing.T) {
		if from != wantFrom {
			t.Errorf("Expected %v, got %v", wantFrom, from)
		}
	})
	t.Run("To is parsed correctly", func(t *testing.T) {
		if to != wantTo {
			t.Errorf("Expected %v, got %v", wantTo, to)
		}
	})

}

func TestMoveSubstrings(t *testing.T) {
	input := "move 2 from 2 to 1"
	stacks := []string{"ABC", "DEF"}
	want := []string{"ABCFE", "D"}
	got := MoveSubstrings(stacks, input)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v got %v", want, got)
	}
}

func TestMoveSubstringsV2(t *testing.T) {
	input := "move 2 from 2 to 1"
	stacks := []string{"ABC", "DEF"}
	want := []string{"ABCEF", "D"}
	got := MoveSubstringsV2(stacks, input)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v got %v", want, got)
	}
}
func TestPartOne(t *testing.T) {
	want := "CMZ"
	got := PartOne(testData())
	toolstest.CompareWithExample(t, want, got)
}

func TestPartTwo(t *testing.T) {
	want := "MCD"
	got := PartTwo(testData())
	toolstest.CompareWithExample(t, want, got)
}

func testData() []string {
	return strings.Split(testInput(), "\n")
}

func testInput() string {
	return `    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
}
