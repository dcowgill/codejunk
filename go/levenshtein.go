package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strings"
)

func newMatrix(x, y int) [][]int {
	mem := make([]int, x*y)
	mat := make([][]int, x)
	for i := 0; i < x; i++ {
		mat[i] = mem[i*x : i*x+y]
	}
	return mat
}

func least(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Levenshtein(s, t string) int {
	d := newMatrix(len(s)+1, len(t)+1)

	for i := 1; i <= len(s); i++ {
		d[i][0] = i
	}
	for i := 1; i <= len(t); i++ {
		d[0][i] = i
	}

	for j := 1; j <= len(t); j++ {
		for i := 1; i <= len(s); i++ {
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1] // no operation required
			} else {
				d[i][j] = least(
					d[i-1][j]+1, // a deletion
					least(
						d[i][j-1]+1,   // an insertion
						d[i-1][j-1]+1, // a substitution
					),
				)
			}
		}
	}

	return d[len(s)][len(t)]
}

func show(a, b int) string {
	v := make([]string, b-a+1)
	for i := 0; i < len(v); i++ {
		v[i] = fmt.Sprintf("%d", a+i)
	}
	return "[" + strings.Join(v, ",") + "]"
}

func main() {

	s := show(1, 1000)
	t := show(2, 1001)

	f, err := os.Create("levenshtein.pprof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 10; i++ {
		Levenshtein(s, t)
	}
}
