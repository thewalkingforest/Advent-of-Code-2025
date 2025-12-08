package main

import (
	"fmt"

	day01 "advent-of-code/aoc2025/day-01"
	day02 "advent-of-code/aoc2025/day-02"
	day03 "advent-of-code/aoc2025/day-03"
	day04 "advent-of-code/aoc2025/day-04"
	day05 "advent-of-code/aoc2025/day-05"
	day06 "advent-of-code/aoc2025/day-06"
	day07 "advent-of-code/aoc2025/day-07"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	input01 := "input/01.txt"
	fmt.Println(day01.Solve1(input01))
	fmt.Println(day01.Solve2(input01))

	input02 := "input/02.txt"
	fmt.Println(day02.Solve1(input02))
	fmt.Println(day02.Solve2(input02))

	input03 := "input/03.txt"
	fmt.Println(day03.Solve1(input03))
	fmt.Println(day03.Solve2(input03))

	input04 := "input/04.txt"
	fmt.Println(day04.Solve1(input04))
	fmt.Println(day04.Solve2(input04))

	input05 := "input/05.txt"
	fmt.Println(day05.Solve1(input05))
	fmt.Println(day05.Solve2(input05))

	input06 := "input/06.txt"
	day06.Solve1(input06)
	day06.Solve2(input06)

	input07 := "input/07.txt"
	day07.Solve1(input07)
	day07.Solve2(input07)
}
