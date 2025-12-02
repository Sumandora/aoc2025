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

func sumValid(ranges []segment, check func(int) bool) int {
	sol := 0

	for _, seg := range ranges {
		for n := seg.begin; n <= seg.end; n++ {
			if check(n) {
				sol += n
			}
		}
	}

	return sol
}

func part1(ranges []segment) int {
	return sumValid(ranges, func(n int) bool {
		str := strconv.Itoa(n)
		return str[0:len(str)/2] == str[len(str)/2:]
	})
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
	return sumValid(ranges, func(n int) bool {
		str := strconv.Itoa(n)
		for l := 1; l <= len(str)/2; l++ {
			if check(str, str[0:l]) {
				return true
			}
		}
		return false
	})
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
