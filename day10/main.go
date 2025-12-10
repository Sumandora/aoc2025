package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/aclements/go-z3/z3"
)

type button struct {
	swaps []int
}

type machine struct {
	wantedState []bool
	buttons     []button
	joltage     []int
}

func parse(input string) []machine {
	var machines []machine
	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		words := strings.Split(line, " ")

		wanted := words[0]
		var wantedBools []bool
		for _, c := range wanted[1 : len(wanted)-1] {
			switch c {
			case '.':
				wantedBools = append(wantedBools, false)
			case '#':
				wantedBools = append(wantedBools, true)
			default:
				log.Fatal(c)
			}
		}

		var buttons []button
		for _, word := range words[1 : len(words)-1] {
			part := word[1 : len(word)-1]
			var nums []int
			for num := range strings.SplitSeq(part, ",") {
				n, err := strconv.Atoi(num)
				if err != nil {
					log.Fatal(err)
				}
				nums = append(nums, n)
			}
			buttons = append(buttons, button{nums})
		}

		lastWord := words[len(words)-1]
		var joltage []int
		for num := range strings.SplitSeq(lastWord[1:len(lastWord)-1], ",") {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			joltage = append(joltage, n)
		}

		machines = append(machines, machine{
			wantedState: wantedBools,
			buttons:     buttons,
			joltage:     joltage,
		})
	}

	return machines
}

func bfs(depth int, machine *machine, states [][]bool) int {
	found := false
	for _, state := range states {
		if slices.Equal(state, machine.wantedState) {
			found = true
			break
		}
	}
	if found {
		return depth
	}

	var nextStates [][]bool
	for _, state := range states {
		for _, btn := range machine.buttons {
			var updatedState = make([]bool, len(state))
			copy(updatedState, state)
			for _, idx := range btn.swaps {
				updatedState[idx] = !updatedState[idx]
			}
			nextStates = append(nextStates, updatedState)
		}
	}
	return bfs(depth+1, machine, nextStates)
}

func part1(machines []machine) int {
	sum := 0
	for _, machine := range machines {
		var b = [][]bool{make([]bool, len(machine.wantedState))}
		sum += bfs(0, &machine, b)
	}
	return sum
}

func part2(machines []machine) int {
	sum := 0
	ctx := z3.NewContext(nil)
	intSort := ctx.IntSort()
	zero := ctx.FromInt(0, intSort).(z3.Int)
	for _, machine := range machines {
		s := z3.NewSolver(ctx)

		var buttonsNums []z3.Int
		for i := range machine.buttons {
			val := ctx.IntConst(string('A' + i))
			buttonsNums = append(buttonsNums, val)
			s.Assert(val.GE(zero))
		}

		for i, target := range machine.joltage {
			var responsible []z3.Int
			for j, button := range machine.buttons {
				if slices.Contains(button.swaps, i) {
					responsible = append(responsible, buttonsNums[j])
				}
			}

			joltSum := responsible[0]
			for _, v := range responsible[1:] {
				joltSum = joltSum.Add(v)
			}

			s.Assert(joltSum.Eq(ctx.FromInt(int64(target), intSort).(z3.Int)))
		}

		total := buttonsNums[0]
		for _, otherBtn := range buttonsNums[1:] {
			total = total.Add(otherBtn)
		}

		bestSum := int64(0)
		for {
			sat, err := s.Check()
			if err != nil {
				log.Fatal(err)
			}
			if !sat {
				break
			}

			model := s.Model()
			val := model.Eval(total, true).(z3.Int)
			newBestSum, _, _ := val.AsInt64()
			bestSum = newBestSum

			s.Assert(total.LT(val))
		}

		if bestSum == 0 {
			log.Fatal("Unsolvable", machine)
		}
		sum += int(bestSum)
	}
	return sum
}

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	machines := parse(string(content))

	fmt.Printf("Part 1: %d\n", part1(machines))
	fmt.Printf("Part 2: %d\n", part2(machines))
}
