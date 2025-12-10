package day09

import (
	"advent-of-code/aoc2025/utils"
	"testing"
)

func TestSolve1_TestInput(t *testing.T) {
	result := Solve1("../input/09-test.txt")
	expected := int64(50)
	if result != expected {
		t.Errorf("Solve1(test input) = %d; want %d", result, expected)
	}
}

func TestSolve1_RealInput(t *testing.T) {
	result := Solve1("../input/09.txt")
	expected := int64(4725826296)
	if result != expected {
		t.Errorf("Solve1(real input) = %d; want %d", result, expected)
	}
}

func TestSolve2_TestInput(t *testing.T) {
	result := Solve2("../input/09-test.txt")
	expected := int64(24)
	if result != expected {
		t.Errorf("Solve2(test input) = %d; want %d", result, expected)
	}
}

// func TestSolve2_RealInput(t *testing.T) {
// 	result := Solve2("../input/09.txt")
// 	expected := int64(-1)
// 	if result != expected {
// 		t.Errorf("Solve2(real input) = %d; want %d", result, expected)
// 	}
// }

func Test_insideShape(t *testing.T) {
	data, err := utils.Ingest("../input/09-test.txt")
	if err != nil {
		t.Errorf("Failed to read input: %v", err)
	}
	points, _ := getPoints(data)
	result := insideShape(Point{8, 1}, points)
	expected := true
	if result != expected {
		t.Errorf("insideShape(test input) = %v; want %v", result, expected)
	}
}
