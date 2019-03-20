package main

import (
	"flag"
	"fmt"
)

func main() {
	part := flag.Int("part", 1, "which part")
	flag.Parse()
	switch *part {
	case 1:
		part1()
	case 2:
		part2()
	}
}

func part1() {
	fmt.Println(lastN(initialState(), 286051, 10))
}

func part2() {
	// fmt.Println(numBefore(initialState(), []int8{9, 2, 5, 1, 0}))
	fmt.Println(numBefore(initialState(), []int8{2, 8, 6, 0, 5, 1}))
}

type state struct {
	scores []int8
	elf1   int
	elf2   int
}

func initialState() *state {
	return &state{scores: []int8{3, 7}, elf1: 0, elf2: 1}
}

// Advances the simulation and reports the number of new recipes created (1 or 2).
func (st *state) step() int {
	// Combine recipes to form 1 or 2 new ones.
	sum := st.scores[st.elf1] + st.scores[st.elf2]
	numAdded := 1
	if sum >= 10 {
		st.scores = append(st.scores, 1)
		sum -= 10
		numAdded = 2
	}
	st.scores = append(st.scores, sum)

	// Find the elves' next starting recipes.
	st.elf1 = st.next(st.elf1)
	st.elf2 = st.next(st.elf2)

	// Return the number of recipe scores added.
	return numAdded
}

func (st *state) next(elf int) int {
	return (elf + int(st.scores[elf]) + 1) % len(st.scores)
}

// Reports the n scores that will follow the given number of steps.
func lastN(st *state, numSteps, n int) string {
	for len(st.scores) < numSteps+n {
		st.step()
	}
	scores := st.scores[len(st.scores)-n:]
	b := make([]rune, n)
	for i, x := range scores {
		b[i] = '0' + rune(x)
	}
	return string(b)
}

// Reports the number of recipes to the left of the pattern of scores in want.
func numBefore(st *state, want []int8) int {
	n := len(want)
	for len(st.scores) <= n {
		st.step()
	}
	for {
		if st.step() == 2 {
			// We added two recipe scores, so there are two "tails".
			tail := st.scores[len(st.scores)-n-1 : len(st.scores)-1]
			if match(tail, want) {
				return len(st.scores) - n - 1
			}
		}
		tail := st.scores[len(st.scores)-n:]
		if match(tail, want) {
			return len(st.scores) - n
		}
	}
}

// Do a and b contain the same values?
// Assumes they have equal length.
func match(a, b []int8) bool {
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}
