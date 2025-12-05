package day05

import (
	"testing"
)

func TestSolve1_TestInput(t *testing.T) {
	result := Solve1("../input/05-test.txt")
	expected := int64(3)
	if result != expected {
		t.Errorf("Solve1(test input) = %d; want %d", result, expected)
	}
}

func TestSolve1_RealInput(t *testing.T) {
	result := Solve1("../input/05.txt")
	expected := int64(811)
	if result != expected {
		t.Errorf("Solve1(real input) = %d; want %d", result, expected)
	}
}

func TestSolve2_TestInput(t *testing.T) {
	result := Solve2("../input/05-test.txt")
	expected := int64(14)
	if result != expected {
		t.Errorf("Solve2(test input) = %d; want %d", result, expected)
	}
}

func TestSolve2_RealInput(t *testing.T) {
	result := Solve2("../input/05.txt")
	expected := int64(338189277144473)
	if result != expected {
		t.Errorf("Solve2(real input) = %d; want %d", result, expected)
	}
}
