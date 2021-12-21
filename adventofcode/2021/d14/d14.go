package d14

import "adventofcode2021/lib"

func Run()         { lib.Run(14, part1, part2) }
func part1() int64 { return solve(realInput, 10) }
func part2() int64 { return solve(realInput, 40) }

func solve(prob problem, nsteps int) int64 {
	xs := positiveCounts(explore(prob.template, prob.rules, nsteps))
	return int64(lib.Greatest(xs) - lib.Least(xs))
}

type pair struct{ a, b byte }

func explore(polymer string, rules map[pair]byte, nsteps int) counter {
	type cacheKey struct {
		p pair
		n int
	}
	var (
		search func(p pair, remaining int) counter
		cache  = make(map[cacheKey]counter)
	)
	search = func(p pair, n int) (result counter) {
		k := cacheKey{p, n}
		if f, ok := cache[k]; ok {
			result = f
			return
		}
		defer func() {
			cache[k] = result
		}()
		if n <= 0 {
			return
		}
		x, found := rules[p]
		if !found {
			result = nil
			return
		}
		if n == 1 {
			result = newCounter()
			result.inc(x)
			return
		}
		l, r := pair{p.a, x}, pair{x, p.b}
		result = search(l, n-1).add(search(r, n-1))
		result.inc(x)
		return
	}

	// Count the letters added by expanding each pair in the polymer the
	// specified number of times.
	var sum counter
	for _, p := range stringToPairs(polymer) {
		sum = sum.add(search(p, nsteps))
	}

	// Also count the letters in the starting polymer, since they're not
	// counted by the loop above.
	for i := 0; i < len(polymer); i++ {
		sum.inc(polymer[i])
	}

	return sum
}

func stringToPairs(s string) []pair {
	pairs := make([]pair, 0, len(s)*2)
	for i := 1; i < len(s); i++ {
		pairs = append(pairs, pair{s[i-1], s[i]})
	}
	return pairs
}

const N = 26

// A counter counts frequencies of the letters A-Z.
type counter []int

// Creates a new, empty counter.
func newCounter() counter {
	return make([]int, N)
}

// Increments the count for the ASCII letter represented by b.
// b must be between 'A' or 'Z' or the function panics.
func (c counter) inc(b byte) {
	c[b-'A']++
}

// Returns a copy of the counter.
func (c counter) copy() counter {
	d := newCounter()
	copy(d, c)
	return d
}

// Returns a new counter that is the sum of the two counters.
// Either counter may be nil; if both are, returns nil.
func (f counter) add(g counter) counter {
	switch {
	case f == nil && g == nil:
		return nil
	case f == nil:
		return g.copy()
	case g == nil:
		return f.copy()
	}
	sum := newCounter()
	for i := 0; i < N; i++ {
		sum[i] = f[i] + g[i]
	}
	return sum
}

// Extracts only the frequencies in c that are >= 1.
func positiveCounts(c counter) []int {
	var a []int
	for _, n := range c {
		if n > 0 {
			a = append(a, n)
		}
	}
	return a
}
