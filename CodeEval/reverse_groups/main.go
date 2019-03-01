package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// CHALLENGE DESCRIPTION:
//
// Given a list of numbers and a positive integer k, reverse the
// elements of the list, k items at a time. If the number of elements is
// not a multiple of k, then the remaining items in the end should be
// left as is.
//
// INPUT SAMPLE:
//
// Your program should accept as its first argument a path to a
// filename. Each line in this file contains a list of numbers and the
// number k, separated by a semicolon. The list of numbers are comma
// delimited. E.g.
//
//		1,2,3,4,5;2
//		1,2,3,4,5;3
//
// OUTPUT SAMPLE:
//
// Print out the new comma separated list of numbers obtained after
// reversing. E.g.
//
//		2,1,4,3,5
//		3,2,1,4,5

func mustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func parseTestCase(s string) ([]int, int) {
	parts := strings.Split(s, ";")
	nums := strings.Split(parts[0], ",")
	xs := make([]int, len(nums))
	for i, t := range nums {
		xs[i] = mustParseInt(t)
	}
	k := mustParseInt(parts[1])
	return xs, k
}

func reverseGroup(xs []int, k int) []int {
	for n := k; n < len(xs); n += k {
		for i, j := n-k, n-1; i < j; i, j = i+1, j-1 {
			xs[i], xs[j] = xs[j], xs[i]
		}
	}
	return xs
}

func commaJoinInts(xs []int) string {
	ss := make([]string, len(xs))
	for i, x := range xs {
		ss[i] = strconv.Itoa(x)
	}
	return strings.Join(ss, ",")
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
		xs, k := parseTestCase(testcase)
		ys := reverseGroup(xs, k)
		fmt.Println(commaJoinInts(ys))
	}
}
