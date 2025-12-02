package day02

import (
	"advent-of-code/aoc2025/utils"
	"fmt"
	"log"
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

func Solve1(input string) {
	data, err := utils.Ingest(input)
	if err != nil {
		log.Fatalf("failed to ingest: %s", err)
	}
	sum := int64(0)
	for idRange := range strings.SplitSeq(data, ",") {
		bounds := strings.Split(idRange, "-")
		start := utils.ParseI64(bounds[0])
		end := utils.ParseI64(bounds[1])
		for i := start; i <= end; i++ {
			if !isValidID_P1(i) {
				sum += i
			}
		}
	}
	fmt.Println(sum)
}

func isValidID_P2(num int64) bool {
	str := strconv.FormatInt(num, 10)
	mid := len(str) / 2
	l := str[:mid]
	r := str[mid:]
	return l != r
}

func Solve2(input string) {
	data, err := utils.Ingest(input)
	if err != nil {
		log.Fatalf("failed to ingest: %s", err)
	}
	sum := int64(0)
	for idRange := range strings.SplitSeq(data, ",") {
		bounds := strings.Split(idRange, "-")
		start := utils.ParseI64(bounds[0])
		end := utils.ParseI64(bounds[1])
		for i := start; i <= end; i++ {
			if !isValidID_P2(i) {
				sum += i
			}
		}
	}
	fmt.Println(sum)
}
