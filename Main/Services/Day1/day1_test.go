package Day1_test

import (
	day1 "AdventOfCode/Services/Day1"
	"reflect"
	"testing"
)

func TestCalculateChallenge1(t *testing.T) {
	inputList1 := []int{7, 1, 4, 2, 6}
	inputList2 := []int{2, 5, 3, 1, 8}

	got := day1.CalculateChallenge1(inputList1, inputList2)
	expected := 3

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}

func TestCalculateChallenge2(t *testing.T) {
	inputList1 := []int{7, 1, 4, 2, 6}
	inputList2 := []int{2, 5, 3, 1, 2}

	got := day1.CalculateChallenge2(inputList1, inputList2)
	expected := 5

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}
