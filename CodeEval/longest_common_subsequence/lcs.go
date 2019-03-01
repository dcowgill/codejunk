package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type key struct{ x, y string }

type lcsSolver map[key]string

// Create a new LCS solver.
func newLcsSolver() lcsSolver {
	return make(map[key]string)
}

// Wraps 'lcs', ensuring its answer is stored in the solver cache.
func (cache lcsSolver) Lcs(s, t string) string {
	k := key{s, t}
	if answer, ok := cache[k]; ok {
		return answer
	}
	answer := cache.lcs(s, t)
	cache[k] = answer
	return answer
}

// Does the actual work of computing the LCS of (s, t).
func (cache lcsSolver) lcs(s, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	if s[len(s)-1] == t[len(t)-1] {
		return cache.Lcs(s[:len(s)-1], t[:len(t)-1]) + string(s[len(s)-1])
	}
	a := cache.Lcs(s[:len(s)-1], t)
	b := cache.Lcs(s, t[:len(t)-1])
	if len(a) > len(b) {
		return a
	}
	return b
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		testcase := scanner.Text()
		if len(testcase) != 0 {
			xs := strings.Split(scanner.Text(), ";")
			if len(xs) != 2 {
				panic("bad test case")
			}
			s, t := xs[0], xs[1]
			fmt.Println(newLcsSolver().Lcs(s, t))
		}
	}
}
