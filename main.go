package main

import (
	"fmt"

	day01 "advent-of-code/aoc2025/day-01"
	day02 "advent-of-code/aoc2025/day-02"
)

func main() {
	input01 := "input/01.txt"
	fmt.Println(day01.Solve1(input01))
	fmt.Println(day01.Solve2(input01))

	input02 := "input/02.txt"
	fmt.Println(day02.Solve1(input02))
	fmt.Println(day02.Solve2(input02))
}
