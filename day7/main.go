package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type xy = struct {
	x int
	y int
}

type data = struct {
	startPos xy
	grid     [][]rune
}

func parse(input string) data {
	var startPos xy
	var grid [][]rune
	for y, line := range strings.Split(input, "\n") {
		for x, c := range line {
			for len(grid) <= y {
				grid = append(grid, []rune{})
			}
			for len(grid[y]) <= x {
				grid[y] = append(grid[y], '#')
			}
			grid[y][x] = c
			if c == 'S' {
				startPos = xy{x, y}
			}
		}
	}

	return data{startPos, grid}
}

var visited = make(map[xy]bool)

func part1_dfs(input data, pos xy) {
	_, hit := visited[pos]
	if hit {
		return
	}

	if pos.y >= len(input.grid) {
		return
	}

	c := input.grid[pos.y][pos.x]
	if c != '^' {
		part1_dfs(input, xy{pos.x, pos.y + 1})
		return
	}

	visited[pos] = true
	part1_dfs(input, xy{pos.x - 1, pos.y})
	part1_dfs(input, xy{pos.x + 1, pos.y})
}

func part1(input data) int {
	part1_dfs(input, input.startPos)
	return len(visited)
}

var cache = make(map[xy]int)

func part2_dfs(input data, pos xy) int {
	count, hit := cache[pos]
	if hit {
		return count
	}

	if pos.y >= len(input.grid) {
		return 1
	}

	c := input.grid[pos.y][pos.x]
	if c != '^' {
		res := part2_dfs(input, xy{pos.x, pos.y + 1})
		cache[pos] = res
		return res
	}

	res := part2_dfs(input, xy{pos.x - 1, pos.y}) + part2_dfs(input, xy{pos.x + 1, pos.y})
	cache[pos] = res
	return res
}

func part2(input data) int {
	return part2_dfs(input, input.startPos)
}

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := parse(strings.TrimSpace(string(content)))

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}
