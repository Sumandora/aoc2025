package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type vec2 struct {
	x int
	y int
}

func (v vec2) Hash() int64 {
	return int64(v.x)<<32 | int64(v.y)
}

func parse(input string) []vec2 {
	var points []vec2
	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		points = append(points, vec2{x, y})
	}
	return points
}

func part1(points []vec2) int {
	sol := 0

	for i, p1 := range points {
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]

			w := p2.x - p1.x
			if w < 0 {
				w = -w
			}
			h := p2.y - p1.y
			if h < 0 {
				h = -h
			}
			w += 1
			h += 1
			area := w * h

			if area > sol {
				sol = area
			}
		}
	}

	return sol
}

type rect = struct {
	p1   vec2
	p2   vec2
	area int
}

type seg = struct {
	beg    int
	length int
}

type span = struct {
	begin int
	end   int
	x     int
}

func hittest(lines *[]span, p vec2) int {
	n := 0

	for _, s := range *lines {
		if s.x <= p.x {
			continue
		}

		if s.begin <= p.y && p.y < s.end {
			n++
		}
	}

	return n
}

func part2(points []vec2) int {
	var horLines []span
	var vertLines []span

	last := points[len(points)-1]
	for _, p := range points {
		if last.x == p.x {
			vertLines = append(vertLines, span{min(p.y, last.y), max(p.y, last.y), p.x})
		} else if last.y == p.y {
			horLines = append(horLines, span{min(p.x, last.x), max(p.x, last.x), p.y})
		} else {
			log.Fatal(last, p)
		}

		last = p
	}

	var rects []rect
	for i, p1 := range points {
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]

			w := p2.x - p1.x
			if w < 0 {
				w = -w
			}
			h := p2.y - p1.y
			if h < 0 {
				h = -h
			}
			w += 1
			h += 1
			area := w * h

			rects = append(rects, rect{p1, p2, area})
		}
	}

	slices.SortFunc(rects, func(a, b rect) int {
		return b.area - a.area
	})

	for _, rect := range rects {
		p1 := rect.p1
		p2 := rect.p2

		minx := min(p1.x, p2.x)
		maxx := max(p1.x, p2.x)
		miny := min(p1.y, p2.y)
		maxy := max(p1.y, p2.y)

		if hittest(&vertLines, vec2{minx, miny})%2 != 1 {
			continue
		}

		valid := true

		for _, s := range vertLines {
			if minx < s.x && s.x < maxx {
				if max(s.begin, miny) < min(s.end, maxy) {
					valid = false
					break
				}
			}
		}

		if !valid {
			continue
		}

		for _, s := range horLines {
			if miny < s.x && s.x < maxy {
				if max(s.begin, minx) < min(s.end, maxx) {
					valid = false
					break
				}
			}
		}

		if !valid {
			continue
		}

		return rect.area
	}

	return -1
}

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	points := parse(string(content))

	fmt.Printf("Part 1: %d\n", part1(points))
	fmt.Printf("Part 2: %d\n", part2(points))
}
