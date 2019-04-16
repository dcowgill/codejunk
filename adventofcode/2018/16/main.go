package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
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
	examples, _ := readInput(os.Stdin)
	answer := 0 // number of examples matching 3 or more opcodes
nextExample:
	for _, ex := range examples {
		numMatches := 0
		for _, op := range ops {
			if execOp(op, ex.before, ex.inst) == ex.after {
				numMatches++
				if numMatches >= 3 {
					answer++
					continue nextExample
				}
			}
		}
	}
	fmt.Printf("%d examples match 3 or more ops\n", answer)
}

func part2() {
	// Make a slice of the ops whose opcodes we do not know. Define a helper
	// function to remove an op from the list once we know its code.
	unknown := make([]*Op, len(ops))
	copy(unknown, ops)
	eliminateUnknown := func(i int) {
		n := len(unknown)
		unknown[i], unknown[n-1] = unknown[n-1], unknown[i]
		unknown = unknown[:n-1]
	}

	examples, instructions := readInput(os.Stdin)
	opcodeToOp := make(map[int]*Op, len(ops))

	// Assumption: there will always be at least one opcode whose associated
	// examples can be solved by exactly one op that remains in "unknown";
	// in other words, we assume that backtracking will *not* be required.
outer:
	for len(unknown) != 0 {
	nextExample:
		for _, ex := range examples {
			match := -1
			for i, op := range unknown {
				if execOp(op, ex.before, ex.inst) == ex.after {
					if match >= 0 {
						continue nextExample
					}
					match = i
				}
			}
			opcodeToOp[ex.inst.opcode()] = unknown[match]
			eliminateUnknown(match)
			examples = filterExamplesInPlace(examples, func(e *Example) bool {
				return e.inst.opcode() != ex.inst.opcode()
			})
			continue outer
		}
	}

	// Show the mapping from opcode to op. (Not part of the answer.)
	for code, op := range opcodeToOp {
		fmt.Printf("opcode %2d = %s\n", code, op.name)
	}

	// Run the program.
	var regs Registers
	for _, inst := range instructions {
		op := opcodeToOp[inst.opcode()]
		regs = execOp(op, regs, inst)
	}

	// Show the final state of the registers.
	fmt.Printf("final registers: %+v\n", regs)
}

// Reports whether the given op fits with the example.
func opWorksForExample(op *Op, ex *Example) bool {
	return execOp(op, ex.before, ex.inst) == ex.after
}

// Executes an op; returns the new registers.
func execOp(op *Op, r Registers, inst Instruction) Registers {
	op.exec(&r, inst.a(), inst.b(), inst.c())
	return r
}

const numRegisters = 4

func readInput(r io.Reader) ([]*Example, []Instruction) {
	scanner := bufio.NewScanner(r)
	mustScan := func() string {
		if !scanner.Scan() {
			panic("unexpected EOF")
		}
		return scanner.Text()
	}
	var (
		examples     []*Example
		instructions []Instruction
	)
	for scanner.Scan() {
		var ex Example
		s := scanner.Text()
		if strings.HasPrefix(s, "Before: [") {
			ex.before = parseRegisters(s)
			s = mustScan()
			ex.inst = parseInstruction(s)
			s = mustScan()
			ex.after = parseRegisters(s)
			examples = append(examples, &ex)
		} else if s != "" {
			instructions = append(instructions, parseInstruction(s))
		}
	}
	return examples, instructions
}

func parseRegisters(s string) Registers {
	var r Registers
	for i, x := range parseInts(s, numRegisters) {
		r[i] = x
	}
	return r
}

func parseInstruction(s string) Instruction {
	var inst Instruction
	for i, x := range parseInts(s, 4) {
		inst[i] = x
	}
	return inst
}

var intRE = regexp.MustCompile(`[0-9]+`)

func parseInts(s string, n int) []int {
	matches := intRE.FindAllString(s, n)
	xs := make([]int, len(matches))
	for i, m := range matches {
		xs[i], _ = strconv.Atoi(m)
	}
	return xs
}

type (
	Registers   [4]int
	Instruction [4]int
	Example     struct {
		before Registers
		inst   Instruction
		after  Registers
	}
)

func (inst Instruction) opcode() int { return inst[0] }
func (inst Instruction) a() int      { return inst[1] }
func (inst Instruction) b() int      { return inst[2] }
func (inst Instruction) c() int      { return inst[3] }

// Removes from examples any items that do not match the predicate.
// Returns the modified slice.
func filterExamplesInPlace(examples []*Example, pred func(*Example) bool) []*Example {
	i := 0
	for _, ex := range examples {
		if pred(ex) {
			examples[i] = ex
			i++
		}
	}
	return examples[:i]
}

type Op struct {
	name string
	exec func(xs *Registers, a, b, c int)
}

var ops = []*Op{
	{name: "addr", exec: func(r *Registers, a, b, c int) { r[c] = r[a] + r[b] }},
	{name: "addi", exec: func(r *Registers, a, b, c int) { r[c] = r[a] + b }},
	{name: "mulr", exec: func(r *Registers, a, b, c int) { r[c] = r[a] * r[b] }},
	{name: "muli", exec: func(r *Registers, a, b, c int) { r[c] = r[a] * b }},
	{name: "banr", exec: func(r *Registers, a, b, c int) { r[c] = r[a] & r[b] }},
	{name: "bani", exec: func(r *Registers, a, b, c int) { r[c] = r[a] & b }},
	{name: "borr", exec: func(r *Registers, a, b, c int) { r[c] = r[a] | r[b] }},
	{name: "bori", exec: func(r *Registers, a, b, c int) { r[c] = r[a] | b }},
	{name: "setr", exec: func(r *Registers, a, b, c int) { r[c] = r[a] }},
	{name: "seti", exec: func(r *Registers, a, b, c int) { r[c] = a }},
	{name: "gtir", exec: func(r *Registers, a, b, c int) { r[c] = intBool(a > r[b]) }},
	{name: "gtri", exec: func(r *Registers, a, b, c int) { r[c] = intBool(r[a] > b) }},
	{name: "gtrr", exec: func(r *Registers, a, b, c int) { r[c] = intBool(r[a] > r[b]) }},
	{name: "eqir", exec: func(r *Registers, a, b, c int) { r[c] = intBool(a == r[b]) }},
	{name: "eqri", exec: func(r *Registers, a, b, c int) { r[c] = intBool(r[a] == b) }},
	{name: "eqrr", exec: func(r *Registers, a, b, c int) { r[c] = intBool(r[a] == r[b]) }},
}

// Returns 1 if b is true, else 0.
func intBool(b bool) int {
	if b {
		return 1
	}
	return 0
}
