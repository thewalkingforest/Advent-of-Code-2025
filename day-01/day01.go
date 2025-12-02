package day01

import (
	"advent-of-code/aoc2025/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Solve1(input string) {
	data, err := utils.Ingest(input)
	if err != nil {
		log.Fatalf("failed to ingest: %s", err)
	}

	var pos uint32 = 50
	pointZero := 0
	for line := range strings.SplitSeq(data, "\n") {
		dir := line[0]
		amtstr := line[1:]
		// fmt.Printf("%s %s\n", string(dir), string(amtstr))
		amtw, perr := strconv.ParseUint(amtstr, 10, 32)
		if perr != nil {
			log.Fatalf("failed to parse %s as unit8: %s", amtstr, err)
		}
		amt := uint32(amtw % 100)
		if string(dir) == "L" {
			if amt > pos {
				pos = 100 - (amt - pos)
			} else {
				pos -= amt
			}
		} else if string(dir) == "R" {
			if amt+pos >= 100 {
				pos = (amt + pos) - 100
			} else {
				pos += amt
			}
		} else {
			log.Fatalf("invalid dir %s", string(dir))
		}
		// fmt.Printf("pos: %d\n", pos)
		if pos == 0 {
			pointZero += 1
		}
	}
	fmt.Println(pointZero)
}

func Solve2(input string) {
	data, err := utils.Ingest(input)
	if err != nil {
		log.Fatalf("failed to ingest: %s", err)
	}

	var pos uint32 = 50
	pointZero := 0
	for line := range strings.SplitSeq(data, "\n") {
		dir := line[0]
		amtstr := line[1:]
		// fmt.Printf("%s %s\n", string(dir), string(amtstr))
		amtw, perr := strconv.ParseUint(amtstr, 10, 32)
		if perr != nil {
			log.Fatalf("failed to parse %s as unit8: %s", amtstr, err)
		}
		clicks := uint32(amtw / 100)
		amt := uint32(amtw % 100)
		if string(dir) == "L" {
			if amt > pos {
				last := pos
				pos = 100 - (amt - pos)
				if last != 0 && pos != 0 {
					clicks += 1
				}
			} else {
				pos -= amt
			}
		} else if string(dir) == "R" {
			if amt+pos >= 100 {
				last := pos
				pos = (amt + pos) - 100
				if last != 0 && pos != 0 {
					clicks += 1
				}
			} else {
				pos += amt
			}
		} else {
			log.Fatalf("invalid dir %s", string(dir))
		}
		if pos == 0 {
			pointZero += 1
		}
		pointZero += int(clicks)
		// fmt.Printf("clicks: %d pos: %d pointZero: %d\n", clicks, pos, pointZero)
	}
	fmt.Println(pointZero)
}
