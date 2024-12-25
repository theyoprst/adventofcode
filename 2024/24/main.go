package main

import (
	"context"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/must"
)

type Operation string

const (
	Const Operation = "-"
	And   Operation = "AND"
	Or    Operation = "OR"
	Xor   Operation = "XOR"
)

type Wire struct {
	name   string
	first  string
	second string
	op     Operation
	value  int // In case of Const operation
}

func parseConstWire(expr string) *Wire {
	name, value := must.Split2(expr, ": ")
	return &Wire{
		name:  name,
		op:    Const,
		value: must.Atoi(value),
	}
}

func parseBinaryWire(expr string) *Wire {
	expr, dst := must.Split2(expr, " -> ")
	src1, operation, src2 := must.Split3(expr, " ")
	return &Wire{
		name:   dst,
		first:  src1,
		second: src2,
		op:     Operation(operation),
	}
}

func SolvePart1(_ context.Context, lines []string) any {
	wires := make(map[string]*Wire, len(lines))
	blocks := aoc.Blocks(lines)
	constants, binaries := blocks[0], blocks[1]
	for _, constExpr := range constants {
		wire := parseConstWire(constExpr)
		wires[wire.name] = wire
	}
	for _, binaryExpr := range binaries {
		wire := parseBinaryWire(binaryExpr)
		wires[wire.name] = wire
	}

	var eval func(string) int
	eval = func(name string) int {
		wire, ok := wires[name]
		must.True(ok)
		if wire.op == Const {
			return wire.value
		}
		src1 := eval(wire.first)
		src2 := eval(wire.second)
		switch wire.op {
		case And:
			return src1 & src2
		case Or:
			return src1 | src2
		case Xor:
			return src1 ^ src2
		default:
			panic(fmt.Sprintf("unknown operation %v", wire.op))
		}
	}

	ans := 0
	for zIdx := 0; ; zIdx++ {
		zName := fmt.Sprintf("z%02d", zIdx)
		if _, ok := wires[zName]; !ok {
			break
		}
		value := eval(zName)
		ans += value << zIdx
	}
	return ans
}

type Expr struct {
	Op     Operation
	First  *Expr
	Second *Expr
	Name   string
}

func newLiteral(name string) *Expr {
	return &Expr{
		Op:   Const,
		Name: name,
	}
}

func newBinary(first *Expr, op Operation, second *Expr) *Expr {
	return &Expr{
		Op:     op,
		First:  first,
		Second: second,
	}
}

func SolvePart2(ctx context.Context, lines []string) any {
	// We've got eletrical scheme for Ripple-carry adder which is broken in 8 places: 4 pairs of wires are swapped.
	// The scheme adds 45-bit numbers x and y, and produces 46-bit result z.
	// There are 3 types of gates: AND, OR, XOR.
	// Looks like the minimal expression for each bit is (a little different for the first two and the last indexes):
	//   z = (x ^ y) ^ c, where
	//   c = (xPrev & yPrev) | (cPrev & (xPrev ^ yPrev))
	// `c` is a carry bit.
	// There are 5 gates (binary operations) used if we re-use (xPrev ^ yPref) from the previous bit.
	// In total, there are 5 gates for 43 bits, 1 gate for the 1st bit, 3 gates for the 2nd bit, 3 gates for the last bit.
	// 222 gates in total for the minimal circuit, exactly 222 gates in the input. So we have a minimal circuit.
	//
	// The idea is for each bit build the expected expression that depend on 5 variables, and try to match it with the input.
	// If not matched, try to swap some wires until it matched. It appears that 1 swap is enough for each bit.
	// z = x ^ (y ^ c) and similar are impossible because we want to reuse x^y later for the carry bit.
	// Matching tries to swap order for each binary operation, 2^5 swaps in total in the worst case, not a big deal.
	if !aoc.GetParams(ctx).Bool("part2") {
		return nil
	}
	wires := make(map[string]*Wire, len(lines))
	for _, binaryExpr := range aoc.Blocks(lines)[1] {
		wire := parseBinaryWire(binaryExpr)
		wires[wire.name] = wire
	}

	notUsedWires := containers.NewSet(slices.Collect(maps.Keys(wires))...)

	var swapped []string
	var prevCarry string
	for zIdx := range zMax + 1 {
		wantExpr, wantCarryExpr := buildExpectedExpr(zIdx, prevCarry)

		zName := fmt.Sprintf("z%02d", zIdx)
		matched, localUsed := matchExpr(zName, wantExpr, wires)
		if !matched {
			swapFound := false
		iterSwapsLoop:
			for first := range localUsed {
				for second := range notUsedWires {
					if first != second {
						wires[first].name, wires[second].name = wires[second].name, wires[first].name
						wires[first], wires[second] = wires[second], wires[first]
						matched, localUsed = matchExpr(zName, wantExpr, wires)
						if matched {
							swapped = append(swapped, first, second)
							swapFound = true
							break iterSwapsLoop
						}
						wires[first].name, wires[second].name = wires[second].name, wires[first].name
						wires[first], wires[second] = wires[second], wires[first]
					}
				}
			}
			if !swapFound {
				panic("no swap found")
			}
		}

		notUsedWires.Remove(slices.Collect(maps.Keys(localUsed))...)

		if wantCarryExpr != nil {
			prevCarry = wantCarryExpr.Name // Name was populated in matchExpr.
		}
	}
	slices.Sort(swapped)
	return strings.Join(swapped, ",")
}

