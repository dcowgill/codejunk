package main

import (
	"bytes"
	"container/list"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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
	node := readInput(os.Stdin) // parse the input
	world := newWorld()         // create a blank world
	node.visit(world, Point{})  // process the parsed AST
	const threshold = 1000
	maxDist, numThreshold := farthestRoomDistance(world, threshold)
	fmt.Printf("distance to farthest room = %d\n", maxDist)
	fmt.Printf("rooms at least %d doors away = %d\n", threshold, numThreshold)
}

func part2() {
	part1()
}

// Reads and parses the input.
func readInput(r io.Reader) Node {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	data = bytes.Trim(data, "^$ \n") // remove extraneous symbols
	node, i := parseSequence(data, 0)
	if i < len(data) {
		panic(fmt.Sprintf("parse error at offset %d", i))
	}
	return node
}

// Parses a sequence of steps, which may be directions or parenthesized branches.
// E.g. "NNNNSSSEEW(...)NNNN(...)SSS(...)EEW".
func parseSequence(data []byte, i int) (Node, int) {
	var seq SequenceNode
	for i < len(data) {
		switch x := data[i]; x {
		case '(':
			node, j := parseBranch(data, i)
			seq = append(seq, node)
			i = j + 1
		case 'N', 'S', 'E', 'W':
			seq = append(seq, DirectionNode(charToDirection(x)))
			i++
		default:
			return seq, i
		}
	}
	return seq, i
}

// Parses a set of branches. Expects data[i] to be an opening parenthesis, which
// must be balanced by a closing parenthesis. E.g. "(NS|EW(NN|SS)||S)".
func parseBranch(data []byte, i int) (Node, int) {
	if data[i] != '(' {
		panic(fmt.Sprintf("internal error at offset %d: expected '(', got %c", i, data[i]))
	}
	i++
	var br BranchNode
	for i < len(data) {
		node, j := parseSequence(data, i)
		if j > i {
			br = append(br, node)
		}
		switch data[j] {
		case ')':
			return br, j
		case '|':
			i = j + 1
		default:
			panic(fmt.Sprintf("parse error at offset %d: expected ')' or '|', got %c (%[1]d)", j, data[j]))
		}
	}
	panic("unexpected end of input")
}

type Direction Point

var (
	north = Direction{0, -1}
	south = Direction{0, +1}
	east  = Direction{+1, 0}
	west  = Direction{-1, 0}
)

func charToDirection(b byte) Direction {
	switch b {
	case 'N':
		return north
	case 'S':
		return south
	case 'E':
		return east
	case 'W':
		return west
	}
	panic("invalid direction char")
}

type Point struct {
	x, y int
}

func (p Point) move(d Direction) Point {
	return Point{p.x + d.x, p.y + d.y}
}

type World struct {
	vertmap  map[Point]int
	vertices []Point
	edges    [][]int
}

func newWorld() *World {
	w := &World{vertmap: make(map[Point]int)}
	w.addVertex(Point{0, 0}) // ensure Point {0, 0} has ID = 0
	return w
}

func (w *World) addVertex(p Point) int {
	if id, ok := w.vertmap[p]; ok {
		return id
	}
	id := len(w.vertices)
	w.vertices = append(w.vertices, p)
	w.vertmap[p] = id
	return id
}

func (w *World) addEdge(p, q Point) {
	pid := w.addVertex(p)
	qid := w.addVertex(q)
	for len(w.edges) <= pid || len(w.edges) <= qid {
		w.edges = append(w.edges, nil)
	}
	addToSet(&w.edges[pid], qid)
	addToSet(&w.edges[qid], pid)
}

func addToSet(a *[]int, x int) {
	for _, y := range *a {
		if x == y {
			return
		}
	}
	*a = append(*a, x)
}

func (w *World) hasVertex(p Point) bool {
	_, ok := w.vertmap[p]
	return ok
}

func (w *World) hasEdge(p, q Point) bool {
	pid, pok := w.vertmap[p]
	qid, qok := w.vertmap[q]
	if !pok || !qok || len(w.edges) <= pid {
		return false
	}
	for _, id := range w.edges[pid] {
		if id == qid {
			return true
		}
	}
	return false
}

func (w *World) neighbors(p Point) []Point {
	id, ok := w.vertmap[p]
	if !ok || len(w.edges) <= id {
		return nil
	}
	qs := make([]Point, len(w.edges[id]))
	for i, id2 := range w.edges[id] {
		qs[i] = w.vertices[id2]
	}
	return qs
}

type Node interface {
	visit(w *World, p Point) Point
}

type DirectionNode Direction

func (n DirectionNode) visit(w *World, p Point) Point {
	q := p.move(Direction(n))
	w.addEdge(p, q)
	return q
}

type SequenceNode []Node

func (n SequenceNode) visit(w *World, p Point) Point {
	for _, node := range n {
		p = node.visit(w, p)
	}
	return p
}

type BranchNode []Node

func (n BranchNode) visit(w *World, p Point) Point {
	for _, node := range n {
		node.visit(w, p)
	}
	return p
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

// Returns two values: the distance to the node the most steps from the origin
// and the number of nodes that are at least threshold steps away.
func farthestRoomDistance(world *World, threshold int) (maxDist, numThreshold int) {
	// Straightforward use of BFS.
	var (
		origin   = Point{0, 0}
		frontier = list.New()
		visited  = map[Point]bool{origin: true}
	)
	type Item struct {
		point Point
		dist  int
	}
	frontier.PushBack(Item{origin, 0})
	for frontier.Len() > 0 {
		elem := frontier.Front()
		frontier.Remove(elem)
		item := elem.Value.(Item)
		maxDist = maxInt(maxDist, item.dist)
		if item.dist >= threshold {
			numThreshold++
		}
		for _, neighbor := range world.neighbors(item.point) {
			if !visited[neighbor] {
				frontier.PushBack(Item{neighbor, item.dist + 1})
				visited[neighbor] = true
			}
		}
	}
	return maxDist, numThreshold
}
