package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
	var c2, c3 int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Count letter frequencies.
		m := make(map[rune]int)
		for _, ch := range scanner.Text() {
			m[ch]++
		}
		// Determine if any letter appears exactly 2 or 3 times.
		var has2, has3 bool
		for _, count := range m {
			switch count {
			case 2:
				has2 = true
			case 3:
				has3 = true
			}
			if has2 && has3 {
				break
			}
		}
		// Increment counters.
		if has2 {
			c2++
		}
		if has3 {
			c3++
		}
	}
	// Print checksum.
	fmt.Println(c2 * c3)
}

func part2() {
	// Reads all box IDs into memory.
	var ids [][]rune
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ids = append(ids, []rune(scanner.Text()))
	}
	// Compare all pairs of IDs until we find the one that differs by one rune.
	for _, a := range ids {
		for _, b := range ids {
			if i := diffpos(a, b); i >= 0 {
				fmt.Printf("%s%s\n", string(a[:i]), string(a[i+1:]))
				return
			}
		}
	}
}

// Reports the index i where a[i] and b[i] are unequal, if and only if there is
// exactly one such index and len(a) == len(b). Else returns -1.
func diffpos(a, b []rune) int {
	const notfound = -1
	if len(a) != len(b) {
		return notfound
	}
	pos := notfound
	for i := range a {
		if a[i] != b[i] {
			if pos != notfound {
				return notfound
			}
			pos = i
		}
	}
	return pos
}
