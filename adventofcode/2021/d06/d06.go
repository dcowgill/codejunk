package d06

import "adventofcode2021/lib"

func Run()         { lib.Run(6, part1, part2) }
func part1() int64 { return sim(realInput, 80) }
func part2() int64 { return sim(realInput, 256) }

func sim(input []int, ndays int) int64 {
	const N = 9
	var curr [N]int64
	for _, n := range input {
		curr[n]++
	}
	for ; ndays > 0; ndays-- {
		var next [N]int64
		next[6] = curr[0]
		next[8] = curr[0]
		for i := 1; i < N; i++ {
			next[i-1] += curr[i]
		}
		curr = next
	}
	var nfish int64
	for _, n := range curr {
		nfish += n
	}
	return nfish
}
