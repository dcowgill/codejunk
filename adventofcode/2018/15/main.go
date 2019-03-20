package main

import (
	"bufio"
	"container/heap"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
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
	numRounds := runUntilDone(world)
	fmt.Printf("final map:\n\n%s\n", world)
	fmt.Printf("rounds = %d\n", numRounds)
	fmt.Printf("HP remaining: %d\n", world.totalHP())
	fmt.Printf("answer: %d\n", numRounds*world.totalHP())
}

func part2() {
	world := readInput(os.Stdin)
	numElvesAtStart := world.numUnitsByKind(elfKind)
	for bonus := 0; ; bonus++ {
		world := world.clone()
		world.elfBonus = bonus
		numRounds := runUntilDone(world)
		if world.numUnitsByKind(elfKind) != numElvesAtStart {
			continue // one or more elves were killed; need more bonus damage
		}
		fmt.Printf("final map:\n\n%s\n", world)
		fmt.Printf("bonus damage = %d\n", bonus)
		fmt.Printf("rounds = %d\n", numRounds)
		fmt.Printf("HP remaining: %d\n", world.totalHP())
		fmt.Printf("answer: %d\n", numRounds*world.totalHP())
		return
	}
}

// Never run a simulation for more than this many rounds.
const maxTestRounds = 100000

func runUntilDone(world *World) (numRounds int) {
	for world.step() {
		numRounds++
		if numRounds > maxTestRounds {
			panic("simulation ran for too many rounds")
		}
	}
	return numRounds
}

func readInput(r io.Reader) *World {
	var (
		scanner     = bufio.NewScanner(r)
		points      []Point
		cells       []*Cell
		y           = 0
		prevLineLen = -1
	)
	for scanner.Scan() {
		line := scanner.Text()
		if prevLineLen < 0 {
			prevLineLen = len(line)
		} else if len(line) != prevLineLen {
			panic(fmt.Sprintf("expected %d bytes in line, got %d: %q", prevLineLen, len(line), line))
		}
		for x := 0; x < len(line); x++ {
			p := Point{x, y}
			var c Cell
			switch line[x] {
			case '#':
				c.isWall = true
			case 'E':
				c.unit = newUnit(p, elfKind)
			case 'G':
				c.unit = newUnit(p, goblinKind)
			case '.':
			default:
				panic(fmt.Sprintf("unexpected cell symbol: %q", line[x]))
			}
			points = append(points, p)
			cells = append(cells, &c)
		}
		y++
	}
	return newWorld(points, cells)
}

// Point is a coordinate position.
type Point struct{ x, y int }

// Up, down, left, or right.
type Direction struct {
	dir Point // vector of length one
	seq int   // lower means higher priority
}

var (
	up    = Direction{Point{0, -1}, 1}
	left  = Direction{Point{-1, 0}, 2}
	right = Direction{Point{+1, 0}, 3}
	down  = Direction{Point{0, +1}, 4}
)

func (d Direction) less(d2 Direction) bool {
	return d.seq < d2.seq
}

func (d Direction) move(p Point) Point {
	return Point{x: p.x + d.dir.x, y: p.y + d.dir.y}
}

// Directions are listed in order of preference.
var directions = []Direction{up, left, right, down}

// UnitKind identifies a type of unit, i.e. elf or goblin.
type UnitKind byte

const (
	elfKind    UnitKind = 'E'
	goblinKind UnitKind = 'G'
)

func (k UnitKind) enemy() UnitKind {
	if k == elfKind {
		return goblinKind
	}
	return elfKind
}

// Unit represents an actual elf or goblin in the world.
type Unit struct {
	pos  Point
	kind UnitKind
	hp   int
}

func newUnit(pos Point, kind UnitKind) *Unit {
	return &Unit{pos: pos, kind: kind, hp: 200}
}

func (u *Unit) clone() *Unit {
	if u == nil {
		return nil
	}
	copy := *u
	return &copy
}

func (u *Unit) isDead() bool {
	return u.hp <= 0
}

// Cell describes a square in the world.
// It may be either a wall, a unit, or an empty space.
type Cell struct {
	unit   *Unit
	isWall bool
}

