package day05

import (
	"advent-of-code/aoc2025/utils"
	"cmp"
	"log"
	"slices"
	"strings"
)

type Range struct {
	start int64
	end   int64
}

func Solve1(input string) int64 {
	data, err := utils.Ingest(input)
	if err != nil {
		log.Fatalf("failed to ingest: %s", err)
	}

	result := int64(0)
	var ranges []Range
	for line := range strings.SplitSeq(data, "\n") {
		if strings.ContainsRune(line, '-') {
			parts := strings.Split(line, "-")
			start := utils.ParseInt64(parts[0])
			end := utils.ParseInt64(parts[1])
			ranges = append(ranges, Range{start, end})
		} else if len(line) == 0 {
			continue
		} else {
			ing := utils.ParseInt64(line)
			for _, r := range ranges {
				if r.start <= ing && ing <= r.end {
					result++
					break
				}
			}
		}
	}

	return result
}

func Solve2(input string) int64 {
	data, err := utils.Ingest(input)
	if err != nil {
		log.Fatalf("failed to ingest: %s", err)
	}

	var ranges []Range
	for line := range strings.SplitSeq(data, "\n") {
		if strings.ContainsRune(line, '-') {
			parts := strings.Split(line, "-")
			start := utils.ParseInt64(parts[0])
			end := utils.ParseInt64(parts[1])
			ranges = append(ranges, Range{start, end})
		}
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		return cmp.Compare(a.start, b.start)
	})

	var merged []Range

	for i := 0; i < len(ranges); i++ {
		start := ranges[i].start
		end := ranges[i].end
		for j := i + 1; j < len(ranges); j++ {
			if ranges[j].start <= end {
				if ranges[j].end > end {
					end = ranges[j].end
				}
			} else {
				break
			}
			i = j
		}
		merged = append(merged, Range{start, end})
	}

	result := int64(0)
	for _, r := range merged {
		result += (r.end - r.start) + 1
	}

	return result
}
