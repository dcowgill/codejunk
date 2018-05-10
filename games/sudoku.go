package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

type Row uint8
const (
	R1 Row = iota
	R2
	R3
	R4
	R5
	R6
	R7
	R8
	R9
)

type Col uint8
const (
	C1 Col = iota
	C2
	C3
	C4
	C5
	C6
	C7
	C8
	C9
)

type Square struct {
	r Row
	c Col
}

func (s Square) String() string {
	return fmt.Sprintf("%c%c", 'A' + s.r, '1' + s.c)
}

func (s Square) Offset() int {
	return int(s.r)*81 + int(s.c)*9
}

func buildSquares(rows []Row, cols []Col) (squares []Square) {
	for _, r := range rows {
		for _, c := range cols {
			squares = append(squares, Square{r,c})
		}
	}
	return squares
}

func unitContainsSquare(a []Square, s Square) bool {
	for _, s2 := range a {
		if s2 == s {
			return true
		}
	}
	return false
}

func filterUnits(predicate func (unit []Square) bool, units [][]Square) (result [][]Square) {
	for _, u := range units {
		if predicate(u) {
			result = append(result, u)
		}
	}
	return result
}

func keys(m map[Square]bool) (ks []Square) {
	for k, _ := range m {
		ks = append(ks, k)
	}
	return ks
}

var rows = []Row{R1, R2, R3, R4, R5, R6, R7, R8, R9}
var cols = []Col{C1, C2, C3, C4, C5, C6, C7, C8, C9}
var squares []Square = buildSquares(rows, cols)
var unitlist [][]Square = func() (unitlist [][]Square) {
	for _, r := range rows {
		unitlist = append(unitlist, buildSquares([]Row{r}, cols))
	}
	for _, c := range cols {
		unitlist = append(unitlist, buildSquares(rows, []Col{c}))
	}
	for _, rs := range [][]Row{{R1,R2,R3}, {R4,R5,R6}, {R7,R8,R9}} {
		for _, cs := range [][]Col{{C1,C2,C3}, {C4,C5,C6}, {C7,C8,C9}} {
			unitlist = append(unitlist, buildSquares(rs, cs))
		}
	}
	return unitlist
}()
var units map[Square][][]Square = func () (units map[Square][][]Square) {
	units = make(map[Square][][]Square)
	for _, s := range squares {
		units[s] = filterUnits(
			func (unit []Square) bool { return unitContainsSquare(unit, s) },
			unitlist)
	}
	return units
}()
var peers map[Square][]Square = func () (peers map[Square][]Square) {
	peers = make(map[Square][]Square)
	for _, s := range squares {
		m := make(map[Square]bool)
		for _, u := range units[s] {
			if unitContainsSquare(u, s) {
				for _, s2 := range u {
					if s2 != s {
						m[s2] = true
					}
				}
			}
		}
		peers[s] = keys(m)
	}
	return peers
}()

const VALUES_LENGTH = 9*9*9 // 9x9 squares, 9 digits per square
type Values [VALUES_LENGTH]bool

// Compute offset into Values array of the given square & digit.
func squareDigitIndex(s Square, d int) int {
	return s.Offset() + d - 1
}

func newValues() *Values {
	v := new(Values)
	for i := 0; i < len(v); i++ {
		v[i] = true
	}
	return v
}

func (src* Values) clone() *Values {
	dst := new(Values)
	// FIXME: use copy(dst, src)?
	for i := 0; i < len(src); i++ {
		dst[i] = src[i]
	}
	return dst
}

func (v *Values) digits(s Square) (rv []int) {
	offset := s.Offset()
	for i := 0; i < 9; i++ {
		if v[i + offset] {
			rv = append(rv, i+1)
		}
	}
	return rv
}

func (v *Values) contains(s Square, d int) bool {
	return v[squareDigitIndex(s, d)]
}

func (v *Values) erase(s Square, d int) bool {
	i := squareDigitIndex(s, d)
	if v[i] {
		v[i] = false
		return true
	}
	return false
}

func (v *Values) count(s Square) int {
	offset := s.Offset()
	c := 0
	for i := 0; i < 9; i++ {
		if v[i + offset] {
			c++
		}
	}
	return c
}

func assign(values *Values, s Square, d int) *Values {
	for _, d2 := range values.digits(s) {
		if d2 != d {
			if eliminate(values, s, d2) == nil {
				return nil
			}
		}
	}
	return values
}

