package d01

import "adventofcode2021/lib"

func Run() {
	lib.Run(1, part1, part2)
}

func part1() int64 {
	depths := realInput
	answer := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			answer++
		}
	}
	return int64(answer)
}

func part2() int64 {
	depths := realInput
	answer := 0
	for i := 3; i < len(depths); i++ {
		if depths[i] > depths[i-3] {
			answer++
		}
	}
	return int64(answer)
}
