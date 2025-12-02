package day02

import (
	"testing"
)

func TestSolve1_TestInput(t *testing.T) {
	result := Solve1("../input/02-test.txt")
	expected := int64(1227775554) // Update this with the actual expected value from the test input
	if result != expected {
		t.Errorf("Solve1(test input) = %d; want %d", result, expected)
	}
}

func TestSolve1_RealInput(t *testing.T) {
	result := Solve1("../input/02.txt")
	expected := int64(24157613387)
	if result != expected {
		t.Errorf("Solve1(real input) = %d; want %d", result, expected)
	}
}

func TestSolve2_TestInput(t *testing.T) {
	result := Solve2("../input/02-test.txt")
	expected := int64(4174379265)
	if result != expected {
		t.Errorf("Solve2(test input) = %d; want %d", result, expected)
	}
}

func TestSolve2_RealInput(t *testing.T) {
	result := Solve2("../input/02.txt")
	expected := int64(33832678380)
	if result != expected {
		t.Errorf("Solve2(real input) = %d; want %d", result, expected)
	}
}