const zMax = 45

// buildExpectedExpr builds the expected expression for the zIdx-th bit of z.
func buildExpectedExpr(zIdx int, prevCarry string) (expr *Expr, carry *Expr) {
	// Expect expr: (xCur ^ yCur)  ^ carry
	//                    |            |
	//             (not last bit)  (not first bit)
	var xorXY *Expr
	if zIdx != zMax {
		xorXY = newBinary(
			newLiteral("x"+fmt.Sprintf("%02d", zIdx)),
			Xor,
			newLiteral("y"+fmt.Sprintf("%02d", zIdx)),
		)
		if zIdx == 0 {
			return xorXY, nil
		}
	}
	// if zIdx = 1: carry = xPrev & yPrev
	// if zIdx > 1: carry = (xPrev & yPrev) | (carry & (xPrev ^ yPrev))
	xPrev := newLiteral("x" + fmt.Sprintf("%02d", zIdx-1))
	yPrev := newLiteral("y" + fmt.Sprintf("%02d", zIdx-1))
	carry = newBinary(xPrev, And, yPrev)
	if zIdx > 1 {
		carry = newBinary(
			carry,
			Or,
			newBinary(newLiteral(prevCarry), And, newBinary(xPrev, Xor, yPrev)),
		)
	}
	if xorXY == nil {
		must.Equal(zIdx, zMax)
		return carry, carry
	}
	return newBinary(xorXY, Xor, carry), carry // Most popular case for zIdx in 1..44
}

// matchExpr tries to match the wire with the expected expression.
// It returns bool for matching and used (visited) binary wires.
// Also populates names in wantExpr (used later to find out name of the carry expression).
func matchExpr(name string, wantExpr *Expr, wires map[string]*Wire) (matched bool, used containers.Set[string]) {
	used = containers.NewSet[string]()

	var matchExprRecursive func(name string, wantExpr *Expr) bool
	matchExprRecursive = func(name string, wantExpr *Expr) bool {
		if wantExpr.Op == Const {
			return name == wantExpr.Name
		}
		wire, isBinaryWire := wires[name]
		if !isBinaryWire {
			return false // wantExpr is binary operation by this moment.
		}
		used.Add(wire.name) // Mark binary wire as visited.
		if wire.op != wantExpr.Op {
			return false
		}
		wantExpr.Name = wire.name
		return matchExprRecursive(wire.first, wantExpr.First) && matchExprRecursive(wire.second, wantExpr.Second) ||
			matchExprRecursive(wire.first, wantExpr.Second) && matchExprRecursive(wire.second, wantExpr.First)
	}

	return matchExprRecursive(name, wantExpr), used
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
