package ConvertHelper_test

import (
	"AdventOfCode/Helpers/ConvertHelper"
	"fmt"
	"reflect"
	"testing"
)

func TestStringToInt(t *testing.T) {
	input := "2"

	got := ConvertHelper.StringToInt(input)
	expected := 2

	if got != expected {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}

func TestStringToList(t *testing.T) {
	test := "2 5"

	got := ConvertHelper.StringToList(test)
	expected := []string{"2", "5"}

	for i := 0; i < len(got); i++ {

		fmt.Println(got[i])
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}

func TestStringToTwoSplices(t *testing.T) {
	inputText := []string{"2 5", "4 2", "1 7"}

	got1, got2 := ConvertHelper.StringToTwoSplices(inputText)
	expected1 := []int{2, 4, 1}
	expected2 := []int{5, 2, 7}

	if !reflect.DeepEqual(got1, expected1) || !reflect.DeepEqual(got2, expected2) {
		t.Errorf("Error, \n1: Expected %v, but got %v \n2: Expected %v, but got %v", expected1, got1, expected2, got2)
	}
}

func TestStringSpliceToIntSplice(t *testing.T) {
	inputText := []string{"25", "42", "17"}

	got := ConvertHelper.StringSpliceToIntSplice(inputText)
	expected := []int{25, 42, 17}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}

func TestStringSpliceToIntSpliceList(t *testing.T) {
	inputText := []string{"2 5", "4 2", "1 7"}

	got := ConvertHelper.StringSpliceToIntSpliceList(inputText)
	expected := [][]int{{2, 5}, {4, 2}, {1, 7}}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}
