package SortHelper_test

import (
	"AdventOfCode/Helpers/SortHelper"
	"reflect"
	"testing"
)

func TestSortList(t *testing.T) {
	inputList := []int{6, 3, 2, 7, 5}

	got := SortHelper.SortList(inputList)
	expected := []int{2, 3, 5, 6, 7}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}
