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

func sumFold(nums []int) int {
	n := 0
	for _, num := range nums {
		n += num
	}
	return n
}

func multFold(nums []int) int {
	n := 1
	for _, num := range nums {
		n *= num
	}
	return n
}

func process(nums []int, op string) int {
	switch op {
	case "+":
		return sumFold(nums)
	case "*":
		return multFold(nums)
	default:
		log.Fatal("Invalid operator: " + op)
	}
	return 0
}

func parse(input string) data {
	lines := strings.Split(input, "\n")
	var nums [][]int
	for _, s := range lines[:len(lines)-1] {
		var line []int
		for field := range strings.FieldsSeq(s) {
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
		var nums []int
		for _, row := range inp.nums {
			nums = append(nums, row[n])
		}

		c += process(nums, op)
	}
	return c
}

func stringToCharMatrix(input string) [][]rune {
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
	return chars
}

func part2(input string) int {
	chars := stringToCharMatrix(input)

	var transposedLines []string
	for _, l := range chars {
		transposedLines = append(transposedLines, strings.TrimSpace(string(l)))
	}

	transposedText := strings.Join(transposedLines, "\n")

	sol := 0
	for group := range strings.SplitSeq(transposedText, "\n\n") {
		fstLine := strings.Split(group, "\n")[0]
		op := fstLine[len(fstLine)-1]

		var nums []int

		for line := range strings.SplitSeq(group, "\n") {
			numString := strings.TrimSpace(strings.TrimRight(line, "+*"))
			if len(numString) == 0 {
				continue
			}
			n, err := strconv.Atoi(strings.TrimSpace(numString))
			if err != nil {
				log.Fatal(err)
			}

			nums = append(nums, n)
		}

		sol += process(nums, string(op))
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
