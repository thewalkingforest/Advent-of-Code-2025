package day11

import (
	"log"
	"strings"

	"advent-of-code/aoc2025/utils"
)

type Routes map[string][]string

func parseRoutes(input string) Routes {
	result := make(map[string][]string)

	for line := range strings.SplitSeq(input, "\n") {
		halves := strings.Split(line, ":")
		dev := halves[0]
		links := strings.FieldsFunc(halves[1], func(r rune) bool { return r == ' ' })
		result[dev] = links
	}

	return result
}

// func findRoute(dev string, route []string, ROUTES Routes) []string {
// 	for _, link := range ROUTES[dev] {
// 		if link == "out" {
// 			return append(route, dev, link)
// 		}
// 		r := findRoute(link, append(route, dev), ROUTES)
// 		if len(r) > 0 && r[len(r)-1] == "out" {
// 			return r
// 		}
// 	}
// 	return route
// }

func countRoutes(dev string, ROUTES Routes) (int64, bool) {
	if dev == "out" {
		return 1, true
	}

	count := int64(0)
	for _, link := range ROUTES[dev] {
		childRoutes, found := countRoutes(link, ROUTES)
		if found {
			count += childRoutes
		}
	}
	return count, true
}

type Conditions struct {
	dac bool
	fft bool
}

func (c *Conditions) allMet() bool {
	return c.dac && c.fft
}

func countValidRoutes(dev string, ROUTES Routes, conds Conditions) (int64, bool) {
	switch dev {
	case "dac":
		conds.dac = true
	case "fft":
		conds.fft = true
	case "out":
		if conds.allMet() {
			return 1, true
		} else {
			return 0, false
		}
	}

	count := int64(0)
	for _, link := range ROUTES[dev] {
		childRoutes, found := countValidRoutes(link, ROUTES, conds)
		if found {
			count += childRoutes
		}
	}

	return count, true
}

func Solve1(input string) int64 {
	data, err := utils.Ingest(input)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	routes := parseRoutes(data)
	count, f := countRoutes("you", routes)
	if !f {
		log.Fatalf("no routes found")
	}
	return count
}

func Solve2(input string) int64 {
	data, err := utils.Ingest(input)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	routes := parseRoutes(data)
	var conds Conditions
	conds.dac = false
	conds.fft = false
	count, _ := countValidRoutes("svr", routes, conds)
	if count == 0 {
		log.Fatalf("no routes found")
	}
	return count
}