func (c *Cell) clone() *Cell {
	return &Cell{unit: c.unit.clone(), isWall: c.isWall}
}

func (c *Cell) isEmpty() bool {
	return !c.isWall && c.unit == nil
}

func (c *Cell) containsUnit(k UnitKind) bool {
	return c.unit != nil && c.unit.kind == k
}

// World represents the state of the simulation.
type World struct {
	cells     [][]*Cell // row-major: [row][col], e.g. [y][x]
	units     []*Unit
	liveCount map[UnitKind]int
	elfBonus  int
}

func newWorld(points []Point, cells []*Cell) *World {
	// Determine the size of the world.
	var max Point
	for _, p := range points {
		max.x = maxInt(max.x, p.x)
		max.y = maxInt(max.y, p.y)
	}
	// Allocate the world cells.
	world := &World{
		liveCount: make(map[UnitKind]int),
	}
	world.cells = make([][]*Cell, max.y+1)
	for y := range world.cells {
		world.cells[y] = make([]*Cell, max.x+1)
	}
	// Initialize the cells.
	for i, p := range points {
		world.cells[p.y][p.x] = cells[i]
	}
	// Initialize the units.
	for _, c := range cells {
		if c.unit != nil {
			world.units = append(world.units, c.unit)
			world.liveCount[c.unit.kind]++
		}
	}
	return world
}

