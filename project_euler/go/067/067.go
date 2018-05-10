/*

By starting at the top of the triangle below and moving to adjacent numbers on
the row below, the maximum total from top to bottom is 23.

3
7 4
2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom in triangle.txt (right click and
'Save Link/Target As...'), a 15K text file containing a triangle with
one-hundred rows.

NOTE: This is a much more difficult version of Problem 18. It is not possible
to try every route to solve this problem, as there are 299 altogether! If you
could check one trillion (1012) routes every second it would take over twenty
billion years to check them all. There is an efficient algorithm to solve it.
;o)

*/

package p067

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func search(triangle [][]int) int {
	type point struct{ x, y int }
	var (
		nrows = len(triangle)
		best  = make(map[point]int)
		dfs   func(x, y, v0 int) int
	)
	dfs = func(x, y, v0 int) int {
		v1 := v0 + triangle[x][y]
		if x == nrows-1 {
			return v1
		}

		key := point{x, y}
		if best[key] >= v1 {
			return 0
		}
		best[key] = v1

		l := dfs(x+1, y, v1)
		r := dfs(x+1, y+1, v1)
		if l > r {
			return l
		}
		return r
	}
	return dfs(0, 0, 0)
}

func solve() int {
	triangle, err := readTriangleFile("067.txt")
	if err != nil {
		log.Fatal(err)
	}
	return search(triangle)
}

func readTriangleFile(path string) (triangle [][]int, err error) {
	var (
		matchInts = regexp.MustCompile("\\d+")
		lines     []string
	)

	lines, err = readLines(path)
	if err != nil {
		return
	}

	for _, line := range lines {
		var (
			row []int
			n   int
		)
		for _, s := range matchInts.FindAllString(line, -1) {
			n, err = strconv.Atoi(s)
			if err != nil {
				return
			}
			row = append(row, n)
		}
		triangle = append(triangle, row)
	}
	err = validateTriangle(triangle)
	return
}

func validateTriangle(triangle [][]int) error {
	if triangle == nil || len(triangle) == 0 {
		return fmt.Errorf("empty triangle")
	}
	prev := 0
	for i, row := range triangle {
		if len(row) != prev+1 {
			return fmt.Errorf("invalid triangle (row %d)", i+1)
		}
		prev = len(row)
	}
	return nil
}

func readLines(filename string) ([]string, error) {
	fp, err := os.Open(filename)
	defer fp.Close()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(fp)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
