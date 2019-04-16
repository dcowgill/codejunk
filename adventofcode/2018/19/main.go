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
	ipReg, instructions := readInput(os.Stdin)
	fmt.Printf("ipReg = %d\n", ipReg)
	fmt.Printf("len(instructions) = %d\n", len(instructions))
	r := executeProgram(Registers{}, ipReg, instructions)
	fmt.Printf("registers = %+v\n", r)
}

func part2() {
	// (used Wolfram Alpha)
	fmt.Printf("answer = %d\n", 1+2+7+14+167+334+1169+2338+4513+9026+
		31591+63182+753671+1507342+5275697+10551394)

	ipReg, instructions := readInput(os.Stdin)
	fmt.Printf("ipReg = %d\n", ipReg)
	fmt.Printf("len(instructions) = %d\n", len(instructions))
	var r Registers
	r[0] = 1
	executeProgram(r, ipReg, instructions)
	// one eternity later...
}

func readInput(r io.Reader) (ipReg int, instructions []Instruction) {
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

func parseInstruction(s string) Instruction {
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
	inst.opname = opname
	inst.op = op
	inst.a, _ = strconv.Atoi(matches[2])
	inst.b, _ = strconv.Atoi(matches[3])
	inst.c, _ = strconv.Atoi(matches[4])
	return inst
}

const numRegisters = 6

type Registers [numRegisters]int

type Op func(xs *Registers, a, b, c int)

type Instruction struct {
	opname  string
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
// until the IP goes out of bounds. Returns the final state of the registers.
func executeProgram(r Registers, ipReg int, instructions []Instruction) Registers {
	prev := -1
outer:
	for {
		if prev != r[0] {
			fmt.Printf("r[0] = %d\n", r[0])
			prev = r[0]
		}
		ip := r[ipReg]
		if ip < 0 || ip >= len(instructions) {
			break
		}
		if ip == 3 {
			// Unroll this inner loop:
			//
			//	3:  mulr 3 1 5
			//	4:  eqrr 5 4 5
			//	5:  addr 5 2 2
			//	6:  addi 2 1 2
			//	7:  addr 3 0 0
			//	8:  addi 1 1 1
			//	9:  gtrr 1 4 5
			//	10: addr 2 5 2
			//	11: seti 2 9 2
			//
			for {
				if r[3]*r[1] == r[4] {
					r[0] += r[3]
				}
				r[1]++
				if r[1] > r[4] {
					r[2] = 12
					continue outer
				}
			}
		}
		inst := instructions[ip]
		inst.op(&r, inst.a, inst.b, inst.c)
		r[ipReg]++
	}
	return r
}
