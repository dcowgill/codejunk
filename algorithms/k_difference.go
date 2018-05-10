// https://www.geeksforgeeks.org/count-pairs-difference-equal-k/
package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func kdiff(a []int, k int) int {
	sort.Ints(a)
	n := 0
	for i, x := range a {
		for j := i + 1; j < len(a); j++ {
			switch d := a[j] - x; {
			case d > k:
				break
			case d == k:
				n++
			}
		}
	}
	return n
}

func main() {
	var a csvInt
	var k int

	flag.Var(&a, "a", "space- or comma-separated list of ints")
	flag.IntVar(&k, "k", 1, "k-difference")
	flag.Parse()

	n := kdiff(a, k)
	fmt.Printf("a = %+v, k = %d, n = %d\n", a, k, n)
}

// Flag parsing:

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
	ss := strings.FieldsFunc(s, func(c rune) bool {
		return c == ',' || unicode.IsSpace(c)
	})
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
