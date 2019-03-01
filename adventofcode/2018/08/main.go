package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
	fmt.Println(metadataSum(readTree()))
}

func part2() {
	fmt.Println(nodeValue(readTree()))
}

type node struct {
	children []*node
	metadata []int
}

// Reads a tree from stdin.
func readTree() *node {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	fields := strings.Fields(string(bytes))
	xs := make([]int, len(fields))
	for i, s := range fields {
		xs[i], _ = strconv.Atoi(s)
	}
	root, _ := parseNode(xs)
	return root
}

// Returns the next node in xs, plus the number of integers consumed.
func parseNode(xs []int) (root *node, consumed int) {
	n := &node{
		children: make([]*node, xs[0]),
		metadata: make([]int, xs[1]),
	}
	j := 2
	for i := 0; i < len(n.children); i++ {
		child, advance := parseNode(xs[j:])
		n.children[i] = child
		j += advance
	}
	for i := 0; i < len(n.metadata); i++ {
		n.metadata[i] = xs[j]
		j++
	}
	return n, j
}

// Recursively sums all metadata in the tree.
func metadataSum(n *node) int {
	sum := 0
	for _, c := range n.children {
		sum += metadataSum(c)
	}
	for _, v := range n.metadata {
		sum += v
	}
	return sum
}

// Computes the value of a node (per the instructions in part two).
func nodeValue(n *node) int {
	value := 0
	if len(n.children) == 0 {
		for _, v := range n.metadata {
			value += v
		}
		return value
	}
	for _, i := range n.metadata {
		if i >= 1 && i <= len(n.children) {
			value += nodeValue(n.children[i-1])
		}
	}
	return value
}
