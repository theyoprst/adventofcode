package main

import (
	"context"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(_ context.Context, lines []string) any {
	return solveGeneric(1, lines)
}

func SolvePart2(_ context.Context, lines []string) any {
	return solveGeneric(9, lines)
}

func solveGeneric(tailSize int, lines []string) int {
	var head fld.Pos
	tailVisited := containers.NewSet(head)
	tails := make([]fld.Pos, tailSize)
	for i := range tails {
		tails[i] = head
	}
	for _, line := range lines {
		cmd, stepsStr := must.Split2(line, " ")
		steps := must.Atoi(stepsStr)
		var dir fld.Pos
		switch cmd {
		case "R":
			dir = fld.Right
		case "L":
			dir = fld.Left
		case "U":
			dir = fld.Up
		case "D":
			dir = fld.Down
		default:
			panic("Unknown command")
		}
		for i := 0; i < steps; i++ {
			head = head.Add(dir)
			prev := head
			for i, tail := range tails {
				if max(aoc.Abs(prev.Row-tail.Row), aoc.Abs(prev.Col-tail.Col)) > 1 {
					delta := prev.Sub(tail)
					delta.Row = min(max(-1, delta.Row), 1)
					delta.Col = min(max(-1, delta.Col), 1)
					tail = tail.Add(delta)
					tails[i] = tail
				}
				prev = tail
			}
			tailVisited.Add(tails[len(tails)-1])
		}
	}
	return len(tailVisited)
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
