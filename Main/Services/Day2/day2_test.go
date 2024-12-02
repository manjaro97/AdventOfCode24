package Day2_test

import (
	"AdventOfCode/Services/Day2"
	"reflect"
	"testing"
)

func TestCalculateChallenge1(t *testing.T) {
	inputList := [][]int{{7, 1, 4, 2, 6}, {1, 2, 3, 4, 5}, {6, 3, 2, 1}, {6, 3, 3, 1}, {1, 6, 7, 8, 9}}

	got := Day2.CalculateChallenge1(inputList)
	expected := 2

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}

func TestCalculateChallenge2(t *testing.T) {
	inputList := [][]int{{7, 1, 4, 2, 6}, {1, 2, 3, 4, 5}, {6, 3, 2, 1}, {6, 3, 3, 1}, {1, 6, 7, 8, 9}}

	got := Day2.CalculateChallenge2(inputList)
	expected := 4

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}