func eliminate(values *Values, s Square, d int) *Values {
	if !values.erase(s, d) {
		return values // already eliminated
	}
	// (1) If a square s is reduced to one value d2, then eliminate d2 from the peers.
	ndigits := values.count(s)
	if (ndigits == 0) {
		return nil // contradiction
	} else if (ndigits == 1) {
		d2 := values.digits(s)[0]
		for _, s2 := range peers[s] {
			if eliminate(values, s2, d2) == nil {
				return nil
			}
		}
	}
	// (2) If a unit u is reduced to only one place for a value d, then put it there.
	for _, u := range units[s] {
		var dplace Square
		flag := 0 // ok, it's gross
		for _, s2 := range u {
			if values.contains(s2, d) {
				if flag == 0 {
					flag = 1
					dplace = s2
				} else {
					flag = 2
					break
				}
			}
		}
		if flag == 0 {
			return nil // contradiction
		} else if flag == 1 {
			if assign(values, dplace, d) == nil {
				return nil
			}
		}
	}
	return values
}

func search(values *Values) *Values {
	if values == nil {
		return nil // failed earlier
	}
	var minSquare Square
	minDigits := 10
	for _, s := range squares {
		if c := values.count(s); c > 1 && c < minDigits {
			minSquare = s
			minDigits = c
		}
	}
	if minDigits > 9 {
		return values // done! no unsolved squares remain
	}
	for _, d := range values.digits(minSquare) {
		if result := search(assign(values.clone(), minSquare, d)); result != nil {
			return result
		}
	}
	return nil
}

func solved(values *Values) bool {
	if values == nil {
		return false
	}
	for _, s := range squares {
		if values.count(s) != 1 {
			return false
		}
	}
	return true
}

func solve(grid string) *Values {
	v := search(parseGrid(grid))
	if !solved(v) {
		panic("Failed to solve a puzzle!")
	}
	return v
}

func gridValues(grid string) (values map[Square]int) {
	var digits []int
	for _, ch := range grid {
		if '0' <= ch && ch <= '9' {
			digits = append(digits, int(ch - '0'))
		} else if ch == '.' {
			digits = append(digits, 0)
		}
	}
	if len(digits) != 81 {
		panic(fmt.Sprintf("unexpected number of digits in grid: %d", len(digits)))
	}
	values = make(map[Square]int)
	for i, s := range squares {
		values[s] = digits[i]
	}
	return values
}

func parseGrid(grid string) *Values {
	values := newValues()
	for s, d := range gridValues(grid) {
		if d != 0 {
			assign(values, s, d)
		}
	}
	return values
}

func fromFile(filename string) (grids []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		grids = append(grids, string(line))
	}
	return grids
}

func test() {
	if len(squares) != 81 {
		panic(fmt.Sprintf("unexpected number of squares: %d", len(squares)))
	}
	if len(unitlist) != 27 {
		panic(fmt.Sprintf("unexpected number of units: %d", len(unitlist)))
	}
	for _, s := range squares {
		if len(units[s]) != 3 {
			panic(fmt.Sprintf("unexpected number of units for square %s: %d", s, len(units[s])))
		}
	}
	for _, s := range squares {
		if len(peers[s]) != 20 {
			panic(fmt.Sprintf("unexpected number of units for square %s: %d", s, len(peers[s])))
		}
	}
	fmt.Printf("All tests pass.\n")
}

func printSolution(values *Values) {
	if !solved(values) {
		panic("showSolution expects a solved puzzle.")
	}
	for _, s := range squares {
		for _, d := range values.digits(s) {
			fmt.Printf("%d", d)
		}
	}
	fmt.Printf("\n")
}

var inputFilename = flag.String("input", "", "input filename")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var nparallel = flag.Int("nparallel", 1, "parallelism")

func parallelSolve(puzzles []string, n int) (result []*Values) {
	solver := func (from chan string, to chan *Values) {
		for g := range from {
			to <- solve(g)
		}
	}
	cp := make(chan string, len(puzzles))
	cs := make(chan *Values, len(puzzles))
	for i := 0; i < n; i++ {
		go solver(cp, cs)
	}
	for _, p := range puzzles {
		cp <- p
	}
	for i := 0; i < len(puzzles); i++ {
		result = append(result, <-cs)
	}
	close(cp)
	close(cs)
	return result
}

func main() {
	test()

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	if *inputFilename == "" {
		log.Fatal("-input is required")
	}
	runtime.GOMAXPROCS(*nparallel)
	grids := fromFile(*inputFilename)
	t1 := time.Now()
	//----------
	solutions := parallelSolve(grids, *nparallel)
	//----------
	elapsed := time.Since(t1)
	for _, v := range solutions {
		printSolution(v)
	}
	mean := elapsed.Seconds() / float64(len(grids))
	hz := float64(len(grids)) / elapsed.Seconds()
	fmt.Printf("Solved %d grids (avg %0.2f secs, %0.2f Hz).\n", len(grids), mean, hz)
}
