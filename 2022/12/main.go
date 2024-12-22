package main

import (
	"context"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

func SolvePart1(_ context.Context, lines []string) any {
	return solve(lines, false)
}

func SolvePart2(_ context.Context, lines []string) any {
	return solve(lines, true)
}

func solve(lines []string, startFromZeros bool) any {
	field := fld.NewByteField(lines)

	height := func(pos fld.Pos) int {
		h := field.Get(pos)
		if h == 'S' {
			h = 'a'
		} else if h == 'E' {
			h = 'z'
		}
		return int(h - 'a')
	}

	start := field.FindFirst('S')
	end := field.FindFirst('E')

	// Let BFS begin.
	dist := map[fld.Pos]int{start: 0}
	queue := []fld.Pos{start}
	if startFromZeros {
		for pos := range field.IterPositions() {
			if height(pos) == 0 {
				dist[pos] = 0
				queue = append(queue, pos)
			}
		}
	}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		if pos == end {
			return dist[pos]
		}

		for _, dir := range fld.DirsSimple {
			next := pos.Add(dir)
			if !field.Inside(next) {
				continue
			}
			if height(next)-height(pos) < 2 {
				if _, ok := dist[next]; !ok {
					dist[next] = dist[pos] + 1
					queue = append(queue, next)
				}
			}
		}
	}
	panic("no path found")
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
