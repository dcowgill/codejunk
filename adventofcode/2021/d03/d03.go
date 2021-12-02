package d03

import (
	"adventofcode2021/lib"
	"fmt"
)

func Run() {
	lib.Run(3, part1, part2)
}

func part1() int64 {
	var (
		numbers = sanity(realInput)
		counts  = countBits(numbers)
		γ, ε    int
		pow     = 1
	)
	for i := len(counts) - 1; i >= 0; i-- {
		m, l := 0, 1
		if counts[i][1] > counts[i][0] {
			m, l = l, m
		}
		γ += m * pow
		ε += l * pow
		pow *= 2
	}
	return int64(γ * ε)
}

func part2() int64 {
	var (
		numbers = sanity(realInput)

		mc = func(n0, n1 int) byte {
			if n0 > n1 {
				return '0'
			}
			return '1'
		}

		lc = func(n0, n1 int) byte {
			if n1 < n0 {
				return '1'
			}
			return '0'
		}

		xs = search(numbers, mc)
		ys = search(numbers, lc)
		x  = binStrToDec(xs[0])
		y  = binStrToDec(ys[0])
	)
	return int64(x * y)
}

func sanity(xs []string) []string {
	if len(xs) == 0 {
		panic("input is empty")
	}
	for i, x := range xs {
		if len(x) != len(xs[0]) {
			panic(fmt.Sprintf("number %d has %d bits, expected %d", i, len(x), len(xs[0])))
		}
	}
	return xs
}

func countBits(xs []string) [][2]int {
	var (
		numBits = len(xs[0])
		counts  = make([][2]int, numBits)
	)
	for _, x := range xs {
		for i := 0; i < len(x); i++ {
			b := 0
			if x[i] == '1' {
				b = 1
			}
			counts[i][b]++
		}
	}
	return counts
}

func search(numbers []string, chooseBit func(n0, n1 int) byte) []string {
	numBits := len(numbers[0])
	numbers = copyStrings(numbers)
	for i := 0; i < numBits; i++ {
		c := countBits(numbers)
		b := chooseBit(c[i][0], c[i][1])
		numbers = filter(numbers, b, i)
		if len(numbers) <= 1 {
			break
		}
	}
	if len(numbers) != 1 {
		panic("search: expected exactly 1 number")
	}
	return numbers
}

func copyStrings(xs []string) []string {
	ys := make([]string, len(xs))
	copy(ys, xs)
	return ys
}

func filter(xs []string, bit byte, pos int) []string {
	j := 0
	for _, x := range xs {
		if x[pos] == bit {
			xs[j] = x
			j++
		}
	}
	return xs[:j]
}

func binStrToDec(x string) int {
	pow := 1
	n := 0
	for i := len(x) - 1; i >= 0; i-- {
		if x[i] == '1' {
			n += pow
		}
		pow *= 2
	}
	return n
}
