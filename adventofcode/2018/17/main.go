package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
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
	world := readInput(os.Stdin)
	world.flow(springPoint)
	fmt.Printf("visited = %d, water = %d\n", len(world.visited), len(world.water))
}

func part2() {
	part1()
}

var springPoint = Point{x: 500, y: 0}

// Parses the input.
func readInput(r io.Reader) *World {
	fail := func(line, reason string) {
		panic(fmt.Sprintf("failed to parse line %q: %s", line, reason))
	}
	w := newWorld()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			fail(line, "expected 2 space-separated parts")
		}
		name1, xvalues := parseAssignRange(parts[0])
		name2, yvalues := parseAssignRange(parts[1])
		if name1 == name2 {
			fail(line, fmt.Sprintf("assignment to duplicate name %q", name1))
		}
		if name1 == "y" {
			xvalues, yvalues = yvalues, xvalues
		}
		for _, x := range xvalues {
			for _, y := range yvalues {
				w.addClay(x, y)
			}
		}
	}
	return w
}

// Parses an assignment, e.g. "VAR=123" or "VAR=456...789", where the only valid
// values for VAR are "x" and "y". Returns the range of integers and the
// variable to which they were assigned: "x" or "y".
func parseAssignRange(s string) (name string, values []int) {
	s = strings.TrimSpace(s)
	fail := func() {
		panic(fmt.Sprintf("parse error: expected 'x=INT' or 'x=INT...INT', got %q", s))
	}
	n := strings.IndexByte(s, '=')
	if n < 0 {
		fail()
	}
	name, vals := s[:n], s[n+1:]
	if name != "x" && name != "y" {
		fail()
	}
	return name, parseRange(vals)
}

// Parses a string that represents either a single integer (e.g. "123") or a
// ".."-separated range of integers (e.g. "456...789").
func parseRange(s string) []int {
	fail := func() {
		panic(fmt.Sprintf("failed to parse integer or range from %q", s))
	}
	parseInt := func(t string) int {
		x, err := strconv.Atoi(t)
		if err != nil {
			fail()
		}
		return x
	}
	n := strings.Index(s, "..")
	if n < 0 {
		x, err := strconv.Atoi(s)
		if err != nil {
			fail()
		}
		return []int{x}
	}
	first, last := parseInt(s[:n]), parseInt(s[n+2:])
	if first > last {
		fail()
	}
	values := make([]int, 0, last-first+1)
	for i := first; i <= last; i++ {
		values = append(values, i)
	}
	return values
}

type Point struct {
	x, y int
}

func westOf(p Point) Point  { return Point{p.x - 1, p.y} }
func eastOf(p Point) Point  { return Point{p.x + 1, p.y} }
func beneath(p Point) Point { return Point{p.x, p.y + 1} }

type World struct {
	clay    map[Point]bool // locations of clay blocks
	visited map[Point]bool // sand through which water has passed
	water   map[Point]bool // where water has settled
	min     Point
	max     Point
}

func newWorld() *World {
	return &World{
		clay:    make(map[Point]bool),
		visited: make(map[Point]bool),
		water:   make(map[Point]bool),
		min:     Point{math.MaxInt64, math.MaxInt64},
		max:     Point{math.MinInt64, math.MinInt64},
	}
}

func (w *World) addClay(x, y int) {
	w.min.x = minInt(w.min.x, x)
	w.min.y = minInt(w.min.y, y)
	w.max.x = maxInt(w.max.x, x)
	w.max.y = maxInt(w.max.y, y)
	w.clay[Point{x, y}] = true
}

func (w *World) String() string {
	var b strings.Builder
	for y := w.min.y; y <= w.max.y; y++ {
		for x := w.min.x - 1; x <= w.max.x+1; x++ {
			p := Point{x, y}
			var r byte
			switch {
			case w.clay[p]:
				r = '#'
			case w.water[p]:
				r = '~'
			case w.visited[p]:
				r = '|'
			case p == springPoint:
				r = '+'
			default:
				r = '.'
			}
			b.WriteByte(r)
		}
		b.WriteByte('\n')
	}
	return b.String()
}

// Simulates the flow of water beginning at the given starting point.
func (w *World) flow(start Point) {
	// Helper funcs.
	var (
		inBounds = func(p Point) bool { return p.y <= w.max.y }
		open     = func(p Point) bool { return !w.clay[p] && !w.water[p] }
		visit    = func(p Point) {
			if p.y >= w.min.y && p.y <= w.max.y {
				w.visited[p] = true
			}
		}
		spread = func(p Point, dir func(Point) Point) Point {
			for open(p) && !open(beneath(p)) {
				visit(p)
				p = dir(p)
			}
			return p
		}
	)

	// Never flow out of bounds or into an obstructed square.
	if !inBounds(start) || !open(start) {
		return
	}

	// First try to flow downward.
	if below := beneath(start); open(below) {
		if !w.visited[start] { // but only if we have not done so before
			visit(start)
			w.flow(below)
			// If flowing into the square below filled it with
			// water, allow water to spread horizontally here.
			if w.water[below] {
				goto spread
			}
		}
		return
	}

spread:
	// We cannot flow downward, so spread to the left and right until obstructed.
	// If water can flow downward at some point in either direction, let it do so.
	// Otherwise, water flow is obstructed from both sides. Fill with water.
	west := spread(start, westOf)
	east := spread(start, eastOf)
	if open(beneath(west)) || open(beneath(east)) {
		w.flow(west)
		w.flow(east)
	} else {
		// Fill.
		for p := eastOf(west); p != east; p = eastOf(p) {
			w.water[p] = true
		}
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
