package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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
	const numSteps = 10
	world := readInput(os.Stdin)
	for i := 0; i < numSteps; i++ {
		world.step()
	}
	lumberyards, trees := world.countLumberyardsAndTrees()
	fmt.Printf("lumberyards=%d, trees=%d, lumberyards*trees=%d\n",
		lumberyards, trees, lumberyards*trees)
}

func part2() {
	const numSteps = 1000 // we hope pattern is apparent before this many steps
	world := readInput(os.Stdin)
	data := make([]int, numSteps)
	for i := 0; i < numSteps; i++ {
		lumberyards, trees := world.countLumberyardsAndTrees()
		data[i] = lumberyards * trees
		world.step()
	}
	// Find the period of the repeating pattern by trying all possible
	// lengths for all possible starting points.
	for n := 1; n < len(data)/2; n++ {
		for i := 0; i < len(data)-2*n; i++ {
			if isPeriodic(data[i:], n) {
				fmt.Printf("period length is %d (starting at offset %d)\n", n, i)
				fmt.Printf("repeating portion:\n")
				fmt.Printf("%+v\n", data[i:i+n])
				const N = 1000000000
				fmt.Printf("after %d steps, answer is %d\n", N, data[i+(N-i)%28])
				return
			}
		}
	}
}

// Reports whether the data in "a" contains repeating patterns of length "n".
func isPeriodic(a []int, n int) bool {
	if len(a) < 2*n {
		panic(fmt.Sprintf("len(a) = %d, n = %d, cannot be periodic", len(a), n))
	}
	init := a[:n]
	for i := n; i < len(a)-n; i += n {
		curr := a[i : i+n]
		if !intSliceEqual(init, curr) {
			return false
		}
	}
	return true
}

// reflect.DeepEqual would be fine, too.
func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}

func readInput(r io.Reader) *World {
	scanner := bufio.NewScanner(r)
	var tiles []Tile
	width := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue // skip blank lines
		}
		if width == 0 {
			width = len(line)
		} else if len(line) != width {
			panic(fmt.Sprintf("mismatched line widths: got %d, expected %d", len(line), width))
		}
		for i := 0; i < width; i++ {
			var tile Tile
			switch line[i] {
			case '.':
				tile = Open
			case '#':
				tile = Lumberyard
			case '|':
				tile = Trees
			default:
				panic(fmt.Sprintf("unrecognized tile: %c", line[i]))
			}
			tiles = append(tiles, tile)
		}
	}
	if len(tiles) == 0 {
		panic("no tiles read")
	}
	return newWorld(tiles, width)
}

type Tile int8

const (
	Open       Tile = '.'
	Lumberyard Tile = '#'
	Trees      Tile = '|'
)

// State of the simulation.
type World struct {
	tiles, aux []Tile
	neighbors  [][]int
	width      int
}

func newWorld(tiles []Tile, width int) *World {
	return &World{
		tiles:     tiles,
		neighbors: makeNeighbors(len(tiles), width),
		width:     width,
	}
}

// Builds an array that links each tile offset to the offsets of its neighbors,
// which are the eight surrounding tiles. For example, given
//
//	neighbors = makeNeighbors(2500, 50)
//
// then the neighbors of tile 1 (coordinates [x:1, y:0]) are
//
//	neighbors[1] // i.e. [0, 2, 49, 50, 51]; 3 of 8 neighbors are off the grid
//
func makeNeighbors(numTiles, width int) [][]int {
	var (
		neighbors = make([][]int, numTiles)
		inBounds  = func(x, y int) bool {
			return 0 <= x && x < width && 0 <= y && y < numTiles/width
		}
		dirs = [][2]int{
			{-1, +0}, // N
			{-1, +1}, // NE
			{+0, +1}, // E
			{+1, +1}, // SE
			{+1, +0}, // S
			{+1, -1}, // SW
			{+0, -1}, // W
			{-1, -1}, // NW
		}
		tileToPoint = func(i int) (x, y int) {
			y = i / width
			x = i - y*width
			return x, y
		}
		pointToTile = func(x, y int) int {
			return y*width + x
		}
	)
	for i := 0; i < numTiles; i++ {
		numNeighbors := 0 // count first to avoid multiple allocs
		x, y := tileToPoint(i)
		for _, d := range dirs {
			if inBounds(x+d[0], y+d[1]) {
				numNeighbors++
			}
		}
		neighbors[i] = make([]int, 0, numNeighbors)
		for _, d := range dirs {
			if x1, y1 := x+d[0], y+d[1]; inBounds(x1, y1) {
				neighbors[i] = append(neighbors[i], pointToTile(x1, y1))
			}
		}
	}
	return neighbors
}

// Reports the number of lumberyard and tree tiles, respectively.
func (w *World) countLumberyardsAndTrees() (lumberyards, trees int) {
	for _, tile := range w.tiles {
		switch tile {
		case Lumberyard:
			lumberyards++
		case Trees:
			trees++
		}
	}
	return lumberyards, trees
}

// Moves the simulation forward one step.
func (w *World) step() {
	// Reports whether the tile at offset i has >=min tiles equal to t.
	hasAtLeastNAdjacent := func(i, min int, t Tile) bool {
		n := 0
		for _, j := range w.neighbors[i] {
			if w.tiles[j] == t {
				n++
				if n >= min {
					return true
				}
			}
		}
		return false
	}
	if w.aux == nil {
		w.aux = make([]Tile, len(w.tiles))
	}
	// Evolve each tile, storing the new values in the "aux" array.
	for i, tile := range w.tiles {
		switch tile {
		case Open:
			if hasAtLeastNAdjacent(i, 3, Trees) {
				tile = Trees
			}
		case Trees:
			if hasAtLeastNAdjacent(i, 3, Lumberyard) {
				tile = Lumberyard
			}
		case Lumberyard:
			if !hasAtLeastNAdjacent(i, 1, Lumberyard) || !hasAtLeastNAdjacent(i, 1, Trees) {
				tile = Open
			}
		}
		w.aux[i] = tile
	}
	// Update the world by swapping the live and auxillary slices.
	w.tiles, w.aux = w.aux, w.tiles
}
