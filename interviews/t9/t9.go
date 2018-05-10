package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	NDIGITS = 8
)

var (
	DIGITS = []int{2, 3, 4, 5, 6, 7, 8, 9}
	KEYPAD = [][]string{
		nil, // 0
		nil, // 1
		[]string{"a", "b", "c"},      // 2
		[]string{"d", "e", "f"},      // 3
		[]string{"g", "h", "i"},      // 4
		[]string{"j", "k", "l"},      // 5
		[]string{"m", "n", "o"},      // 6
		[]string{"p", "q", "r", "s"}, // 7
		[]string{"t", "u", "v"},      // 8
		[]string{"w", "x", "y", "z"}, // 9
	}
)

type prefixMap map[string]bool

// Returns the set of all prefixes of the words in the given dictionary file.
func buildPrefixMapFromFile(path string) prefixMap {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	prefixes := make(prefixMap)
	r := bufio.NewReader(f)
	for {
		// Get next line.
		s, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		// Add all prefixes of the word.
		word := strings.ToLower(strings.TrimSpace(s))
		for i := 1; i <= len(word); i++ {
			prefixes[word[:i]] = true
		}
	}
	return prefixes
}

// Returns the successor set of prefixes derived from appending each of the
// specified digit's letters to parentPrefixes.
func (m prefixMap) nextPrefixes(digit int, parentPrefixes []string) (prefixes []string) {
	for _, p := range parentPrefixes {
		for _, c := range KEYPAD[digit] {
			if s := p + c; m[s] {
				prefixes = append(prefixes, s)
			}
		}
	}
	return
}

// A node in a T9 prefix tree.
type node struct {
	prefixes []string
	children [NDIGITS]*node // 2..9
}

// Builds a tree beginning with the given prefixes. Uses the prefixMap to
// recursively construct children with successively longer prefixes.
func newNode(prefixes []string, m prefixMap) *node {
	if len(prefixes) == 0 {
		prefixes = []string{""}
	}
	n := node{prefixes: prefixes}
	for i, d := range DIGITS {
		if ps := m.nextPrefixes(d, prefixes); len(ps) != 0 {
			n.children[i] = newNode(ps, m)
		}
	}
	return &n
}

// Returns the individual digits of a number, from most- to least-significant.
// For example, given x=12345, returns []int{1,2,3,4,5}.
func digitsOf(x int) []int {
	var a []int
	for x != 0 {
		a = append(a, x%10)
		x /= 10
	}
	for i, j := 0, len(a)-1; i < j; {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	return a
}

// Returns the set of prefixes for the keypad input x.
func (n *node) lookup(x int) []string {
	for _, d := range digitsOf(x) {
		if d >= 2 && d <= 9 { // ignore 0 and 1
			if n = n.children[d-2]; n == nil {
				return nil
			}
		}
	}
	return n.prefixes
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("wrong number of arguments")
	}
	m := buildPrefixMapFromFile(os.Args[1])
	t9 := newNode(nil, m)
	fmt.Println(len(m))
	for {
		var digits int
		_, err := fmt.Scanf("%d\n", &digits)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(strings.Join(t9.lookup(digits), ","))
	}
}
