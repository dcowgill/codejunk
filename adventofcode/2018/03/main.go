package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	part := flag.Int("part", 1, "which part")
	flag.Parse()
	switch *part {
	case 1:
		part1()
	case 2:
		part2()
	}
}

func part1() {
	const N = 1000
	grid := newGrid(N, N)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		c := parseClaim(scanner.Text())
		for x := c.x1; x <= c.x2; x++ {
			for y := c.y1; y <= c.y2; y++ {
				grid[x][y]++
			}
		}
	}
	n := 0
	for _, row := range grid {
		for _, count := range row {
			if count > 1 {
				n++
			}
		}
	}
	fmt.Println(n)
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	var claims []claim
	for scanner.Scan() {
		claims = append(claims, parseClaim(scanner.Text()))
	}
outer:
	for i, c1 := range claims {
		for j, c2 := range claims {
			if i != j && c1.overlaps(c2) {
				continue outer
			}
		}
		fmt.Printf("%+v\n", c1)
	}
}

// Returns an r-by-c matrix of integers.
func newGrid(r, c int) [][]int {
	mem := make([]int, r*c)
	mat := make([][]int, r)
	for i := 0; i < r; i++ {
		mat[i], mem = mem[:c], mem[c:]
	}
	return mat
}

type claim struct {
	id, x1, y1, x2, y2 int
}

func makeClaim(id, x, y, w, h int) claim {
	return claim{id: id, x1: x, y1: y, x2: x + w - 1, y2: y + h - 1}
}

func (c claim) overlaps(d claim) bool {
	return !((c.x2 < d.x1 || c.x1 > d.x2) || (c.y2 < d.y1 || c.y1 > d.y2))
}

var claimRE = regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)

func parseClaim(s string) claim {
	m := claimRE.FindStringSubmatch(s)
	if m == nil {
		panic(fmt.Sprintf("failed to parse %q", s))
	}
	return makeClaim(
		atoi(m[1]), // id
		atoi(m[2]), // x1
		atoi(m[3]), // y1
		atoi(m[4]), // width
		atoi(m[5]), // height
	)
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("strconv.Atoi(%q) failed: %v", s, err))
	}
	return n
}
