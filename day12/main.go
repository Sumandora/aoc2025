package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type problem struct {
	width  int
	height int

	shapes []int
}

type shape = [][]bool

type data struct {
	shapes   []shape
	problems []problem
}

func parse(input string) data {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")

	var shapes []shape
	for _, shape := range parts[:len(parts)-1] {
		lines := strings.Split(shape, "\n")
		var shape [][]bool
		for _, l := range lines[1:] {
			var shapeLine []bool
			for _, c := range l {
				shapeLine = append(shapeLine, c == '#')
			}
			shape = append(shape, shapeLine)
		}
		shapes = append(shapes, shape)
	}

	var problems []problem
	for l := range strings.SplitSeq(parts[len(parts)-1], "\n") {
		parts := strings.Split(l, ": ")
		dims := strings.Split(parts[0], "x")

		width, err := strconv.Atoi(dims[0])
		if err != nil {
			log.Fatal(err)
		}

		height, err := strconv.Atoi(dims[1])
		if err != nil {
			log.Fatal(err)
		}

		var counts []int
		for num := range strings.SplitSeq(parts[1], " ") {
			num, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}

			counts = append(counts, num)
		}

		problems = append(problems, problem{width, height, counts})
	}

	return data{shapes, problems}
}

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := parse(string(content))

	filledIns := make([]int, len(data.shapes))
	for i, shape := range data.shapes {
		c := 0
		for _, shapeLine := range shape {
			for _, filled := range shapeLine {
				if filled {
					c++
				}
			}
		}
		filledIns[i] = c
	}

	sol := 0
	for _, problem := range data.problems {
		size := problem.width * problem.height
		needed := 0
		for shape, count := range problem.shapes {
			needed += filledIns[shape] * count
		}

		if needed <= size {
			sol++
		}
	}

	fmt.Printf("Part 1: %d\n", sol)
}
