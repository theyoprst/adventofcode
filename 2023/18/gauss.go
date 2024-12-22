package main

import (
	"context"

	"github.com/theyoprst/adventofcode/aoc/fld"
	"github.com/theyoprst/adventofcode/must"
)

func SolveGenericGauss(commands []Command) int {
	pos := fld.Zero
	pp := []fld.Pos{pos}
	pathLen := 0
	for _, cmd := range commands {
		pos = pos.Add(cmd.dir.Mult(cmd.steps))
		pp = append(pp, pos)
		pathLen += cmd.steps
	}
	must.Equal(pp[len(pp)-1], fld.Zero)
	area := 0
	for i := 0; i < len(pp)-1; i++ {
		cur := pp[i]
		next := pp[i+1]
		area += (cur.Row + next.Row) * (cur.Col - next.Col)
	}
	must.Equal(area%2, 0)
	area /= 2
	area += pathLen/2 + 1
	return area
}

func SolvePart1Gauss(_ context.Context, lines []string) any {
	return SolveGenericGauss(ParseCommands(lines))
}

func SolvePart2Gauss(_ context.Context, lines []string) any {
	return SolveGenericGauss(ParseCommands2(lines))
}
