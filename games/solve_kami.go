// Accepts iPhone screenshots of a Kami gameboard and prints a solution.
// See https://itunes.apple.com/us/app/kami/id710724007
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
	"os"
	"sort"
	"strconv"

	colorful "github.com/lucasb-eyer/go-colorful"
)

const (
	ROWS      = 16
	COLS      = 10
	MAX_COLOR = 10
)

type empty struct{}

// Creates a two-dimensional slice of ints with r rows and c cols,
// ensuringe the rows are in contiguous memory.
func newIntMatrix(r, c int) [][]int {
	mem := make([]int, r*c)
	mat := make([][]int, r)
	for i := 0; i < r; i++ {
		mat[i], mem = mem[:c], mem[c:]
	}
	return mat
}

func processImage(src image.Image, numColors int) *Board {
	img := convertToRGBA(src)
	b := img.Bounds()
	tileSize := b.Dx() / COLS
	m := image.NewRGBA(image.Rect(0, 0, COLS*tileSize, ROWS*tileSize))
	draw.Draw(m, m.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)

	grid := newIntMatrix(ROWS, COLS)
	swatches := extractSwatches(img, numColors)
	for col := 0; col < COLS; col++ {
		for row := 0; row < ROWS; row++ {
			sr := image.Rect(
				b.Min.X+col*tileSize, b.Min.Y+row*tileSize,
				b.Min.X+(col+1)*tileSize, b.Min.Y+(row+1)*tileSize,
			)
			sr = trimRect(sr, 5)
			c, x := nearestSwatch(swatches, averageColor(img.SubImage(sr)))
			dr := image.Rect(col*tileSize, row*tileSize, (col+1)*tileSize, (row+1)*tileSize)
			dr = trimRect(dr, 5)
			draw.Draw(m, dr, &image.Uniform{c}, dr.Min, draw.Src)
			grid[row][col] = x + 1
		}
	}
	savePNG("processed", m)
	findRegions(grid)
	return newBoard(grid)
}

// Shrinks the rectangle by n pixels in each dimension.
func trimRect(r image.Rectangle, n int) image.Rectangle {
	return image.Rect(r.Min.X+n, r.Min.Y+n, r.Max.X-n, r.Max.Y-n)
}

// Processes the swatches section of the image board. Also creates some
// files in /tmp for diagnostic purposes.
func extractSwatches(src *image.RGBA, numColors int) []colorful.Color {
	const (
		W = 400
		H = 75
	)
	var swatches []colorful.Color
	b := src.Bounds()
	sw := W / numColors
	for i := 0; i < numColors; i++ {
		m := src.SubImage(trimRect(image.Rect(b.Min.X+i*sw, b.Max.Y-H, b.Min.X+(i+1)*sw, b.Max.Y), 10))
		swatches = append(swatches, toColorful(averageColor(m)))
		savePNG(strconv.Itoa(i), m) // for debugging
	}
	const dim = 50
	m := image.NewRGBA(image.Rect(0, 0, dim*len(swatches), dim))
	for i, c := range swatches {
		r := image.Rect(i*dim, 0, (i+1)*dim, dim)
		draw.Draw(m, r, &image.Uniform{fromColorful(c)}, image.ZP, draw.Src)
	}
	savePNG("swatches", m) // for debugging
	return swatches
}

// Reports which of the swatches is most similar to c.
func nearestSwatch(swatches []colorful.Color, c color.Color) (color.Color, int) {
	c0 := toColorful(c)
	minDist := math.MaxFloat64
	var best colorful.Color
	bestIndex := -1
	for i, s := range swatches {
		if d := c0.DistanceCIE94(s); d < minDist {
			minDist, best, bestIndex = d, s, i
		}
	}
	return fromColorful(best), bestIndex
}

const outputPrefix = "/tmp/kami_" // see savePNG

// Encodes the image as PNG with the filename outputPrefix+name+".png".
func savePNG(name string, img image.Image) {
	filename := outputPrefix + name + ".png"
	fp, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(fp, img)
	if err != nil {
		log.Fatal(err)
	}
}

// Decodes the specified file as a PNG image.
func loadPNG(filename string) image.Image {
	fp, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	img, err := png.Decode(fp)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func toColorful(c color.Color) colorful.Color {
	c0 := color.RGBAModel.Convert(c).(color.RGBA)
	return colorful.Color{
		R: float64(c0.R) / float64(0xFFFF),
		G: float64(c0.G) / float64(0xFFFF),
		B: float64(c0.B) / float64(0xFFFF),
	}
}

func fromColorful(c colorful.Color) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}

// Reports the average color of an image.
func averageColor(src image.Image) color.Color {
	b := src.Bounds()
	var sum struct{ c, y, m, k float64 }
	n := 0
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c := color.CMYKModel.Convert(src.At(x, y)).(color.CMYK)
			sum.c += float64(c.C)
			sum.m += float64(c.M)
			sum.y += float64(c.Y)
			sum.k += float64(c.K)
			n++
		}
	}
	d := float64(n)
	return color.CMYK{
		C: uint8(sum.c / d),
		M: uint8(sum.m / d),
		Y: uint8(sum.y / d),
		K: uint8(sum.k / d),
	}
}

