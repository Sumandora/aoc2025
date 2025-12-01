package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type insn = struct {
	dir byte
	n   int
}

func parse(input string) []insn {
	lines := strings.SplitSeq(strings.TrimSpace(input), "\n")

	var insns []insn

	for line := range lines {
		dir := line[0]
		n, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		insns = append(insns, insn{
			dir, n,
		})
	}

	return insns
}

func part1(insns []insn) int {
	var dial = 50
	var sol = 0

	for _, insn := range insns {

		switch insn.dir {
		case 'L':
			dial = (dial - insn.n) % 100
		case 'R':
			dial = (dial + insn.n) % 100
		default:
			log.Fatal(insn.dir)
		}

		if dial == 0 {
			sol++
		}
	}

	return sol
}

func part2(insns []insn) int {
	var dial = 50
	var sol = 0

	for _, insn := range insns {

		// There's a mathematical way to this, but its too error-prone, so I settled with this.
		for n := insn.n; n > 0; n-- {
			switch insn.dir {
			case 'L':
				dial--
			case 'R':
				dial++
			default:
				log.Fatal(insn.dir)
			}
			if dial%100 == 0 {
				sol++
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

	insns := parse(string(content))

	fmt.Printf("Part 1: %d\n", part1(insns))
	fmt.Printf("Part 2: %d\n", part2(insns))
}
