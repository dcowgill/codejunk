package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
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
	ipReg, instructions := readInput(os.Stdin)
	executeProgram(Registers{}, ipReg, instructions)
}

func part2() {
	part1() // very slow
}

func readInput(r io.Reader) (ipReg int, instructions []*Instruction) {
	scanner := bufio.NewScanner(r)
	ipReg = parseIPRegHeader(scanner)
	for scanner.Scan() {
		instructions = append(instructions, parseInstruction(scanner.Text()))
	}
	return ipReg, instructions
}

// Reads line from the scanner and parses it as "#ip <int>". Returns the int.
func parseIPRegHeader(sc *bufio.Scanner) int {
	const prefix = "#ip "
	fail := func(err error) {
		msg := fmt.Sprintf("invalid header %q: expected '#ip <int>'")
		if err != nil {
			msg += " " + err.Error()
		}
		panic(msg)
	}
	if !sc.Scan() {
		fail(nil)
	}
	line := sc.Text()
	if !strings.HasPrefix(line, prefix) {
		fail(nil)
	}
	ipReg, err := strconv.Atoi(line[len(prefix):])
	if err != nil {
		fail(err)
	}
	return ipReg
}

var instRE = regexp.MustCompile(`^([a-z]+) (\d+) (\d+) (\d+)$`)

func parseInstruction(s string) *Instruction {
	matches := instRE.FindStringSubmatch(s)
	if matches == nil {
		panic(fmt.Sprintf("failed to match %q with regexp %s", s, instRE))
	}
	opname := matches[1]
	op, ok := ops[opname]
	if !ok {
		panic(fmt.Sprintf("invalid opname: %q", opname))
	}
	var inst Instruction
	inst.op = op
	inst.a, _ = strconv.Atoi(matches[2])
	inst.b, _ = strconv.Atoi(matches[3])
	inst.c, _ = strconv.Atoi(matches[4])
	return &inst
}

const numRegisters = 6

type Registers [numRegisters]int

type Op func(xs *Registers, a, b, c int)

type Instruction struct {
	op      Op
	a, b, c int
}

var ops = map[string]Op{
	"addr": func(r *Registers, a, b, c int) { r[c] = r[a] + r[b] },
	"addi": func(r *Registers, a, b, c int) { r[c] = r[a] + b },
	"mulr": func(r *Registers, a, b, c int) { r[c] = r[a] * r[b] },
	"muli": func(r *Registers, a, b, c int) { r[c] = r[a] * b },
	"banr": func(r *Registers, a, b, c int) { r[c] = r[a] & r[b] },
	"bani": func(r *Registers, a, b, c int) { r[c] = r[a] & b },
	"borr": func(r *Registers, a, b, c int) { r[c] = r[a] | r[b] },
	"bori": func(r *Registers, a, b, c int) { r[c] = r[a] | b },
	"setr": func(r *Registers, a, b, c int) { r[c] = r[a] },
	"seti": func(r *Registers, a, b, c int) { r[c] = a },
	"gtir": func(r *Registers, a, b, c int) { r[c] = intBool(a > r[b]) },
	"gtri": func(r *Registers, a, b, c int) { r[c] = intBool(r[a] > b) },
	"gtrr": func(r *Registers, a, b, c int) { r[c] = intBool(r[a] > r[b]) },
	"eqir": func(r *Registers, a, b, c int) { r[c] = intBool(a == r[b]) },
	"eqri": func(r *Registers, a, b, c int) { r[c] = intBool(r[a] == b) },
	"eqrr": func(r *Registers, a, b, c int) { r[c] = intBool(r[a] == r[b]) },
}

// Returns 1 if b is true, else 0.
func intBool(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Given a starting set of register values, the index of the register containing
// the instruction pointer, and a set of instructions, executes the instructions
// until the IP goes out of bounds. Returns the final state of the registers and
// the number of instructions executed.
func executeProgram(r Registers, ipReg int, instructions []*Instruction) Registers {
	var (
		maxr = math.MinInt64
		seen = make(map[int]bool)
		prev = 0
	)
	for {
		ip := r[ipReg]
		if ip < 0 || ip >= len(instructions) {
			break
		}
		if ip == 28 { // eqrr 5 0 2
			if r[5] > maxr {
				fmt.Printf("new halt max r[0] = %d\n", r[5])
				maxr = r[5]
			}
			if seen[r[5]] {
				fmt.Printf("duplicate in r[0] = %d (previous = %d)\n", r[5], prev)
				break
			}
			seen[r[5]] = true
			prev = r[5]
		}
		inst := instructions[ip]
		inst.op(&r, inst.a, inst.b, inst.c)
		r[ipReg]++
	}
	return r
}
