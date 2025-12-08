package day06

import (
	"testing"
)

func TestSolve1_TestInput(t *testing.T) {
	result := Solve1("../input/06-test.txt")
	expected := int64(4277556)
	if result != expected {
		t.Errorf("Solve1(test input) = %d; want %d", result, expected)
	}
}

func TestSolve1_RealInput(t *testing.T) {
	result := Solve1("../input/06.txt")
	expected := int64(6757749566978)
	if result != expected {
		t.Errorf("Solve1(real input) = %d; want %d", result, expected)
	}
}

func TestSolve2_TestInput(t *testing.T) {
	result := Solve2("../input/06-test.txt")
	expected := int64(3263827)
	if result != expected {
		t.Errorf("Solve2(test input) = %d; want %d", result, expected)
	}
}

// func TestSolve2_RealInput(t *testing.T) {
// 	result := Solve2("../input/06.txt")
// 	expected := int64(-1)
// 	if result != expected {
// 		t.Errorf("Solve2(real input) = %d; want %d", result, expected)
// 	}
// }
