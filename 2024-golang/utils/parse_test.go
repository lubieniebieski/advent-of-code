package utils

import (
	"testing"
)

func TestGetInts(t *testing.T) {
	input := "1 2 3"
	expected := []int{1, 2, 3}

	result, _ := GetIntsFromString(input)
	if len(result) != len(expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("expected %v, got %v", expected, result)
		}
	}

}
