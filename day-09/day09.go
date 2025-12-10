package day09

import (
	"fmt"
	"log"
	"math"
	"strings"

	"advent-of-code/aoc2025/utils"
)

type Point struct {
	x, y int64
}

type Side struct {
	l Point
	r Point
}

type Extremes struct {
	xMin, xMax int64
	yMin, yMax int64
}

func area(l, r Point) int64 {
	w := r.x - l.x
	if w == 0 {
		w = 1
	} else if w < 0 {
		w = (l.x - r.x) + 1
	}
	h := r.y - l.y
	if h == 0 {
		h = 1
	} else if h < 0 {
		h = (l.y - r.y) + 1
	}
	return w * h
}

func getPoints(input string) ([]Point, Extremes) {
	var points []Point
	xMin := int64(math.MaxInt64)
	xMax := int64(0)
	yMin := int64(math.MaxInt64)
	yMax := int64(0)
	for line := range strings.SplitSeq(input, "\n") {
		parts := strings.Split(line, ",")
		x := utils.ParseInt64(parts[0])
		if x < xMin {
			xMin = x
		}
		if x > xMax {
			xMax = x
		}
		y := utils.ParseInt64(parts[1])
		if y < yMin {
			yMin = y
		}
		if y > yMax {
			yMax = y
		}
		points = append(points, Point{x, y})
	}
	return points, Extremes{xMin, xMax, yMin, yMax}
}

func Solve1(input string) int64 {
	data, err := utils.Ingest(input)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	points, _ := getPoints(data)
	max := int64(0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			a := area(points[i], points[j])
			if a > max {
				max = a
			}
		}
	}
	return max
}

func insideShape(p Point, ps []Point) bool {
	result := false

	p1 := ps[0]
	var p2 Point
	for i := 0; i <= len(ps); i++ {
		p2 = ps[i%len(ps)]
		if p.y > utils.MinInt64(p1.y, p2.y) {
			if p.y <= utils.MaxInt64(p1.y, p2.y) {
				if p.x <= utils.MaxInt64(p1.x, p2.x) {
					intersection := (p.y-p1.y)*(p2.x-p1.x)/(p2.y-p1.y) + p1.x
					if p1.x == p2.x || p.x < intersection {
						result = !result
					}
				}
			}
		}
		p1 = p2
	}

	return result
}

func pointsToSides(ps []Point) []Side {
	var sides []Side
	for i := range ps {
		l := ps[i]
		if i+1 == len(ps) {
			r := ps[0]
			sides = append(sides, Side{l, r})
			break
		} else {
			r := ps[i+1]
			sides = append(sides, Side{l, r})
		}
	}
	return sides
}

func Solve2(input string) int64 {
	data, err := utils.Ingest(input)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	points, extremes := getPoints(data)
	sides := pointsToSides(points)
	fmt.Println(sides)
	_ = points
	_ = extremes
	return -1
}
