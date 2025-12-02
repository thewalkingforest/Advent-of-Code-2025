package day01

import (
	"testing"
)

func TestSolve1_TestInput(t *testing.T) {
	result := Solve1("../input/01-test.txt")
	expected := 3
	if result != expected {
		t.Errorf("Solve1(test input) = %d; want %d", result, expected)
	}
}

func TestSolve1_RealInput(t *testing.T) {
	result := Solve1("../input/01.txt")
	expected := 1154
	if result != expected {
		t.Errorf("Solve1(real input) = %d; want %d", result, expected)
	}
}

func TestSolve2_TestInput(t *testing.T) {
	result := Solve2("../input/01-test.txt")
	expected := 6 // Update this with the actual expected value from the test input
	if result != expected {
		t.Errorf("Solve2(test input) = %d; want %d", result, expected)
	}
}

func TestSolve2_RealInput(t *testing.T) {
	result := Solve2("../input/01.txt")
	expected := 6819
	if result != expected {
		t.Errorf("Solve2(real input) = %d; want %d", result, expected)
	}
}
