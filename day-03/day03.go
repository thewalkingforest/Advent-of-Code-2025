package day03

import (
	"advent-of-code/aoc2025/utils"
	"log"
	"strings"
)

func solveAux(input string, lenth int) int64 {
	data, err := utils.Ingest(input)

	if err != nil {
		log.Fatalf("failed to ingest: %s", err)
	}

	result := int64(0)
	for line := range strings.SplitSeq(data, "\n") {
		elms := []byte(line)
		var selected []byte
		start := 0
		for remaining := lenth - 1; remaining >= 0; remaining-- {
			end := len(elms) - remaining
			max, i := utils.MaxIndex(elms[start:end])
			if i < 0 {
				log.Fatalf("failed to get max element")
			}
			selected = append(selected, max)
			start += i + 1
		}
		c := utils.ParseI64(string(selected))
		result += c
	}
	return result
}

func Solve1(input string) int64 {
	return solveAux(input, 2)
}

func Solve2(input string) int64 {
	return solveAux(input, 12)
}
