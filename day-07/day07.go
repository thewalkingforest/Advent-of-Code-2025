package day07

import (
	"log"
	"strings"

	"advent-of-code/aoc2025/utils"
)

func Solve1(input string) int64 {
	data, err := utils.Ingest(input)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	manifold := strings.Split(data, "\n")
	start := strings.Index(manifold[0], "S")
	active := make([]bool, len(manifold[0]))
	active[start] = true
	result := int64(0)
	for i := 1; i < len(manifold); i++ {
		splits := utils.IndexAll(manifold[i], "^")
		for _, s := range splits {
			if active[s] {
				result++
				active[s] = false
				active[s-1] = true
				active[s+1] = true
			}
		}
	}

	return result
}

func Solve2(input string) int64 {
	data, err := utils.Ingest(input)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	manifold := strings.Split(data, "\n")
	start := strings.Index(manifold[0], "S")
	active := make([]bool, len(manifold[0]))
	active[start] = true
	result := int64(0)
	for i := 1; i < len(manifold); i++ {
		splits := utils.IndexAll(manifold[i], "^")
		for _, s := range splits {
			if active[s] {
				active[s] = false
				active[s-1] = true
				active[s+1] = true
			}
		}
	}

	return result
}
