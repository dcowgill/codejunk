package lib

import (
	"fmt"
	"os"
)

func Run(day int, p1, p2 func() int64) {
	runPart(day, 1, p1)
	runPart(day, 2, p2)
}

func runPart(day, part int, f func() int64) {
	answer := f()
	fmt.Printf("day %2d part %d: %d%s\n", day, part, answer, verify(day, part, answer))
}

func verify(day, part int, answer int64) string {
	var knownAnswers = map[[2]int]int64{
		{1, 1}:  1215,
		{1, 2}:  1150,
		{2, 1}:  2147104,
		{2, 1}:  2147104,
		{2, 2}:  2044620088,
		{3, 1}:  2035764,
		{3, 2}:  2817661,
		{4, 1}:  16674,
		{4, 2}:  7075,
		{5, 1}:  7085,
		{5, 2}:  20271,
		{6, 1}:  383160,
		{6, 2}:  1721148811504,
		{7, 1}:  349357,
		{7, 2}:  96708205,
		{8, 1}:  362,
		{8, 2}:  1020159,
		{9, 1}:  588,
		{9, 2}:  964712,
		{10, 1}: 321237,
		{10, 2}: 2360030859,
		{11, 1}: 1601,
		{11, 2}: 368,
		{12, 1}: 3450,
		{12, 2}: 96528,
		{13, 1}: 695,
		{14, 1}: 3095,
		{14, 2}: 3152788426516,
		{15, 1}: 562,
		{15, 2}: 2874,
		{16, 1}: 873,
		{16, 2}: 402817863665,
		{17, 1}: 4560,
		{17, 2}: 3344,
		{18, 1}: 4173,
		{18, 2}: 4706,
	}
	if expected, ok := knownAnswers[[2]int{day, part}]; ok {
		if answer != expected {
			fmt.Printf("ERROR: day %02d part %d: got %d, expected %d", day, part, answer, expected)
			os.Exit(1)
		}
		return ""
	}
	return " (UNVERIFIED)"
}