// Formats the world as a 2D map grid.
func (world *World) String() string {
	var b strings.Builder
	for _, row := range world.cells {
		for _, c := range row {
			var r byte = '?'
			switch {
			case c.isWall:
				r = '#'
			case c.unit == nil:
				r = '.'
			case c.unit.kind == elfKind:
				r = 'E'
			case c.unit.kind == goblinKind:
				r = 'G'
			}
			b.WriteByte(r)
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func (world *World) clone() *World {
	copy := &World{}
	copy.cells = make([][]*Cell, len(world.cells))
	copy.units = make([]*Unit, 0, len(world.units))
	for i, row := range world.cells {
		copy.cells[i] = make([]*Cell, len(row))
		for j, cell := range row {
			cell := cell.clone() // shadow
			copy.cells[i][j] = cell
			if cell.unit != nil {
				copy.units = append(copy.units, cell.unit)
			}
		}
	}
	copy.liveCount = make(map[UnitKind]int, len(world.liveCount))
	for k, v := range world.liveCount {
		copy.liveCount[k] = v
	}
	return copy
}

func (world *World) cell(p Point) *Cell {
	return world.cells[p.y][p.x]
}

func (world *World) numUnitsByKind(k UnitKind) int {
	return world.liveCount[k]
}

func (world *World) totalHP() int {
	sum := 0
	for _, u := range world.units {
		if u.hp > 0 {
			sum += u.hp
		}
	}
	return sum
}

// Advances the simulation by one round. Returns true if the round was
// completed, implying that the simulation should continue.
func (world *World) step() bool {
	world.reapDeadUnits()
	sort.Sort(byInitiative(world.units))

	visited := make(map[Point]bool)
nextUnit:
	for _, u := range world.units {
		if u.isDead() {
			continue // units may be killed during the step; skip them
		}
		// If no enemies remain, the simulation halts immediately.
		enemyKind := u.kind.enemy()
		if world.numUnitsByKind(enemyKind) == 0 {
			return false
		}
		// Clear the visited-square set.
		for k := range visited {
			delete(visited, k)
		}
		// Units first try to attack an adjacent enemy.
		if world.tryAttack(u) {
			continue // attack made; turn is over
		}
		// Search for a reachable square that is adjacent to any enemy.
		frontier := newFrontier(world)
		for _, dir := range directions {
			if adj := dir.move(u.pos); world.cell(adj).isEmpty() {
				visited[adj] = true
				heap.Push(frontier, &FrontierItem{
					pos:  adj,
					dir:  dir,
					dist: 1,
				})
			}
		}
		for frontier.Len() > 0 {
			curr := heap.Pop(frontier).(*FrontierItem)
			for _, dir := range directions {
				dst := dir.move(curr.pos)
				cell := world.cell(dst)
				if cell.containsUnit(enemyKind) {
					// Found an enemy unit; immediately move
					// along the path indicated by the
					// frontier item. Also: if the total
					// path length is one, we must be next
					// to an enemy, so attack immediately.
					world.moveUnit(u, curr.dir)
					if curr.dist == 1 {
						world.tryAttack(u)
					}
					continue nextUnit
				}
				// If the cell is passable and we haven't seen
				// it before, add it to the frontier.
				if cell.isEmpty() && !visited[dst] {
					visited[dst] = true
					heap.Push(frontier, &FrontierItem{
						pos:  dst,
						dir:  curr.dir,
						dist: curr.dist + 1,
					})
				}

			}
		}
	}
	return true
}

func (world *World) moveUnit(u *Unit, dir Direction) {
	dst := dir.move(u.pos)
	curCell := world.cell(u.pos)
	dstCell := world.cell(dst)
	if !dstCell.isEmpty() {
		panic("tried to move into non-empty cell")
	}
	curCell.unit, dstCell.unit = nil, u
	u.pos = dst
}

const baseAttackPower = 3

// Attacks an adjacent enemy minion, if one exists.
// Returns true if an attack was made.
func (world *World) tryAttack(u *Unit) bool {
	enemyKind := u.kind.enemy()
	var target *Cell
	for _, dir := range directions {
		cell := world.cell(dir.move(u.pos))
		if cell.containsUnit(enemyKind) && (target == nil || cell.unit.hp < target.unit.hp) {
			target = cell
		}
	}
	if target == nil {
		return false // no adjacent enemy
	}
	attackPower := baseAttackPower
	if u.kind == elfKind {
		attackPower += world.elfBonus
	}
	target.unit.hp -= attackPower
	if target.unit.hp <= 0 {
		world.processUnitDeath(target)
	}
	return true
}

func (world *World) processUnitDeath(cell *Cell) {
	for _, u := range world.units {
		if u == cell.unit && u.isDead() {
			world.liveCount[u.kind]--
			cell.unit = nil
			return
		}
	}
	panic("unit not found")
}

func (world *World) reapDeadUnits() {
	numDead := 0
	for _, u := range world.units {
		if u.isDead() {
			numDead++
		}
	}
	if numDead == 0 {
		return
	}
	j := 0
	for i, u := range world.units {
		if !u.isDead() {
			world.units[j] = world.units[i]
			j++
		}
	}
	world.units = world.units[:j]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Sorts units into the order in which the move each turn.
type byInitiative []*Unit

func (a byInitiative) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byInitiative) Len() int      { return len(a) }
func (a byInitiative) Less(i, j int) bool {
	switch {
	case a[i].pos.y < a[j].pos.y:
		return true
	case a[i].pos.y > a[j].pos.y:
		return false
	default:
		return a[i].pos.x < a[j].pos.x
	}
}

// FrontierItem is an item on the frontier heap.
type FrontierItem struct {
	pos  Point     // position being explored
	dir  Direction // first step along this path
	dist int       // distance to pos
}

// Frontier is a BFS heap.
type Frontier struct {
	world *World
	heap  []*FrontierItem
}

func newFrontier(world *World) *Frontier {
	return &Frontier{world: world}
}

func (f *Frontier) Len() int           { return len(f.heap) }
func (f *Frontier) Swap(i, j int)      { f.heap[i], f.heap[j] = f.heap[j], f.heap[i] }
func (f *Frontier) Push(x interface{}) { f.heap = append(f.heap, x.(*FrontierItem)) }
func (f *Frontier) Pop() interface{} {
	old := f.heap
	n := len(old)
	x := old[n-1]
	f.heap = old[0 : n-1]
	return x
}
func (f *Frontier) Less(i, j int) bool {
	// Closest first, then by destination square.
	switch {
	case f.heap[i].dist < f.heap[j].dist:
		return true
	case f.heap[i].dist > f.heap[j].dist:
		return false
	case f.heap[i].pos.y < f.heap[j].pos.y:
		return true
	case f.heap[i].pos.y > f.heap[j].pos.y:
		return false
	case f.heap[i].pos.x < f.heap[j].pos.x:
		return true
	}
	return false
}
