package MathHelper_test

import (
	"AdventOfCode/Helpers/MathHelper"
	"testing"
)

func TestAbsDiffInt(t *testing.T) {
	//Test positives different
	inputInt1 := 2
	inputInt2 := 4
	got := MathHelper.AbsDiffInt(inputInt1, inputInt2)
	expected := 2
	if got != expected {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}

	//Test positives same
	inputInt1 = 2
	inputInt2 = 2
	got = MathHelper.AbsDiffInt(inputInt1, inputInt2)
	expected = 0
	if got != expected {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}

	//Test negatives same
	inputInt1 = -2
	inputInt2 = -2
	got = MathHelper.AbsDiffInt(inputInt1, inputInt2)
	expected = 0
	if got != expected {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}

	//Test negatives different
	inputInt1 = -3
	inputInt2 = -2
	got = MathHelper.AbsDiffInt(inputInt1, inputInt2)
	expected = 1
	if got != expected {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}

	//Test both negatives and positives
	inputInt1 = -3
	inputInt2 = 2
	got = MathHelper.AbsDiffInt(inputInt1, inputInt2)
	expected = 5
	if got != expected {
		t.Errorf("Error, Expected %v, but got %v", expected, got)
	}
}
