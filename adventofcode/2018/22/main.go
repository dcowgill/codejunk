package main

import (
	"container/heap"
	"flag"
	"fmt"
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
	cs := defaultCaveSystem()
	totalRiskLevel := 0
	for y := cs.origin.y; y <= cs.target.y; y++ {
		for x := cs.origin.x; x <= cs.target.x; x++ {
			totalRiskLevel += cs.regionType(Point{x, y}).riskLevel()
		}
	}
	fmt.Printf("total risk level = %d\n", totalRiskLevel)
}

func part2() {
	fmt.Printf("%+v\n", shortestPathToTarget(defaultCaveSystem()))
}

func defaultCaveSystem() *CaveSystem {
	return newCaveSystem(3198, Point{}, Point{12, 757})
}

type Point struct {
	x, y int
}

func (p Point) move(d Direction) Point {
	return Point{p.x + d.x, p.y + d.y}
}

type Direction Point

var directions = []Direction{
	Direction{0, -1},
	Direction{0, +1},
	Direction{+1, 0},
	Direction{-1, 0},
}

// From the instructions:
//
// If the erosion level modulo 3 is 0, the region's type is rocky.
// If the erosion level modulo 3 is 1, the region's type is wet.
// If the erosion level modulo 3 is 2, the region's type is narrow.
type RegionType int

const (
	rocky  RegionType = 0
	wet    RegionType = 1
	narrow RegionType = 2
)

func (t RegionType) riskLevel() int              { return int(t) }
func erosionLevelToRegionType(el int) RegionType { return RegionType(el % 3) }

type CaveSystem struct {
	depth  int
	origin Point
	target Point

	// Memoizes the result of CaveSystem.erosionLevel().
	erosionLevelCache map[Point]int
}

func newCaveSystem(depth int, origin, target Point) *CaveSystem {
	return &CaveSystem{
		depth:             depth,
		origin:            origin,
		target:            target,
		erosionLevelCache: make(map[Point]int),
	}
}

func (cs *CaveSystem) geologicIndex(p Point) int {
	switch {
	case p == cs.origin || p == cs.target:
		return 0
	case p.y == 0:
		return p.x * 16807
	case p.x == 0:
		return p.y * 48271
	default:
		return cs.erosionLevel(Point{p.x - 1, p.y}) * cs.erosionLevel(Point{p.x, p.y - 1})
	}
}

func (cs *CaveSystem) erosionLevel(p Point) int {
	if el, ok := cs.erosionLevelCache[p]; ok {
		return el
	}
	el := (cs.geologicIndex(p) + cs.depth) % 20183
	cs.erosionLevelCache[p] = el
	return el
}

func (cs *CaveSystem) regionType(p Point) RegionType {
	return erosionLevelToRegionType(cs.erosionLevel(p))
}

type Equipment uint8

const (
	climbingGear Equipment = 1 // 001
	torch        Equipment = 2 // 010
	neither      Equipment = 4 // 100
)

var validEquipment = []Equipment{
	rocky:  climbingGear | torch,
	wet:    climbingGear | neither,
	narrow: torch | neither,
}

type State struct {
	pos       Point     // current position
	minutes   int       // time spent on journey so far
	equipment Equipment // currently equipped tools

	priority int
}

func (st *State) move(dst Point) *State {
	return &State{pos: dst, minutes: st.minutes + 1, equipment: st.equipment}
}

func (st *State) changeEquipment(e Equipment) {
	st.minutes += 7
	st.equipment = e
}

func move(cs *CaveSystem, curr *State, dir Direction) *State {
	dst := curr.pos.move(dir)
	if dst.x < 0 || dst.y < 0 {
		return nil
	}
	st := curr.move(dst)
	validDst := validEquipment[cs.regionType(dst)]
	if validDst&curr.equipment != 0 {
		return st
	}
	st.changeEquipment(validEquipment[cs.regionType(curr.pos)] & validDst)
	return st
}

func shortestPathToTarget(cs *CaveSystem) *State {
	type Key struct {
		pos       Point
		equipment Equipment
	}
	visited := make(map[Key]int) // key -> minutes
	var frontier PriorityQueue
	heap.Push(&frontier, makeItem(&State{equipment: torch}, cs.target))
	for len(frontier) != 0 {
		st := heap.Pop(&frontier).(Item).state
		// If we are at the target...
		if st.pos == cs.target {
			// ...and have a torch equipped, we're done.
			if st.equipment == torch {
				return st
			}
			// Otherwise, equip the torch and continue.
			st2 := *st
			st2.changeEquipment(torch)
			heap.Push(&frontier, makeItem(&st2, cs.target))
			continue
		}
		// Try to move in every legal direction.
		for _, dir := range directions {
			st2 := move(cs, st, dir)
			if st2 == nil {
				continue // cannot move in this direction
			}
			// If this is the fastest route to this position with
			// these tools equipped, add it to the frontier.
			k := Key{st2.pos, st2.equipment}
			best, found := visited[k]
			if st2.minutes < best || !found {
				heap.Push(&frontier, makeItem(st2, cs.target))
				visited[k] = st2.minutes
			}
		}
	}
	return nil // no route found
}

// Item in a priority queue.
type Item struct {
	cost  int
	state *State
}

type PriorityQueue []Item

func (h PriorityQueue) Len() int           { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h PriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PriorityQueue) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *PriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func makeItem(st *State, target Point) Item {
	cost := st.minutes + manhattanDist(st.pos, target)
	if st.equipment != torch {
		cost += 7
	}
	return Item{cost, st}
}
func manhattanDist(p, q Point) int {
	return abs(p.x-q.x) + abs(p.y-q.y)
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
