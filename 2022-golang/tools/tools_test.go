package tools

import (
	"reflect"
	"testing"
)

func TestFindMinAndMax(t *testing.T) {
	input := []int{-1, 5, -10, 5, 10}
	wantMin := -10
	wantMax := 10
	gotMin, gotMax := FindMinAndMax(input)
	if wantMin != gotMin {
		t.Errorf("Wanted min to be %v, got %v instead", wantMin, gotMin)
	}
	if wantMax != gotMax {
		t.Errorf("Wanted max to be %v, got %v instead", wantMax, gotMax)
	}
}

func TestCreateBoolMatrix(t *testing.T) {
	want := [][]bool{{false, false}, {false, false}}
	got := CreateBoolMatrix(2)

	if len(got) != len(want) {
		t.Errorf("Matrix length is incorrect")
	}
	for i := 0; i < len(want); i++ {
		for j := 0; j < len(want); j++ {
			if want[i][j] != got[i][j] {
				t.Errorf("Expected %v, got %v", want, got)
			}
		}
	}
}

func TestExtractIntegersFromString(t *testing.T) {
	input := `1
  2
  3`
	want := []int{1, 2, 3}
	got := ExtractIntegersFromString(input)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestExtractStringsFromString(t *testing.T) {
	input := `abc
  cba
  bac
  two words`
	want := []string{"abc", "cba", "bac", "two words"}
	got := ExtractStringsFromString(input)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}
