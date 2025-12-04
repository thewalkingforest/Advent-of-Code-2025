package day04

import (
	"advent-of-code/aoc2025/utils"
	"log"
	"strings"
)

type Pair struct {
	x int
	y int
}

func countNBors(i int, slots []byte, stride int) int {
	nbors := []Pair{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	y := i / stride
	x := i % stride

	result := 0
	for _, p := range nbors {
		nx := x + p.x
		ny := y + p.y
		if nx < 0 || ny < 0 || nx >= stride {
			continue
		}
		idx := nx + (ny * stride)
		if idx >= len(slots) {
			continue
		}
		if slots[idx] == '@' {
			result++
		}
	}

	return result
}

func rollsRemoved(slots []byte, stride int) (int64, []int) {
	result := int64(0)
	var removeable []int
	for i, e := range slots {
		if e == '@' {
			if countNBors(i, slots, stride) < 4 {
				result++
				removeable = append(removeable, i)
			}
		}
	}
	return result, removeable
}

func Solve1(input string) int64 {
	data, err := utils.Ingest(input)
	if err != nil {
		log.Fatalf("failed to ingest: %s", err)
	}
	lines := strings.Split(data, "\n")
	stride := len(lines[0])
	slots := []byte(strings.Join(lines, ""))
	removed, _ := rollsRemoved(slots, stride)
	return removed
}

func Solve2(input string) int64 {
	data, err := utils.Ingest(input)
	if err != nil {
		log.Fatalf("failed to ingest: %s", err)
	}
	result := int64(0)
	lines := strings.Split(data, "\n")
	stride := len(lines[0])
	slots := []byte(strings.Join(lines, ""))
	removed := int64(0)
	var removeable []int
	for true {
		removed, removeable = rollsRemoved(slots, stride)
		if removed == 0 {
			break
		}
		result += removed
		for _, p := range removeable {
			slots[p] = '.'
		}
	}

	return result
}
