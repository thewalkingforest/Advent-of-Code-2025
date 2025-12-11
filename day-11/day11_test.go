package day11

import (
	"testing"
)

func TestSolve1_TestInput(t *testing.T) {
	result := Solve1("../input/11-test.txt")
	expected := int64(5)
	if result != expected {
		t.Errorf("Solve1(test input) = %d; want %d", result, expected)
	}
}

func TestSolve1_RealInput(t *testing.T) {
	result := Solve1("../input/11.txt")
	expected := int64(506)
	if result != expected {
		t.Errorf("Solve1(real input) = %d; want %d", result, expected)
	}
}

func TestSolve2_TestInput(t *testing.T) {
	result := Solve2("../input/11-test-p2.txt")
	expected := int64(2)
	if result != expected {
		t.Errorf("Solve2(test input) = %d; want %d", result, expected)
	}
}

func TestSolve2_RealInput(t *testing.T) {
	result := Solve2("../input/11.txt")
	expected := int64(0)
	if result != expected {
		t.Errorf("Solve2(real input) = %d; want %d", result, expected)
	}
}
