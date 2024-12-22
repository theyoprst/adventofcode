package main

import (
	"context"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
	"github.com/theyoprst/adventofcode/must"
)

var sandSrc = fld.Pos{Row: 0, Col: 500}

func SolvePart1(_ context.Context, lines []string) any {
	field, abyss := parseField(lines)
	count := 0
	for {
		pos := fallSand(field, abyss)
		if pos.Row == abyss {
			break
		}
		count++
		field.Add(pos)
	}
	return count
}

func SolvePart2(_ context.Context, lines []string) any {
	field, abyss := parseField(lines)
	count := 0
	for !field.Has(sandSrc) {
		field.Add(fallSand(field, abyss))
		count++
	}
	return count
}

func parseField(lines []string) (field containers.Set[fld.Pos], abyss int) {
	field = containers.NewSet[fld.Pos]()
	maxY := 0
	for _, line := range lines {
		path := parsePath(line)
		for _, p := range path {
			maxY = max(maxY, p.Row)
		}
		for i := range len(path) - 1 {
			p1, p2 := path[i], path[i+1]
			diff := p2.Sub(p1)
			must.True(diff.Row == 0 || diff.Col == 0)
			diff.Row = min(max(diff.Row, -1), 1)
			diff.Col = min(max(diff.Col, -1), 1)
			for p := p1; p != p2; p = p.Add(diff) {
				field.Add(p)
			}
			field.Add(p2)
		}
	}
	return field, maxY + 1
}

func parsePath(line string) []fld.Pos {
	var path []fld.Pos
	nn := aoc.Ints(line)
	must.Equal(len(nn)%2, 0)
	for i := 0; i < len(nn); i += 2 {
		path = append(path, fld.Pos{Col: nn[i], Row: nn[i+1]})
	}
	return path
}

func fallSand(field containers.Set[fld.Pos], abyss int) fld.Pos {
	pos := sandSrc
fallLoop:
	for pos.Row < abyss {
		for _, dir := range []fld.Pos{fld.Down, fld.DownLeft, fld.DownRight} {
			npos := pos.Add(dir)
			if !field.Has(npos) {
				pos = npos
				continue fallLoop
			}
		}
		break
	}
	return pos
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
