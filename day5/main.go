package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type span = struct {
	from int
	end  int
}

type data = struct {
	ranges []span
	nums   []int
}

func parse(input string) data {
	var ranges []span
	var nums []int

	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	for seg := range strings.SplitSeq(parts[0], "\n") {
		sides := strings.Split(seg, "-")
		from, err := strconv.Atoi(sides[0])
		if err != nil {
			log.Fatal(err)
		}
		to, err := strconv.Atoi(sides[1])
		if err != nil {
			log.Fatal(err)
		}
		ranges = append(ranges, span{from, to})
	}

	for line := range strings.SplitSeq(parts[1], "\n") {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, num)
	}

	return data{ranges, nums}
}

func part1(data data) int {
	c := 0

	for _, num := range data.nums {
		for _, seg := range data.ranges {
			if seg.from <= num && num <= seg.end {
				c++
				break
			}
		}
	}

	return c
}

func part2(data data) int {
	c := 0

	for i := 0; i < len(data.ranges); i++ {
		data.ranges[i].end++
	}

	sort.Slice(data.ranges, func(i, j int) bool {
		return data.ranges[i].from > data.ranges[j].from
	})

	for i := 0; i < len(data.ranges); i++ {
		curr := data.ranges[i]
		full_overlap := false
		for j := i + 1; j < len(data.ranges); j++ {
			other := data.ranges[j]
			if curr.from >= other.from && curr.end <= other.end {
				full_overlap = true
				break
			}
			if curr.from >= other.from && curr.from <= other.end {
				curr.from = other.end
			}
			if curr.end >= other.from && curr.end <= other.end {
				curr.end = other.from
			}
		}
		if full_overlap {
			continue
		}
		c += curr.end - curr.from
	}

	return c
}

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := parse(string(content))

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}
