package main

import (
	"context"
	"iter"
	"slices"
	"strconv"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

type OpCode int

const (
	OpCodeADV OpCode = 0
	OpCodeBXL OpCode = 1
	OpCodeBST OpCode = 2
	OpCodeJNZ OpCode = 3
	OpCodeBXC OpCode = 4
	OpCodeOUT OpCode = 5
	OpCodeBDV OpCode = 6
	OpCodeCDV OpCode = 7
)

func SolvePart1(_ context.Context, lines []string) any {
	a := aoc.Ints(lines[0])[0]
	b := aoc.Ints(lines[1])[0]
	c := aoc.Ints(lines[2])[0]
	var output []string
	for x := range runProgram(aoc.Ints(lines[4]), a, b, c) {
		output = append(output, strconv.Itoa(x))
	}
	return strings.Join(output, ",")
}

func SolvePart2(_ context.Context, lines []string) any {
	program := aoc.Ints(lines[4])
	wantOut := slices.Clone(program)
	slices.Reverse(wantOut)

	aShift := 0
	for pc := 0; pc < len(program); pc += 2 {
		op := OpCode(program[pc])
		operand := program[pc+1]
		switch op {
		case OpCodeADV:
			// Only solve the problem when A is shifted by a constant (1, 2, 3), ...
			must.Greater(operand, 0)
			must.Less(operand, 4)
			aShift += operand
		case OpCodeJNZ:
			// ... and the jump is always from the end to the beginning.
			must.Equal(operand, 0)
			must.Equal(pc+2, len(program))
		}
	}

	var goodA []int

	var dfs func(a, iter int)
	dfs = func(a int, iter int) {
		if iter == len(wantOut) {
			goodA = append(goodA, a)
			return
		}
		for word := range 8 {
			na := a<<aShift + word
			if firstOutput(program, na, 0, 0) == wantOut[iter] {
				dfs(na, iter+1)
			}
		}
	}
	dfs(0, 0)

	return slices.Min(goodA)
}

func firstOutput(program []int, a, b, c int) int {
	for x := range runProgram(program, a, b, c) {
		return x
	}
	panic("no output")
}

// runProgram runs the program and iterates over the output.
func runProgram(program []int, a, b, c int) iter.Seq[int] {
	combo := func(operand int) int {
		switch operand {
		case 0, 1, 2, 3:
			return operand
		case 4:
			return a
		case 5:
			return b
		case 6:
			return c
		default:
			panic("invalid operand")
		}
	}
	return func(yield func(int) bool) {
		pc := 0
		for pc < len(program)-1 {
			operand := program[pc+1]
			switch OpCode(program[pc]) {
			case OpCodeADV:
				a >>= combo(operand)
			case OpCodeBXL:
				b ^= operand
			case OpCodeBST:
				b = combo(operand) & 7
			case OpCodeJNZ:
				if a != 0 {
					pc = operand
					continue
				}
			case OpCodeBXC:
				b ^= c
			case OpCodeOUT:
				if !yield(combo(operand) & 7) {
					return
				}
			case OpCodeBDV:
				b = a >> combo(operand)
			case OpCodeCDV:
				c = a >> combo(operand)
			}
			pc += 2
		}
	}
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
