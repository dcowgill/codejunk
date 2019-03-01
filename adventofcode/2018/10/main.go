package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
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
	const closeEnough = 80 // supposition: message must fit in a terminal
	points := readPoints()
	numSteps := 0
	for {
		box := boundingBox(points)
		if box.width() < closeEnough && box.height() < closeEnough {
			fmt.Printf("after %d steps:\n", numSteps)
			show(points)
		}
		step(points)
		numSteps++
	}
}

func part2() {
	part1()
}

type point struct {
	position vec
	velocity vec
}

type vec struct {
	x, y int
}

var pointRE = regexp.MustCompile(`position=<\s*(-?\d+),\s*(-?\d+)> velocity=<\s*(-?\d+),\s*(-?\d+)>`)

func readPoints() []point {
	scanner := bufio.NewScanner(os.Stdin)
	var points []point
	for scanner.Scan() {
		m := pointRE.FindStringSubmatch(scanner.Text())
		px, _ := strconv.Atoi(m[1])
		py, _ := strconv.Atoi(m[2])
		vx, _ := strconv.Atoi(m[3])
		vy, _ := strconv.Atoi(m[4])
		points = append(points, point{position: vec{px, py}, velocity: vec{vx, vy}})
	}
	return points
}

func step(points []point) {
	for i, p := range points {
		points[i].position.x += p.velocity.x
		points[i].position.y += p.velocity.y
	}
}

var (
	minVec vec = vec{-math.MaxInt64, -math.MaxInt64}
	maxVec vec = vec{+math.MaxInt64, +math.MaxInt64}
)

type box struct {
	min, max vec
}

func (b box) width() int  { return b.max.x - b.min.x + 1 }
func (b box) height() int { return b.max.y - b.min.y + 1 }

func boundingBox(points []point) box {
	min, max := maxVec, minVec
	for _, p := range points {
		min.x = minInt(min.x, p.position.x)
		min.y = minInt(min.y, p.position.y)
		max.x = maxInt(max.x, p.position.x)
		max.y = maxInt(max.y, p.position.y)
	}
	return box{min, max}
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

func show(points []point) {
	grid := newGrid(boundingBox(points))
	for _, p := range points {
		grid.set(p.position)
	}
	fmt.Println(grid)
}

type grid struct {
	runes [][]rune
	box   box
}

func newGrid(b box) *grid {
	g := &grid{box: b}
	g.runes = make([][]rune, g.box.height())
	for i := range g.runes {
		a := make([]rune, g.box.width())
		for j := range a {
			a[j] = '.'
		}
		g.runes[i] = a
	}
	return g
}

func (g *grid) set(p vec) {
	row := p.y - g.box.min.y
	col := p.x - g.box.min.x
	g.runes[row][col] = '#'
}

func (g *grid) String() string {
	var b strings.Builder
	for _, row := range g.runes {
		b.WriteString(string(row))
		b.WriteByte('\n')
	}
	return b.String()
}
