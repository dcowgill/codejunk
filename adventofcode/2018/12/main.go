package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const plantMark = '#'

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
	plants, rules := readInput()
	for i := 0; i < 20; i++ {
		plants = step(plants, rules)
	}
	fmt.Println(calcsum(plants))
}

func part2() {
	plants, rules := readInput()

	// Display 10 iterations of 1K steps
	prevSum := 0
	gen := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 1000; j++ {
			plants = step(plants, rules)
			gen++
		}
		sum := calcsum(plants)
		fmt.Printf("gen = %5d, count = %3d, min = %5d, max = %5d, sum = %8d (%5d)\n",
			gen, len(plants), plants[0], plants[len(plants)-1], sum, sum-prevSum)
		prevSum = sum
	}

	// Pattern is stable?
	const N = 50 * 1000 * 1000 * 1000
	fmt.Printf("\nafter %d generations: %d\n", N, 52*(N-20)+105872)
}

// Example: the rule ".#.##" is represented by bits "01011".
type rule int8

func makeRule(s string) rule {
	var r rule
	for i := 0; i < 5 && i < len(s); i++ {
		if s[i] == '#' {
			r = r.set(i)
		}
	}
	return r
}

func (r rule) set(n int) rule { return r | 1<<uint(4-n) }
func (r rule) shift() rule    { return (r & ((1 << 4) - 1)) << 1 }

func readInput() (plants []int, rules map[rule]bool) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // read first line
	initial := scanner.Text()[len("initial state: "):]
	for i, b := range initial {
		if b == plantMark {
			plants = append(plants, i)
		}
	}
	rules = make(map[rule]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasSuffix(line, " => #") || len(line) != 10 {
			continue
		}
		rules[makeRule(line)] = true

	}
	return plants, rules
}

func step(plants []int, rules map[rule]bool) []int {
	n := len(plants)
	if n == 0 {
		return nil
	}
	min, max := plants[0], plants[n-1]
	i := 0
	var out []int
	var view rule
	view.set(4) // init
	for pos := min - 2; pos <= max+2; pos++ {
		view = view.shift()
		if i < n && plants[i] == pos+2 {
			view = view.set(4)
			i++
		}
		if rules[view] {
			out = append(out, pos)
		}
	}
	return out
}

func calcsum(xs []int) int {
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum
}
