package main

import (
	"fmt"

	day01 "advent-of-code/aoc2025/day-01"
	day02 "advent-of-code/aoc2025/day-02"
	day03 "advent-of-code/aoc2025/day-03"
	day04 "advent-of-code/aoc2025/day-04"
)

func main() {
	input01 := "input/01.txt"
	fmt.Println(day01.Solve1(input01))
	fmt.Println(day01.Solve2(input01))

	input02 := "input/02.txt"
	fmt.Println(day02.Solve1(input02))
	fmt.Println(day02.Solve2(input02))

	input03 := "input/03.txt"
	fmt.Println(day03.Solve1(input03))
	fmt.Println(day03.Solve2(input03))

	input04 := "input/04-test.txt"
	fmt.Println(day04.Solve1(input04))
	fmt.Println(day04.Solve2(input04))
}
