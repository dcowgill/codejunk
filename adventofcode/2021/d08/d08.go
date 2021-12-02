package d08

import (
	"adventofcode2021/lib"
	"fmt"
	"sort"
	"strings"
)

var DIGITS = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

func Run() {
	lib.Run(8, part1, part2)
}

func part1() int64 {
	entries := parseInput(realInput)
	n := 0
	for _, e := range entries {
		for _, d := range e.digits {
			switch len(d) {
			case 2, 3, 4, 7:
				n++
			}
		}
	}
	return int64(n)
}

func part2() int64 {
	// Under the correct wiring, the segments appear with the following
	// frequencies in the digits 0 through 9:
	//
	//	A: 8
	//	B: 6
	//  C: 8
	//  D: 7
	//	E: 4
	//	F: 9
	//	G: 7
	//
	// Thus we can trivially deduce which letters in the scrambled input
	// correspond to segments B, E, F since their frequencies are unique.
	//
	// The remaining unknowns are A, C (8) and D, G (7). Since there are only
	// four permutations, we can simply test each one via substitution.

	// First sort the correct segment patterns so they can be compared to
	// proposed translations one-for-one.
	segments := make([]string, 0, len(DIGITS))
	for k := range DIGITS {
		segments = append(segments, k)
	}
	sort.Strings(segments)

	// Translates the patterns given a mapping from old to new bytes. The
	// individual bytes in each translated pattern are sorted.
	subst := func(tr map[byte]byte, patterns []string) []string {
		var translations []string
		buf := make([]byte, 0, 7)
		for _, pat := range patterns {
			buf = buf[:0]
			for i := 0; i < len(pat); i++ {
				buf = append(buf, tr[pat[i]])
			}
			sort.Slice(buf, func(i, j int) bool { return buf[i] < buf[j] })
			translations = append(translations, string(buf))
		}
		return translations
	}

	// Reports whether the translations segments are correct.
	segmentsMatch := func(translations []string) bool {
		sort.Strings(translations)
		for i, s := range segments {
			if translations[i] != s {
				return false
			}
		}
		return true
	}

	// Tries each of the four possible byte-mappings, returning the one that
	// correctly translates the scrambled patterns back to the originals.
	search := func(patterns []string) map[byte]byte {
		freq := byteFreq(patterns)
		tr := map[byte]byte{
			keysByValue(freq, 6)[0]: 'b',
			keysByValue(freq, 4)[0]: 'e',
			keysByValue(freq, 9)[0]: 'f',
		}
		ac := keysByValue(freq, 8)
		dg := keysByValue(freq, 7)
		for i := range ac {
			for j := range dg {
				tr[ac[i]], tr[ac[(i+1)%2]] = 'a', 'c'
				tr[dg[j]], tr[dg[(j+1)%2]] = 'd', 'g'
				if segmentsMatch(subst(tr, patterns)) {
					return tr
				}
			}
		}
		panic(fmt.Sprintf("no solution found for %q", patterns))
	}

	entries := parseInput(realInput)
	sum := 0
	for _, entry := range entries {
		tr := search(entry.patterns)
		digits := subst(tr, entry.digits)
		sum += digitsToInt(digits)
	}
	return int64(sum)
}

func digitsToInt(digits []string) int {
	pow := 1000
	n := 0
	for _, d := range digits {
		n += pow * DIGITS[d]
		pow /= 10
	}
	return n
}

func byteFreq(a []string) map[byte]int {
	m := make(map[byte]int, 7)
	for _, s := range a {
		for i := 0; i < len(s); i++ {
			m[s[i]]++
		}
	}
	return m
}

func keysByValue(m map[byte]int, value int) []byte {
	var a []byte
	for k, v := range m {
		if v == value {
			a = append(a, k)
		}
	}
	return a
}

type entry struct {
	patterns, digits []string
}

func parseInput(input []string) []entry {
	entries := make([]entry, 0, len(input))
	for _, line := range input {
		parts := strings.Split(line, " | ")
		entries = append(entries, entry{
			patterns: strings.Split(parts[0], " "),
			digits:   strings.Split(parts[1], " "),
		})
	}
	return entries
}
