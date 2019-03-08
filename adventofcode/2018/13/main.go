package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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
	grid, carts := readInput()
	for {
		collision, ok := step1(grid, carts)
		if !ok {
			fmt.Printf("first collision: %s\n", collision)
			return
		}
	}
}

func part2() {
	grid, carts := readInput()
	for {
		var ok bool
		carts, ok = step2(grid, carts)
		if !ok {
			if final := reapCarts(carts); len(final) > 0 {
				fmt.Printf("final cart position: %s\n", final[0].pos)
			}
			return
		}
	}
}

type TrackType byte

const (
	none         TrackType = ' '
	horizontal   TrackType = '-'
	vertical     TrackType = '|'
	intersection TrackType = '+'
	turnF        TrackType = '/'
	turnB        TrackType = '\\'
)

type Direction byte

const (
	north Direction = '^'
	south Direction = 'v'
	east  Direction = '>'
	west  Direction = '<'
)

type NextTurn int8

const (
	left     NextTurn = 0
	straight NextTurn = 1
	right    NextTurn = 2
)

func (t NextTurn) apply(d Direction) Direction {
	switch t {
	case left:
		switch d {
		case west:
			return south
		case east:
			return north
		case north:
			return west
		case south:
			return east
		}
	case right:
		switch d {
		case west:
			return north
		case east:
			return south
		case north:
			return east
		case south:
			return west
		}
	}
	return d // straight
}

func (t NextTurn) rotate() NextTurn {
	return NextTurn((t + 1) % 3)
}

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func (p Point) north() Point { return Point{p.x, p.y - 1} }
func (p Point) south() Point { return Point{p.x, p.y + 1} }
func (p Point) east() Point  { return Point{p.x + 1, p.y} }
func (p Point) west() Point  { return Point{p.x - 1, p.y} }

func (p Point) advance(dir Direction) Point {
	switch dir {
	case north:
		return p.north()
	case south:
		return p.south()
	case east:
		return p.east()
	case west:
		return p.west()
	}
	panic(fmt.Sprintf("invalid direction: %+v", dir))
}

type Cart struct {
	pos  Point
	dir  Direction
	next NextTurn
	dead bool
}

type TrackPiece struct {
	pos Point
	tt  TrackType
}

type Grid struct {
	tracks [][]TrackType // row major order: [y][x]
}

func newGrid(tps []TrackPiece) *Grid {
	// Determine bounds.
	var max Point
	for _, t := range tps {
		if t.pos.x > max.x {
			max.x = t.pos.x
		}
		if t.pos.y > max.y {
			max.y = t.pos.y
		}
	}
	// Allocate and initialize the grid.
	tracks := make([][]TrackType, max.y+1)
	for y := range tracks {
		tracks[y] = make([]TrackType, max.x+1)
		for x := range tracks[y] {
			tracks[y][x] = none
		}
	}
	// Set the track pieces.
	for _, t := range tps {
		tracks[t.pos.y][t.pos.x] = t.tt
	}
	return &Grid{tracks: tracks}
}

// Reports the track type at the given position.
func (g *Grid) track(p Point) TrackType {
	return g.tracks[p.y][p.x]
}

func readInput() (*Grid, []*Cart) {
	var (
		tracks []TrackPiece
		carts  []*Cart
		y      = 0
	)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for x, b := range line {
			var t TrackType
			p := Point{x, y}
			switch b {
			case '-':
				t = horizontal
			case '|':
				t = vertical
			case '+':
				t = intersection
			case '/':
				t = turnF
			case '\\':
				t = turnB
			case '^':
				t = vertical
				carts = append(carts, &Cart{pos: p, dir: north})
			case 'v':
				t = vertical
				carts = append(carts, &Cart{pos: p, dir: south})
			case '>':
				t = horizontal
				carts = append(carts, &Cart{pos: p, dir: east})
			case '<':
				t = horizontal
				carts = append(carts, &Cart{pos: p, dir: west})
			case ' ':
				t = none
			default:
				panic(fmt.Sprintf("unknown track symbol at (%s): %q", p, b))
			}
			if t != none {
				tracks = append(tracks, TrackPiece{p, t})
			}
		}
		y++
	}
	return newGrid(tracks), carts
}

// Part 1: returns false as soon as there is a collision.
func step1(g *Grid, carts []*Cart) (collision Point, ok bool) {
	sortCarts(carts)
	for i, c := range carts {
		advanceCart(g, c)
		for j, c2 := range carts {
			if i != j && c.pos == c2.pos {
				return c.pos, false
			}
		}
	}
	return collision, true // no carts collided
}

// Part 2: colliding carts are destroyed.
// Returns false if fewer than two carts remain.
func step2(g *Grid, carts []*Cart) (newcarts []*Cart, ok bool) {
	carts = sortCarts(reapCarts(carts))
	numLiveCarts := len(carts)
	for i, c := range carts {
		if c.dead {
			continue // ignore dead carts
		}
		advanceCart(g, c)
		for j, c2 := range carts {
			if i != j && !c2.dead && c.pos == c2.pos {
				c.dead, c2.dead = true, true
				numLiveCarts -= 2
			}
		}
	}
	return carts, numLiveCarts > 1
}

// Removes dead carts.
func reapCarts(cs []*Cart) []*Cart {
	j := 0
	for i, c := range cs {
		if !c.dead {
			cs[j] = cs[i]
			j++
		}
	}
	return cs[:j]
}

// Sorts by initiative.
// Carts move top-to-bottom, then left-to-right.
// So sort first by y, then by x.
func sortCarts(cs []*Cart) []*Cart {
	sort.Slice(cs, func(i, j int) bool {
		switch {
		case cs[i].pos.y < cs[j].pos.y:
			return true
		case cs[i].pos.y > cs[j].pos.y:
			return false
		}
		return cs[i].pos.x < cs[j].pos.x
	})
	return cs
}

// Updates the cart's position, direction, and next turn decision.
func advanceCart(grid *Grid, cart *Cart) {
	cart.pos = cart.pos.advance(cart.dir)
	switch grid.track(cart.pos) {
	case intersection: // i.e. "+"
		cart.dir = cart.next.apply(cart.dir)
		cart.next = cart.next.rotate()
	case turnF: // i.e. "/"
		switch cart.dir {
		case west:
			cart.dir = south
		case east:
			cart.dir = north
		case north:
			cart.dir = east
		case south:
			cart.dir = west
		}
	case turnB: // i.e. "\"
		switch cart.dir {
		case west:
			cart.dir = north
		case east:
			cart.dir = south
		case north:
			cart.dir = west
		case south:
			cart.dir = east
		}
	}
}