// Converts an image to the RGBA type.
func convertToRGBA(src image.Image) *image.RGBA {
	b := src.Bounds()
	m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(m, m.Bounds(), src, b.Min, draw.Src)
	return m
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

type Tile struct{ row, col int }

func (t Tile) add(v Tile) Tile {
	return Tile{row: t.row + v.row, col: t.col + v.col}
}

var dirs = []Tile{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}}

func adjacentTiles(t Tile) [4]Tile {
	var a [4]Tile
	for i, d := range dirs {
		a[i] = t.add(d)
	}
	return a
}

type TileSet struct {
	id    int
	color int
	tiles map[Tile]bool
}

func newTileSet(id, color int) *TileSet {
	return &TileSet{
		id:    id,
		color: color,
		tiles: make(map[Tile]bool),
	}
}
func (ts *TileSet) add(t Tile)    { ts.tiles[t] = true }
func (ts *TileSet) remove(t Tile) { delete(ts.tiles, t) }
func (ts *TileSet) pop() Tile {
	var t Tile
	for t = range ts.tiles {
		break
	}
	ts.remove(t)
	return t
}
func (ts *TileSet) contains(t Tile) bool { return ts.tiles[t] }
func (ts *TileSet) empty() bool          { return len(ts.tiles) == 0 }
func (ts *TileSet) topLeft() Tile {
	t0 := Tile{row: ROWS, col: COLS}
	for t1 := range ts.tiles {
		if t1.row < t0.row || (t1.row == t0.row && t1.col < t0.col) {
			t0 = t1
		}
	}
	return t0
}

func findRegions(grid [][]int) map[int]*Region {
	tilesLeft := newTileSet(-1, -1)
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			tilesLeft.add(Tile{r, c})
		}
	}

	var findReachable func(tiles *TileSet, color int, from Tile)
	findReachable = func(tiles *TileSet, color int, from Tile) {
		for _, tile := range adjacentTiles(from) {
			if tilesLeft.contains(tile) && grid[tile.row][tile.col] == color {
				tilesLeft.remove(tile)
				tiles.add(tile)
				findReachable(tiles, color, tile)
			}
		}
	}

	var tileSets []*TileSet
	idSeq := 0
	for !tilesLeft.empty() {
		tile := tilesLeft.pop()
		color := grid[tile.row][tile.col]
		tiles := newTileSet(idSeq, color)
		tiles.add(tile)
		idSeq++
		findReachable(tiles, color, tile)
		tileSets = append(tileSets, tiles)
	}

	adjacent := func(ts1, ts2 *TileSet) bool {
		for a := range ts1.tiles {
			for _, b := range adjacentTiles(a) {
				if ts2.contains(b) {
					return true
				}
			}
		}
		return false
	}

	regions := make(map[int]*Region)
	for i, ts := range tileSets {
		regions[i] = &Region{
			id:    ts.id,
			color: ts.color,
			tile:  ts.topLeft(),
		}
	}

	for i, ts1 := range tileSets {
		r1 := regions[ts1.id]
		for _, ts2 := range tileSets[i+1:] {
			if adjacent(ts1, ts2) {
				r2 := regions[ts2.id]
				r1.neighbors.add(r2.id)
				r2.neighbors.add(r1.id)
			}
		}
	}

	return regions
}

type Region struct {
	id, color int
	tile      Tile
	neighbors intSet
}

func (r *Region) Copy() *Region {
	r1 := *r
	r1.neighbors = r.neighbors.Copy()
	return &r1
}

type Board struct {
	regions map[int]*Region
}

func newBoard(grid [][]int) *Board { return &Board{regions: findRegions(grid)} }

func (b *Board) solved() bool {
	return len(b.regions) == 1
}

func (b *Board) numColors() int {
	seen := make([]bool, MAX_COLOR)
	for _, r := range b.regions {
		seen[r.color] = true
	}
	n := 0
	for _, b := range seen {
		if b {
			n++
		}
	}
	return n
}

func (b *Board) colorsAdjacentToRegion(region *Region) intSet {
	colors := newIntSet(MAX_COLOR)
	for _, neighborID := range region.neighbors {
		neighbor := b.regions[neighborID]
		colors.add(neighbor.color)
	}
	return colors
}

