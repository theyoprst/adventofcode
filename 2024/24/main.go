package main

import (
	"context"
	"fmt"
	"iter"
	"os"
	"slices"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(_ context.Context, lines []string) any {
	wires := make(map[string]*Wire, len(lines))
	blocks := aoc.Blocks(lines)
	constants, binaries := blocks[0], blocks[1]
	for _, constExpr := range constants {
		wire := parseConstWire(constExpr)
		wires[wire.dst] = wire
	}
	for _, binaryExpr := range binaries {
		wire := parseBinaryWire(binaryExpr)
		wires[wire.dst] = wire
	}

	var eval func(string) int
	eval = func(name string) int {
		wire, ok := wires[name]
		must.True(ok)
		if wire.op == Const {
			return wire.value
		}
		src1 := eval(wire.src1)
		src2 := eval(wire.src2)
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

func SolvePart2(ctx context.Context, lines []string) any {
	if !aoc.GetParams(ctx).Bool("part2") {
		return nil
	}
	wires := make(map[string]*Wire, len(lines))
	blocks := aoc.Blocks(lines)
	binaries := blocks[1]
	for _, binaryExpr := range binaries {
		wire := parseBinaryWire(binaryExpr)
		wires[wire.dst] = wire
	}

	var eval func(zIdx int, name string, carries containers.Set[string], depth int) (expr *Expression, carryName string, used containers.Set[string])
	eval = func(zIdx int, name string, carries containers.Set[string], depth int) (expr *Expression, carryName string, used containers.Set[string]) {
		if depth > 5 { // TODO: better cycle detection?
			return nil, "", nil
		}
		used = containers.NewSet[string]()
		if strings.HasPrefix(name, "x") || strings.HasPrefix(name, "y") {
			var carryName string
			if must.Atoi(name[1:]) < zIdx {
				carryName = name
			}
			return newLiteralExpression(name), carryName, used
		}
		wire, ok := wires[name]
		must.True(ok)
		if carries.Has(name) {
			return newLiteralExpression(wire.dst), wire.dst, used
		}
		used.Add(name)
		if strings.HasPrefix(wire.dst, "z") {
			if must.Atoi(wire.dst[1:]) < zIdx {
				return newLiteralExpression(wire.dst), wire.dst, used
			}
		}
		left, leftCarryName, leftUsed := eval(zIdx, wire.src1, carries, depth+1)
		right, rightCarryName, rightUsed := eval(zIdx, wire.src2, carries, depth+1)
		if left == nil || right == nil {
			return nil, "", nil
		}
		expr = newExpression(name, wire.op, left, right)
		if leftCarryName != "" && rightCarryName != "" {
			carryName = name
		} else if leftCarryName != "" {
			carryName = leftCarryName
		} else if rightCarryName != "" {
			carryName = rightCarryName
		}
		return expr, carryName, used.Union(leftUsed).Union(rightUsed)
	}

	globalUsed := containers.NewSet[string]()

	type WirePair struct {
		first, second string
	}
	iterWirePairsToSwap := func(firstNames containers.Set[string]) iter.Seq[WirePair] {
		return func(yield func(WirePair) bool) {
			for firstName := range firstNames {
				for secondName, second := range wires {
					if firstName == secondName {
						continue
					}
					if globalUsed.Has(secondName) {
						continue
					}
					if !yield(WirePair{first: firstName, second: second.dst}) {
						return
					}
				}
			}
		}
	}

	carries := containers.NewSet[string]()
	var swapped []string
	var prevCarry string
	for zIdx := 0; ; zIdx++ {
		zName := fmt.Sprintf("z%02d", zIdx)
		if _, ok := wires[zName]; !ok {
			break
		}
		expr, carryName, localUsed := eval(zIdx, zName, carries, 0)
		if !isCorrectSumExpression(expr, zIdx, prevCarry) {
			var swap *WirePair
			count := 0
			for pair := range iterWirePairsToSwap(localUsed) {
				count++
				wires[pair.first].dst, wires[pair.second].dst = wires[pair.second].dst, wires[pair.first].dst
				wires[pair.first], wires[pair.second] = wires[pair.second], wires[pair.first]
				expr, carryName, localUsed = eval(zIdx, zName, carries, 0)
				if isCorrectSumExpression(expr, zIdx, prevCarry) {
					swapped = append(swapped, pair.first, pair.second)
					swap = &pair
					break
				}
				wires[pair.first].dst, wires[pair.second].dst = wires[pair.second].dst, wires[pair.first].dst
				wires[pair.first], wires[pair.second] = wires[pair.second], wires[pair.first]
			}
			if swap == nil {
				fmt.Println("No swap found. Count: ", count)
				os.Exit(1)
			}
		}

		carries.Add(carryName)
		globalUsed.Update(localUsed)
		prevCarry = carryName
	}
	slices.Sort(swapped)
	return strings.Join(swapped, ",")
}

type Operation string

const (
	Const Operation = "-"
	And   Operation = "&"
	Or    Operation = "|"
	Xor   Operation = "^"
)

type Wire struct {
	dst   string
	src1  string
	src2  string
	op    Operation
	value int // In case of Const operation
}

func parseWires(lines []string) map[string]*Wire {
	wires := make(map[string]*Wire, len(lines))
	blocks := aoc.Blocks(lines)
	constants, binaries := blocks[0], blocks[1]
	for _, constExpr := range constants {
		wire := parseConstWire(constExpr)
		wires[wire.dst] = wire
	}
	for _, binaryExpr := range binaries {
		wire := parseBinaryWire(binaryExpr)
		wires[wire.dst] = wire
	}
	return wires
}

func parseConstWire(expr string) *Wire {
	name, value := must.Split2(expr, ": ")
	return &Wire{
		dst:   name,
		op:    Const,
		value: must.Atoi(value),
	}
}

func parseBinaryWire(expr string) *Wire {
	expr, dst := must.Split2(expr, " -> ")
	src1, operation, src2 := must.Split3(expr, " ")
	return &Wire{
		dst:  dst,
		src1: src1,
		src2: src2,
		op:   parseOperation(operation),
	}
}

func parseOperation(op string) Operation {
	switch op {
	case "AND":
		return And
	case "OR":
		return Or
	case "XOR":
		return Xor
	default:
		panic(fmt.Sprintf("unknown operation %q", op))
	}
}

type Expression struct {
	Name     string
	Operator Operation
	Literal  string
	Operands []*Expression
}

func newExpression(name string, op Operation, operands ...*Expression) *Expression {
	return &Expression{
		Name:     name,
		Operator: op,
		Operands: operands,
	}
}

func newLiteralExpression(value string) *Expression {
	return &Expression{
		Literal: value,
	}
}

func (e *Expression) Render() string {
	return e.renderRec(0)
}

func (e *Expression) renderRec(depth int) string {
	if e == nil {
		return ""
	}
	if e.Literal != "" {
		return e.Literal
	}
	var result strings.Builder
	if depth > 0 {
		result.WriteString("(")
	}
	for i, operand := range e.Operands {
		if i > 0 {
			result.WriteString(" ")
			result.WriteString(string(e.Operator))
			result.WriteString(" ")
		}
		result.WriteString(operand.renderRec(depth + 1))
	}
	if depth > 0 {
		result.WriteString(")")
	}
	return result.String()
}

func (e *Expression) Normalize() {
	if e == nil {
		return
	}
	if e.Literal != "" {
		must.Equal(e.Operator, "")
		return
	}
	must.NotEqual(e.Operator, "")
	var newOperands []*Expression
	operandsChanged := false
	for _, operand := range e.Operands {
		operand.Normalize()
		if operand.Operator == e.Operator {
			newOperands = append(newOperands, operand.Operands...)
			operandsChanged = true
		} else {
			newOperands = append(newOperands, operand)
		}
	}
	if operandsChanged {
		e.Name = ""
	}
	slices.SortFunc(newOperands, func(a, b *Expression) int {
		return strings.Compare(a.Render(), b.Render())
	})
	e.Operands = newOperands
}

func isCorrectSumExpression(expr *Expression, zIdx int, prevCarry string) bool {
	expr.Normalize()
	if zIdx == 0 {
		return expr.Render() == "x00 ^ y00"
	}
	if zIdx == 1 {
		return expr.Render() == "(x00 & y00) ^ x01 ^ y01"
	}
	if zIdx == 45 { // TODO: 45 is a magic number
		want := fmt.Sprintf(
			"(%s & (x%02d ^ y%02d)) | (x%02d & y%02d)",
			prevCarry, zIdx-1, zIdx-1, zIdx-1, zIdx-1,
		)
		return expr.Render() == want
	}
	want := fmt.Sprintf(
		"((%s & (x%02d ^ y%02d)) | (x%02d & y%02d)) ^ x%02d ^ y%02d",
		prevCarry, zIdx-1, zIdx-1, zIdx-1, zIdx-1, zIdx, zIdx,
	)
	return expr.Render() == want
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
