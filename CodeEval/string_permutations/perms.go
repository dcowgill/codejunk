package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

type Bytes []byte

func (a Bytes) Len() int           { return len(a) }
func (a Bytes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Bytes) Less(i, j int) bool { return a[i] < a[j] }

// Knuth; algorithm L.
func generateLexicographicPermutations(a []byte, visit func([]byte)) {
	for {
		// L1 [Visit.]
		visit(a)

		// L2 [Find j.]
		j := len(a) - 2
		for j >= 0 && a[j] >= a[j+1] {
			j--
		}
		if j < 0 {
			return
		}

		// L3 [Increase Aj.]
		l := len(a) - 1
		for a[j] >= a[l] {
			l--
		}
		a[j], a[l] = a[l], a[j]

		// L4 [Reverse Aj+1 ... An.]
		for k, l := j+1, len(a)-1; k < l; k, l = k+1, l-1 {
			a[k], a[l] = a[l], a[k]
		}
	}
}

func factorial(n int) int {
	f := 1
	for n > 0 {
		f *= n
		n--
	}
	return f
}

func printPermutations(s string) {
	bytes := Bytes(s)
	sort.Sort(bytes)
	var (
		i = 0
		n = factorial(len(bytes))
	)
	generateLexicographicPermutations(bytes, func(a []byte) {
		os.Stdout.Write(a)
		i++
		if i < n {
			os.Stdout.WriteString(",")
		}
	})
	os.Stdout.WriteString("\n")
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		testcase := scanner.Text()
		if len(testcase) != 0 {
			printPermutations(testcase)
		}
	}
}
