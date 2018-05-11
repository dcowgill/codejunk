/*

There's a staircase with N steps, and you can climb 1 or 2 steps at a time.
Given N, write a function that returns the number of unique ways you can climb
the staircase. The order of the steps matters.

For example, if N is 4, then there are 5 unique ways:

1, 1, 1, 1
2, 1, 1
1, 2, 1
1, 1, 2
2, 2

What if, instead of being able to climb 1 or 2 steps at a time, you could climb
any number from a set of positive integers X? For example, if X = {1, 3, 5}, you
could climb 1, 3, or 5 steps at a time. Generalize your function to take in X.

*/
package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func staircase(n int, X []int) int {
	cache := make([]int, n+1)
	cache[0] = 1
	for i := 1; i <= n; i++ {
		t := 0
		for _, x := range X {
			if i-x >= 0 {
				t += cache[i-x]
			}
		}
		cache[i] = t
	}
	return cache[n]
}

func main() {
	var steps int
	var sizes csvInt

	flag.IntVar(&steps, "n", 4, "number of steps")
	flag.Var(&sizes, "x", "set of allowed steps per move (CSV)")
	flag.Parse()

	if len(sizes) == 0 {
		sizes = []int{1}
	}

	fmt.Println(staircase(steps, sizes))
}

type csvInt []int

func (v *csvInt) String() string {
	var b strings.Builder
	sep := ""
	for _, x := range *v {
		b.WriteString(sep)
		b.WriteString(strconv.Itoa(x))
		sep = ","
	}
	return b.String()
}

func (v *csvInt) Set(s string) error {
	ss := strings.FieldsFunc(s, func(c rune) bool { return c == ',' || unicode.IsSpace(c) })
	xs := make([]int, len(ss))
	for i, s := range ss {
		x, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		xs[i] = x
	}
	*v = xs
	return nil
}
