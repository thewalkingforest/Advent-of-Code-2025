package day06

import (
	"advent-of-code/aoc2025/utils"
	"log"
	"strings"
)

func Solve1(inputPath string) int64 {
	input, err := utils.Ingest(inputPath)
	if err != nil {
		log.Fatalf("Failed to ingest input: %v", err)
	}

	var data [][]string
	for line := range strings.SplitSeq(input, "\n") {
		parts := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' })
		data = append(data, parts)
	}
	var aligned [][]int64
	for x := 0; x < len(data[0]); x++ {
		var row []int64
		for y := 0; y < len(data)-1; y++ {
			row = append(row, utils.ParseInt64(data[y][x]))
		}
		aligned = append(aligned, row)
	}
	result := int64(0)
	ops := data[len(data)-1]
	for i, op := range ops {
		var lineResult int64
		switch op {
		case "+":
			lineResult = 0
			for _, num := range aligned[i] {
				lineResult += num
			}
		case "*":
			lineResult = 1
			for _, num := range aligned[i] {
				lineResult *= num
			}
		default:
			log.Fatalf("Invalid operator %s", op)
		}
		result += lineResult
	}

	return result
}

func Solve2(inputPath string) int64 {
	input, err := utils.Ingest(inputPath)
	if err != nil {
		log.Fatalf("Failed to ingest input: %v", err)
	}

	var data [][]string
	for line := range strings.SplitSeq(input, "\n") {
		parts := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' })
		data = append(data, parts)
	}

	result := int64(0)

	return result
}
