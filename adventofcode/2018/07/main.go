package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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
	ins := readInstructions()
	var avail []string // steps with zero dependencies remaining
	ins.forEachParent(func(step string) {
		if ins.numParents(step) == 0 {
			avail = append(avail, step)
		}
	})
	var sequence []string
	for len(avail) > 0 {
		step := popMin(&avail)
		ins.forEachChildOf(step, func(child string) {
			ins.removeParentFromChild(step, child)
			if ins.numParents(child) == 0 {
				avail = append(avail, child)
			}
		})
		sequence = append(sequence, step)
	}
	fmt.Println(strings.Join(sequence, ""))
}

func part2() {
	const numWorkers = 5
	duration := func(step string) int {
		return int([]rune(step)[0]-'A') + 61
	}

	ins := readInstructions()
	var avail []string // steps with zero dependencies remaining
	ins.forEachParent(func(step string) {
		if ins.numParents(step) == 0 {
			avail = append(avail, step)
		}
	})

	type task struct {
		step string
		done int
	}

	// Removes from the slice the earliest ending task and returns it.
	popNextTask := func(tasks *[]task) task {
		a := *tasks
		n := len(a)
		next, nextIndex := a[0], 0
		for i := 1; i < n; i++ {
			if t := a[i]; t.done < next.done {
				next, nextIndex = t, i
			}
		}
		a[nextIndex], a[n-1] = a[n-1], a[nextIndex]
		*tasks = a[:n-1]
		return next
	}

	var (
		working []task
		curTime = 0
	)
	for len(avail) > 0 || len(working) > 0 {
		for len(avail) > 0 && len(working) < numWorkers {
			step := popMin(&avail)
			working = append(working, task{step, curTime + duration(step)})
		}
		next := popNextTask(&working)
		ins.forEachChildOf(next.step, func(child string) {
			ins.removeParentFromChild(next.step, child)
			if ins.numParents(child) == 0 {
				avail = append(avail, child)
			}
		})
		curTime = next.done
	}
	fmt.Println(curTime)
}

var dependencyRE = regexp.MustCompile(`^Step (.*) must be finished before step (.*) can begin.$`)

// Parses the input.
func readInstructions() *instructions {
	type dependency struct {
		parent, child string // child depends on parent
	}
	var deps []dependency
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		m := dependencyRE.FindStringSubmatch(scanner.Text())
		deps = append(deps, dependency{parent: m[1], child: m[2]})
	}
	p2cs := make(map[string]stringSet) // parent -> children
	c2ps := make(map[string]stringSet) // child -> parents
	for _, dep := range deps {
		if p2cs[dep.parent] == nil {
			p2cs[dep.parent] = make(stringSet)
		}
		if c2ps[dep.child] == nil {
			c2ps[dep.child] = make(stringSet)
		}
		p2cs[dep.parent][dep.child] = true
		c2ps[dep.child][dep.parent] = true
	}
	return &instructions{p2cs: p2cs, c2ps: c2ps}
}

type (
	instructions struct {
		p2cs map[string]stringSet // maps a step (parent) to those that depend on it (children)
		c2ps map[string]stringSet // maps a step (child) to those it depends upon (parents)
	}
	stringSet map[string]bool
)

// Calls f for each step that has at least one dependent step.
func (ins *instructions) forEachParent(f func(string)) {
	for parent := range ins.p2cs {
		f(parent)
	}
}

// Calls f for each step that depends on the parent step.
func (ins *instructions) forEachChildOf(parent string, f func(string)) {
	for child := range ins.p2cs[parent] {
		f(child)
	}
}

// Removes a dependency relationship: child no longer depends on parent.
func (ins *instructions) removeParentFromChild(parent, child string) {
	delete(ins.c2ps[child], parent)
}

// Reports the number of steps that must be completed before the specified one.
func (ins *instructions) numParents(step string) int {
	return len(ins.c2ps[step])
}

// Removes the lexicographically first string from the slice, then returns it.
func popMin(p *[]string) string {
	a := *p
	var min string
	index := 0
	for i, s := range a {
		if min == "" || s < min {
			min, index = s, i
		}
	}
	n := len(a)
	a[index], a[n-1] = a[n-1], a[index]
	*p = a[:n-1]
	return min
}
