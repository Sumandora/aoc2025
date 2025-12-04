package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parse(input string) [][]rune {
	var grid [][]rune
	lines := strings.SplitSeq(strings.TrimSpace(input), "\n")

	for line := range lines {
		gridLine := make([]rune, len(line))
		for j, c := range line {
			gridLine[j] = c
		}
		grid = append(grid, gridLine)
	}

	return grid
}

func charAt(grid [][]rune, x int, y int) rune {
	if y < 0 || y >= len(grid) {
		return '.'
	}
	line := grid[y]
	if x < 0 || x >= len(line) {
		return '.'
	}
	return line[x]
}

func countNeighbors(grid [][]rune, x int, y int) int {
	c := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if charAt(grid, x+i, y+j) == '@' {
				c++
			}
		}
	}

	return c
}

type xy = struct {
	x int
	y int
}

func accessibleRoles(grid [][]rune) []xy {
	var pos []xy

	for y, gridLine := range grid {
		for x := range gridLine {
			if charAt(grid, x, y) == '@' {
				if countNeighbors(grid, x, y) < 4 {
					pos = append(pos, xy{x, y})
				}
			}
		}
	}

	return pos
}

func part1(grid [][]rune) int {
	return len(accessibleRoles(grid))
}

func part2(grid [][]rune) int {
	sol := 0
	for {
		roles := accessibleRoles(grid)
		if len(roles) == 0 {
			break
		}
		sol += len(roles)
		for _, pos := range roles {
			grid[pos.y][pos.x] = '.'
		}
	}
	return sol
}

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid := parse(string(content))

	fmt.Printf("Part 1: %d\n", part1(grid))
	fmt.Printf("Part 2: %d\n", part2(grid))
}
