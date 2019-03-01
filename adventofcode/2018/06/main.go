package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
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
	points := readPoints()

	// Find the interior points. Create a map to store their areas.
	interiorPoints := interiorPoints(points)
	interior := make(map[int]int, len(interiorPoints)) // id -> area
	for _, p := range interiorPoints {
		interior[p.id] = 0 // insert key
	}

	// Determine the bounding box.
	var max point
	for _, p := range points {
		max.x = maxInt(p.x, max.x)
		max.y = maxInt(p.y, max.y)
	}

	// For every point in the bounding box, find its nearest neighbors. If
	// there is one nearest neighbor and it is interior, increment its area.
	for x := 0; x <= max.x; x++ {
		for y := 0; y <= max.y; y++ {
			var (
				closest   point
				minDist   = math.MaxInt64
				contested bool
			)
			for _, p := range points {
				d := distance(p.x, p.y, x, y)
				if d < minDist {
					closest, minDist, contested = p, d, false
				} else if d == minDist {
					contested = true
				}
			}
			if !contested {
				if _, ok := interior[closest.id]; ok {
					interior[closest.id]++
				}
			}
		}
	}

	// Print the largest area.
	{
		max := 0
		for _, v := range interior {
			max = maxInt(v, max)
		}
		fmt.Println(max)
	}
}

func part2() {
	points := readPoints()
	var max point
	for _, p := range points {
		max.x = maxInt(p.x, max.x)
		max.y = maxInt(p.y, max.y)
	}
	area := 0
	for x := 0; x <= max.x; x++ {
		for y := 0; y <= max.y; y++ {
			total := 0
			for _, p := range points {
				total += distance(p.x, p.y, x, y)
			}
			if total < 10000 {
				area++
			}
		}
	}
	fmt.Println(area)
}

type point struct {
	id   int // one-based
	x, y int
}

func readPoints() []point {
	scanner := bufio.NewScanner(os.Stdin)
	var points []point
	id := 1
	for scanner.Scan() {
		parts := trimSpaceAll(strings.SplitN(scanner.Text(), ",", 2))
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, point{id: id, x: x, y: y})
		id++
	}
	return points
}

// Applies strings.TrimSpace to the slice of strings, in-place.
func trimSpaceAll(ss []string) []string {
	for i, s := range ss {
		ss[i] = strings.TrimSpace(s)
	}
	return ss
}

// Determines which points are interior, i.e. do not have an infinite area.
// An interior point must have a neighbor to its top-left, top-right, bottom-left, and bottom-right.
func interiorPoints(points []point) []point {
	var interior []point
	sort.Slice(points, func(i, j int) bool {
		return points[i].x < points[j].x
	})
	for i, p := range points {
		var top, bot bool
		for j := i - 1; j >= 0; j-- {
			q := points[j]
			if q.x < p.x {
				if q.y < p.y {
					top = true
				} else if q.y > p.y {
					bot = true
				}
				if top && bot {
					break
				}
			}
		}
		if !(top && bot) {
			continue
		}
		top, bot = false, false
		for j := i + 1; j < len(points); j++ {
			q := points[j]
			if q.x > p.x {
				if q.y < p.y {
					top = true
				} else if q.y > p.y {
					bot = true
				}
				if top && bot {
					break
				}
			}
		}
		if top && bot {
			interior = append(interior, p)
		}
	}
	return interior
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func distance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
