package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type segment = struct {
	begin int
	end   int
}

func parse(input string) []segment {
	var ranges []segment

	for segs := range strings.SplitSeq(strings.TrimSpace(input), ",") {
		parts := strings.Split(segs, "-")
		beg, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		ranges = append(ranges, segment{
			begin: beg,
			end:   end,
		})
	}

	return ranges
}

func part1(ranges []segment) int {
	sol := 0

	for _, seg := range ranges {
		for n := seg.begin; n <= seg.end; n++ {
			str := strconv.Itoa(n)
			if str[0:len(str)/2] == str[len(str)/2:] {
				sol += n
			}
		}
	}

	return sol
}

func check(n string, sub string) bool {
	for {
		if len(n) == 0 {
			return true
		}
		next, found := strings.CutPrefix(n, sub)
		if !found {
			return false
		}
		n = next
	}
}

func part2(ranges []segment) int {
	sol := 0

	for _, seg := range ranges {
		for n := seg.begin; n <= seg.end; n++ {
			str := strconv.Itoa(n)
			for l := 1; l <= len(str)/2; l++ {
				if check(str, str[0:l]) {
					sol += n
					break
				}
			}
		}
	}

	return sol
}

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	ranges := parse(string(content))

	fmt.Printf("Part 1: %d\n", part1(ranges))
	fmt.Printf("Part 2: %d\n", part2(ranges))
}
