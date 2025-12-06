package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type data = struct {
	nums [][]int
	ops  []string
}

func parse(input string) data {
	lines := strings.Split(input, "\n")
	var nums [][]int
	for _, s := range lines[:len(lines)-1] {
		var line []int
		for _, field := range strings.Fields(s) {
			num, err := strconv.Atoi(field)
			if err != nil {
				log.Fatal(err)
			}
			line = append(line, num)
		}
		nums = append(nums, line)
	}
	ops := strings.Fields(lines[len(lines)-1])
	return data{nums, ops}
}

func part1(inp data) int {
	c := 0
	for n := 0; n < len(inp.nums[0]); n++ {
		op := inp.ops[n]
		switch op {
		case "*":
			x := 1
			for _, row := range inp.nums {
				x *= row[n]
			}
			c += x
		case "+":
			x := 0
			for _, row := range inp.nums {
				x += row[n]
			}
			c += x
		}
	}
	return c
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	var chars [][]rune
	for i, line := range lines {
		for j, c := range line {
			for len(chars) <= j {
				chars = append(chars, []rune{})
			}
			for len(chars[j]) <= i {
				chars[j] = append(chars[j], '_')
			}
			chars[j][i] = c
		}
	}

	var transposed_lines []string
	for _, l := range chars {
		s := ""
		for _, r := range l {
			s += string(r)
		}
		transposed_lines = append(transposed_lines, s)
	}

	transposed_text := ""

	for _, l := range transposed_lines {
		transposed_text += strings.TrimSpace(l) + "\n"
	}

	sol := 0
	for group := range strings.SplitSeq(transposed_text, "\n\n") {
		fst_line := strings.Split(group, "\n")[0]
		op := fst_line[len(fst_line)-1]
		var c int
		switch op {
		case '+':
			c = 0
		case '*':
			c = 1
		}
		for i, line := range strings.Split(group, "\n") {
			var line2 string
			if i == 0 {
				line2 = line[:len(line)-1]
			} else {
				line2 = line
			}
			n, err := strconv.Atoi(strings.TrimSpace(line2))
			if err != nil {
				// empty lines
				switch op {
				case '+':
					n = 0
				case '*':
					n = 1
				}
			}

			switch op {
			case '+':
				c += n
			case '*':
				c *= n
			}
		}
		sol += c
	}
	return sol
}

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	trimmed := strings.TrimSpace(string(content))

	fmt.Printf("Part 1: %d\n", part1(parse(trimmed)))
	fmt.Printf("Part 2: %d\n", part2(trimmed))

}
