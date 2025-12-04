package day04

import (
	"testing"
)

func TestSolve1_TestInput(t *testing.T) {
	result := Solve1("../input/04-test.txt")
	expected := int64(13)
	if result != expected {
		t.Errorf("Solve1(test input) = %d; want %d", result, expected)
	}
}

func TestSolve1_RealInput(t *testing.T) {
	result := Solve1("../input/04.txt")
	expected := int64(1508)
	if result != expected {
		t.Errorf("Solve1(real input) = %d; want %d", result, expected)
	}
}

func TestSolve2_TestInput(t *testing.T) {
	result := Solve2("../input/04-test.txt")
	expected := int64(43)
	if result != expected {
		t.Errorf("Solve2(test input) = %d; want %d", result, expected)
	}
}

func TestSolve2_RealInput(t *testing.T) {
	result := Solve2("../input/04.txt")
	expected := int64(8538)
	if result != expected {
		t.Errorf("Solve2(real input) = %d; want %d", result, expected)
	}
}
