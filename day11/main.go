package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type connections = map[string][]string

func parse(input string) connections {
	conns := make(connections)
	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		parts := strings.Split(line, ": ")
		connections := strings.Split(parts[1], " ")
		conns[parts[0]] = connections
	}
	return conns
}

func dfs(cache *map[string]int, conns *connections, curr string) int {
	if curr == "out" {
		return 1
	}
	sol, hit := (*cache)[curr]
	if hit {
		return sol
	}

	next := (*conns)[curr]
	c := 0
	for _, conn := range next {
		c += dfs(cache, conns, conn)
	}

	(*cache)[curr] = c
	return c
}

func part1(conns connections) int {
	var cache = make(map[string]int)
	return dfs(&cache, &conns, "you")
}

type state struct {
	curr   string
	sawFft bool
	sawDac bool
}

func dfs2(cache *map[state]int, conns *connections, s state) int {
	if s.curr == "out" {
		if s.sawFft && s.sawDac {
			return 1
		} else {
			return 0
		}
	}
	sol, hit := (*cache)[s]
	if hit {
		return sol
	}

	next := (*conns)[s.curr]
	c := 0
	for _, conn := range next {
		c += dfs2(cache, conns, state{conn, s.sawFft || s.curr == "fft", s.sawDac || s.curr == "dac"})
	}

	(*cache)[s] = c
	return c
}

func part2(conns connections) int {
	var cache = make(map[state]int)
	return dfs2(&cache, &conns, state{"svr", false, false})
}

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	conns := parse(string(content))

	fmt.Printf("Part 1: %d\n", part1(conns))
	fmt.Printf("Part 2: %d\n", part2(conns))
}
