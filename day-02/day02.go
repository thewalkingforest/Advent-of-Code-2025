package day02

import (
	"advent-of-code/aoc2025/utils"
	"log"
	"slices"
	"strconv"
	"strings"
)

func isValidID_P1(num int64) bool {
	str := strconv.FormatInt(num, 10)
	mid := len(str) / 2
	l := str[:mid]
	r := str[mid:]
	return l != r
}

func Solve1(input string) int64 {
	data, err := utils.Ingest(input)
	if err != nil {
		log.Fatalf("failed to ingest: %s", err)
	}
	sum := int64(0)
	for idRange := range strings.SplitSeq(data, ",") {
		bounds := strings.Split(idRange, "-")
		start := utils.ParseInt64(bounds[0])
		end := utils.ParseInt64(bounds[1])
		for i := start; i <= end; i++ {
			if !isValidID_P1(i) {
				sum += i
			}
		}
	}
	return sum
}

func allSame(in []string) bool {
	for i := 0; i < len(in)-1; i++ {
		if in[i] != in[i+1] {
			return false
		}
	}
	return true
}

func isValidID_P2(num int64) bool {
	notValid := false
	valid := true
	if num < 10 {
		return valid
	}
	str := strconv.FormatInt(num, 10)

	if allSame(strings.Split(str, "")) {
		return notValid
	}
	if len(str) == 2 || len(str) == 3 {
		return valid
	}
	if len(str)%2 == 0 {
		mid := len(str) / 2
		l := str[:mid]
		r := str[mid:]
		if l == r {
			return notValid
		}
		if len(str) == 4 {
			return valid
		}
	}
	var validSubstrLens []int
	for i := 2; i < len(str); i++ {
		if len(str)%i == 0 {
			validSubstrLens = append(validSubstrLens, i)
		}
	}
	if len(validSubstrLens) == 0 {
		return valid
	}
	for si := 0; si < len(validSubstrLens); si++ {
		var chunks [][]string = slices.Collect(slices.Chunk(strings.Split(str, ""), validSubstrLens[si]))
		cs := make([]string, len(chunks))
		for i, v := range chunks {
			cs[i] = strings.Join(v, "")
		}
		if allSame(cs) {
			return notValid
		}
	}
	return valid
}

func Solve2(input string) int64 {
	data, err := utils.Ingest(input)
	if err != nil {
		log.Fatalf("failed to ingest: %s", err)
	}
	sum := int64(0)
	rounds := -1
	for idRange := range strings.SplitSeq(data, ",") {
		if rounds == 0 {
			break
		}
		bounds := strings.Split(idRange, "-")
		start := utils.ParseInt64(bounds[0])
		end := utils.ParseInt64(bounds[1])
		for i := start; i <= end; i++ {
			if !isValidID_P2(i) {
				sum += i
			}
		}
		rounds -= 1
	}

	return sum
}
