package main

import (
	"container/list"
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

func toNum(l *list.List) int {
	if l == nil {
		return 0
	}
	n := 0
	for v := l.Front(); v != nil; v = v.Next() {
		n = n*10 + v.Value.(int)
	}
	return n
}

func dfs(cache map[state]*list.List, bank []int, s state) *list.List {
	sol, hit := cache[s]
	if hit {
		return sol
	}

	if len(bank)-s.offset < s.remaining {
		return nil
	}
	if len(bank) == s.offset {
		return nil
	}
	if s.remaining == 0 {
		return nil
	}

	fst := bank[s.offset]

	taken := list.New()
	taken.PushFront(fst)
	if takenList := dfs(cache, bank, state{
		remaining: s.remaining - 1,
		offset:    s.offset + 1,
	}); takenList != nil {
		taken.PushBackList(takenList)
	}
	missed := dfs(cache, bank, state{
		remaining: s.remaining,
		offset:    s.offset + 1,
	})

	var res *list.List
	if toNum(taken) > toNum(missed) {
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
		cache := make(map[state]*list.List)
		max := dfs(cache, bank, state{depth, 0})
		sol += toNum(max)
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
