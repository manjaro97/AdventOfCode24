package ListHelper_test

import (
	"AdventOfCode/Helpers/ListHelper"
	"reflect"
	"testing"
)

func TestGetListSum(t *testing.T) {
	inputList := []int{7, 1, 4, 2, 6}

	got := ListHelper.GetListSum(inputList)
	expected := 20

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}

func TestGetListDiff(t *testing.T) {
	inputList1 := []int{7, 1, 4, 2, 6}
	inputList2 := []int{2, 3, 4, 5, 6}

	got := ListHelper.GetListDiff(inputList1, inputList2)
	expected := []int{5, 2, 0, 3, 0}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}

func TestGetTotalOccurenceInList(t *testing.T) {
	inputValue := 3
	inputList := []int{2, 3, 4, 3, 6}

	got := ListHelper.GetTotalOccurenceInList(inputValue, inputList)
	expected := 2

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}

func TestIsAscending(t *testing.T) {
	inputList := []int{2, 3, 4, 6}
	got := ListHelper.IsAscending(inputList)
	expected := true
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}

	inputList = []int{7, 4, 3, 1}
	got = ListHelper.IsAscending(inputList)
	expected = false
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}

func TestIsDescending(t *testing.T) {
	inputList := []int{2, 3, 4, 3, 6}
	got := ListHelper.IsDescending(inputList)
	expected := false
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}

	inputList = []int{7, 4, 3, 1}
	got = ListHelper.IsDescending(inputList)
	expected = true
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}

func TestDiffersByAtleast(t *testing.T) {
	inputValue := 1
	inputList := []int{2, 3, 3, 4, 6}
	got := ListHelper.DiffersByAtleast(inputList, inputValue)
	expected := false
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}

	inputValue = 1
	inputList = []int{7, 4, 3, 1}
	got = ListHelper.DiffersByAtleast(inputList, inputValue)
	expected = true
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}

func TestDiffersByAtmost(t *testing.T) {
	inputValue := 3
	inputList := []int{2, 3, 3, 4, 9}
	got := ListHelper.DiffersByAtmost(inputList, inputValue)
	expected := false
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}

	inputValue = 3
	inputList = []int{7, 4, 3, 1}
	got = ListHelper.DiffersByAtmost(inputList, inputValue)
	expected = true
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}

func TestRemoveIndex(t *testing.T) {
	inputValue := 1
	inputList := []int{7, 4, 3, 1}
	got := ListHelper.RemoveIndex(inputList, inputValue)
	expected := []int{7, 3, 1}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}