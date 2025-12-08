package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type vec3 = struct {
	x int
	y int
	z int
}

func parse(input string) []vec3 {
	var boxes []vec3

	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		coords := strings.Split(line, ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatal(err)
		}
		z, err := strconv.Atoi(coords[2])
		if err != nil {
			log.Fatal(err)
		}

		boxes = append(boxes, vec3{x, y, z})
	}

	return boxes
}

type closest = struct {
	a    vec3
	b    vec3
	dist float64
}

func dist(a, b vec3) float64 {
	d := vec3{
		x: b.x - a.x,
		y: b.y - a.y,
		z: b.z - a.z,
	}
	return math.Sqrt(float64(d.x*d.x + d.y*d.y + d.z*d.z))

}

func findClosest(coords []vec3) []closest {
	var l []closest

	for i, a := range coords {
		for j := i + 1; j < len(coords); j++ {
			b := coords[j]
			distance := dist(a, b)

			l = append(l, closest{a, b, distance})
		}

	}

	slices.SortFunc(l, func(a, b closest) int { return int(a.dist - b.dist) })

	return l
}

type connection = struct {
	a vec3
	b vec3
}

func flood(connections []connection, i vec3, acc *([]vec3)) {
	if slices.Contains(*acc, i) {
		return
	}
	*acc = append(*acc, i)
	for _, c := range connections {
		if i == c.a {
			flood(connections, c.b, acc)
		}
		if i == c.b {
			flood(connections, c.a, acc)
		}
	}
}

func part1(boxes []vec3) int {
	var circuits []connection

	dists := findClosest(boxes)
	for range 1000 {
		pair := dists[0]
		circuits = append(circuits, connection{pair.b, pair.a})
		circuits = append(circuits, connection{pair.a, pair.b})
		dists = dists[1:]
	}

	var counts []int

	for len(circuits) != 0 {
		var acc []vec3
		flood(circuits, circuits[0].a, &acc)
		circuits = slices.DeleteFunc(circuits, func(i connection) bool { return slices.Contains(acc, i.a) || slices.Contains(acc, i.b) })
		counts = append(counts, len(acc))
	}

	slices.Sort(counts)
	slices.Reverse(counts)

	c := 1
	for _, c2 := range counts[:3] {
		c *= c2
	}
	return c
}

func part2(boxes []vec3) int {
	circuits := []connection{}
	dists := findClosest(boxes)
	for {
		pair := dists[0]
		circuits = append(circuits, connection{pair.b, pair.a})
		circuits = append(circuits, connection{pair.a, pair.b})
		dists = dists[1:]

		var acc []vec3
		flood(circuits, circuits[0].a, &acc)
		if len(acc) == len(boxes) {
			return pair.a.x * pair.b.x
		}
	}
}

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	boxes := parse(string(content))

	fmt.Printf("Part 1: %d\n", part1(boxes))
	fmt.Printf("Part 2: %d\n", part2(boxes))

}
