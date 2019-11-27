package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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
	bots := readInput(os.Stdin)
	var maxbot Nanobot
	for _, bot := range bots {
		if bot.radius > maxbot.radius {
			maxbot = bot
		}
	}
	numInRange := 0
	for _, bot := range bots {
		if distance(maxbot.pos, bot.pos) <= maxbot.radius {
			numInRange++
		}
	}
	fmt.Println(numInRange)
}

func part2() {
	bots := readInput(os.Stdin)
	p := search(bots)
	fmt.Printf("best = %+v (range=%d) (dist=%d)\n", p, numInRange(p, bots), distance(p, Point{0, 0}))

	// positions := make([]Point, len(bots))
	// for i, bot := range bots {
	// 	positions[i] = bot.pos
	// }
	// centroid := centroid(positions)
	// fmt.Printf("centroid = %+v, num in range = %d\n", centroid, numInRange(centroid, bots))

}

var nanobotRE = regexp.MustCompile(`^pos=<(-?\d+),(-?\d+),(-?\d+)>, r=(\d+)$`)

func readInput(r io.Reader) []Nanobot {
	var bots []Nanobot
	scanner := bufio.NewScanner(r)
	idSeq := 1
	for scanner.Scan() {
		m := nanobotRE.FindStringSubmatch(scanner.Text())
		if m == nil {
			panic(fmt.Sprintf("failed to parse line: %q", scanner.Text()))
		}
		bots = append(bots, Nanobot{
			id:     idSeq,
			pos:    Point{x: atoi(m[1]), y: atoi(m[2]), z: atoi(m[3])},
			radius: atoi(m[4]),
		})
		idSeq++
	}
	return bots
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

type Nanobot struct {
	id     int
	pos    Point
	radius int
}

type Point struct {
	x, y, z int
}

func (p Point) move(d Direction) Point {
	return Point{p.x + d.x, p.y + d.y, p.z + d.z}
}

type Box struct {
	min, max Point
}

type Direction Point

var directions = []Direction{
	Direction{x: +1},
	Direction{x: -1},
	Direction{y: -1},
	Direction{y: +1},
	Direction{z: -1},
	Direction{z: +1},
}

// Reports how many bots are in range of the point.
func numInRange(p Point, bots []Nanobot) int {
	n := 0
	for _, bot := range bots {
		if distance(p, bot.pos) <= bot.radius {
			n++
		}
	}
	return n
}

func octree(min, max Point) []Box {
	// (10,10,10) (19,19,19)

	// top half:
	//
	// 1: (10,14,10) (14,19,14)
	// 2: (14,14,10) (19,19,14)
	// 3: (10,14,14) (14,19,19)
	// 4: (14,14,14) (19,19,19)

	// bottom half:
	//
	// 5: (10,10,10) (14,14,14)
	// 6: (14,10,10) (19,14,14)
	// 7: (10,10,14) (14,14,19)
	// 8: (14,10,14) (19,14,19)
}

func isPointInRangeOfBox(box Box, p Point, radius int) bool {
	if (box.min.x <= p.x && p.x <= box.max.x) &&
		(box.min.y <= p.y && p.y <= box.max.y) &&
		(box.min.z <= p.z && p.z <= box.max.z) {
		return true
	}

	// for each plane, calculate distance?

	return false
}

func distanceToRect(min, max, p Point) int {

}

func numInBotsInRangeOfBox(box Point, bots []Nanobot) int {
	n := 0
	for _, bot := range bots {
		if isPointInRangeOfBox(box, bot.pos, bot.radius) {
			n++
		}
	}
	return n
}

// // Hill climbing attempt.
// func search(bots []Nanobot) Point {
// 	// Create the initial frontier from the bots' positions.
// 	var (
// 		frontier = make(ItemHeap, len(bots))
// 		visited  = make(map[Point]bool)
// 	)
// 	for i, bot := range bots {
// 		frontier[i] = Item{bot.pos, numInRange(bot.pos, bots)}
// 		visited[bot.pos] = true
// 	}
// 	heap.Init(&frontier)
// 	// Climb.
// 	var best Item
// 	for len(frontier) != 0 {
// 		curr := heap.Pop(&frontier).(Item)
// 		if curr.n > best.n || (curr.n == best.n && originDist(curr.p) < originDist(best.p)) {
// 			fmt.Printf("new best = %+v\n", curr)
// 			best = curr
// 		}
// 		for x := -1; x <= +1; x++ {
// 			for y := -1; y <= +1; y++ {
// 				for z := -1; z <= +1; z++ {
// 					if p := curr.p.move(Direction{x, y, z}); !visited[p] {
// 						visited[p] = true
// 						if n := numInRange(p, bots); n > curr.n || n >= best.n {
// 							heap.Push(&frontier, Item{p, n})
// 						}
// 					}
// 				}
// 			}
// 		}
// 		// for _, dir := range directions {
// 		// 	if p := curr.p.move(dir); !visited[p] {
// 		// 		visited[p] = true
// 		// 		if n := numInRange(p, bots); n > curr.n || n >= best.n {
// 		// 			heap.Push(&frontier, Item{p, n})
// 		// 		}
// 		// 	}
// 		// }
// 	}
// 	return best.p
// }
// type Item struct {
// 	p Point
// 	n int // number of bots in range of p
// }
// type ItemHeap []Item
// func (h ItemHeap) Len() int           { return len(h) }
// func (h ItemHeap) Less(i, j int) bool { return h[i].n > h[j].n }
// func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h *ItemHeap) Push(x interface{}) {
// 	*h = append(*h, x.(Item))
// }
// func (h *ItemHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

func distance(p, q Point) int {
	return abs(p.x-q.x) + abs(p.y-q.y) + abs(p.z-q.z)
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