func (b *Board) recolor(regionID, color int) *Board {
	recolored := b.regions[regionID].Copy()
	recolored.color = color

	removed := newIntSet(len(recolored.neighbors))
	for _, id := range recolored.neighbors {
		neighbor := b.regions[id]
		if neighbor.color == color {
			removed.add(id)
		}
	}

	regions := make(map[int]*Region, len(b.regions)-len(removed))
	for _, r := range b.regions {
		if !removed.contains(r.id) {
			regions[r.id] = r
		}
	}
	regions[recolored.id] = recolored

	copyOnWrite := func(r *Region) *Region {
		if regions[r.id] != b.regions[r.id] {
			return r
		}
		c := r.Copy()
		regions[c.id] = c
		return c
	}

	for _, removedID := range removed {
		for _, neighborID := range b.regions[removedID].neighbors {
			if neighborID == recolored.id {
				continue
			}
			if nn, ok := regions[neighborID]; ok {
				nn = copyOnWrite(nn)
				nn.neighbors.remove(removedID)
				nn.neighbors.add(recolored.id)
				recolored.neighbors.add(nn.id)
			}
		}
		recolored.neighbors.remove(removedID)
	}

	return &Board{regions: regions}
}

type intSet []int

func newIntSet(cap int) intSet { return make([]int, 0, cap) }
func (s intSet) Copy() intSet {
	a := make([]int, len(s))
	copy(a, s)
	return a
}
func (s intSet) contains(x int) bool {
	for _, y := range s {
		if x == y {
			return true
		}
	}
	return false
}
func (s *intSet) add(x int) {
	if !s.contains(x) {
		*s = append(*s, x)
	}
}
func (s *intSet) remove(x int) {
	for i, y := range *s {
		if x == y {
			(*s)[i] = (*s)[len(*s)-1]
			*s = (*s)[:len(*s)-1]
			return
		}
	}
}

func getRegions(b *Board) []*Region {
	a := make([]*Region, 0, len(b.regions))
	for _, r := range b.regions {
		a = append(a, r)
	}
	sort.Sort(byNumNeighbors(a))
	return a
}

type byNumNeighbors []*Region

func (a byNumNeighbors) Len() int           { return len(a) }
func (a byNumNeighbors) Less(i, j int) bool { return len(a[i].neighbors) > len(a[j].neighbors) }
func (a byNumNeighbors) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type Move struct {
	tile  Tile
	color int
}

type work struct {
	board     *Board
	regionID  int
	color     int
	movesLeft int
}

func search(b *Board, regionID int, movesLeft int) []Move {
	switch {
	case b.numColors() > movesLeft+1:
		return nil
	case b.solved():
		return []Move{}
	case movesLeft <= 0:
		return nil
	}
	r := b.regions[regionID]
	for _, color := range b.colorsAdjacentToRegion(r) {
		moves := search(b.recolor(r.id, color), r.id, movesLeft-1)
		if moves != nil {
			return append(moves, Move{r.tile, color})
		}
	}
	return nil
}

func workerProcess(c1 <-chan work, c2 chan<- []Move) {
	for work := range c1 {
		newBoard := work.board.recolor(work.regionID, work.color)
		moves := search(newBoard, work.regionID, work.movesLeft-1)
		if moves != nil {
			r := newBoard.regions[work.regionID]
			c2 <- append(moves, Move{r.tile, work.color})
			return
		}
	}
	c2 <- nil
}

func solve(b *Board, maxMoves int, numWorkers int) []Move {
	workChan := make(chan work)
	solutionChan := make(chan []Move)

	// Launch consumers
	for i := 0; i < numWorkers; i++ {
		go workerProcess(workChan, solutionChan)
	}

	// Launch producer
	go func() {
		for _, region := range getRegions(b) {
			colors := b.colorsAdjacentToRegion(region)
			for _, color := range colors {
				fmt.Printf("go region %3d: color %d -> %d\n", region.id, region.color, color)
				workChan <- work{
					board:     b,
					regionID:  region.id,
					color:     color,
					movesLeft: maxMoves,
				}
			}
		}
		close(workChan) // no more work
	}()

	// Wait for a solution
	for i := 0; i < numWorkers; i++ {
		moves := <-solutionChan
		if moves != nil {
			for i, j := 0, len(moves)-1; i < j; i, j = i+1, j-1 {
				moves[i], moves[j] = moves[j], moves[i]
			}
			return moves
		}
	}
	return nil
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 3 {
		fmt.Printf("usage: gokami <ncolors> <nmoves> <filename>\n")
		os.Exit(0)
	}
	numColors := parseInt(args[0])
	numMoves := parseInt(args[1])
	filename := args[2]
	board := processImage(loadPNG(filename), numColors)
	printBoard(board)
	fmt.Println("")
	solution := solve(board, numMoves, 8)
	fmt.Println("")
	for i, move := range solution {
		fmt.Printf("move %2d: (%d, %d) -> %d\n", i+1, move.tile.row, move.tile.col, move.color)
	}
}

func printBoard(b *Board) {
	var ids []int
	for _, r := range b.regions {
		ids = append(ids, r.id)
	}
	sort.Ints(ids)
	for _, id := range ids {
		r := b.regions[id]
		fmt.Printf("region %3d: (%2d, %2d) -> %d %v\n", id, r.tile.row, r.tile.col, r.color, r.neighbors)
	}
}
