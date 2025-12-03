package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse(input string) [][]int {
	var banks [][]int

	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		var bank []int
		for num := range strings.SplitSeq(line, "") {
			parsed, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			bank = append(bank, parsed)
		}
		banks = append(banks, bank)
	}

	return banks
}

type state = struct {
	remaining int
	offset    int
}

func dfs(cache map[state]string, bank []int, s state) string {
	sol, hit := cache[s]
	if hit {
		return sol
	}

	if len(bank)-s.offset < s.remaining {
		return ""
	}
	if len(bank) == s.offset {
		return ""
	}
	if s.remaining == 0 {
		return ""
	}

	fst := bank[s.offset]

	taken := fmt.Sprint(fst) + dfs(cache, bank, state{
		remaining: s.remaining - 1,
		offset:    s.offset + 1,
	})
	missed := dfs(cache, bank, state{
		remaining: s.remaining,
		offset:    s.offset + 1,
	})

	a, err := strconv.Atoi(taken)
	if err != nil {
		a = 0
	}
	b, err := strconv.Atoi(missed)
	if err != nil {
		b = 0
	}

	var res string
	if a > b {
		res = taken
	} else {
		res = missed
	}

	cache[s] = res
	return res
}

func greatestNum(banks [][]int, depth int) int {
	sol := 0
	for _, bank := range banks {
		cache := make(map[state]string)
		max := dfs(cache, bank, state{depth, 0})
		n, err := strconv.Atoi(max)
		if err != nil {
			log.Fatal(err)
		}
		sol += n
	}
	return sol
}

func part1(banks [][]int) int {
	return greatestNum(banks, 2)
}

func part2(banks [][]int) int {
	return greatestNum(banks, 12)
}

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	banks := parse(string(content))

	fmt.Printf("Part 1: %d\n", part1(banks))
	fmt.Printf("Part 2: %d\n", part2(banks))
}
