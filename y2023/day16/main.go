package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

var Map = map[byte]map[fld.Pos][]fld.Pos{
	'-': {
		fld.Right: {fld.Right},
		fld.Left:  {fld.Left},
		fld.Up:    {fld.Right, fld.Left},
		fld.Down:  {fld.Right, fld.Left},
	},
	'|': {
		fld.Right: {fld.Up, fld.Down},
		fld.Left:  {fld.Up, fld.Down},
		fld.Up:    {fld.Up},
		fld.Down:  {fld.Down},
	},
	'\\': {
		fld.Right: {fld.Down},
		fld.Left:  {fld.Up},
		fld.Up:    {fld.Left},
		fld.Down:  {fld.Right},
	},
	'/': {
		fld.Right: {fld.Up},
		fld.Left:  {fld.Down},
		fld.Up:    {fld.Right},
		fld.Down:  {fld.Left},
	},
	'.': {
		fld.Right: {fld.Right},
		fld.Left:  {fld.Left},
		fld.Up:    {fld.Up},
		fld.Down:  {fld.Down},
	},
}

func CountEnergized(field fld.ByteField, start, startDir fld.Pos) int {
	type State struct {
		p, dir fld.Pos
	}
	seen := containers.NewSet[State]()
	seenPoints := containers.NewSet[fld.Pos]()
	var dfs func(p, dir fld.Pos)
	dfs = func(p, dir fld.Pos) {
		if !field.Inside(p) {
			return
		}
		if seen.Has(State{p, dir}) {
			return
		}
		seen.Add(State{p, dir})
		ch := field.Get(p)
		seenPoints.Add(p)

		for _, ndir := range Map[ch][dir] {
			np := p.Add(ndir)
			dfs(np, ndir)
		}
	}
	dfs(start, startDir)
	return len(seenPoints)
}

func SolvePart1(lines []string) any {
	field := fld.NewByteField(lines)
	return CountEnergized(field, fld.NewPos(0, 0), fld.Right)
}

func SolvePart2(lines []string) any {
	field := fld.NewByteField(lines)
	ans := 0
	for col := 0; col < field.Cols(); col++ {
		ans = max(ans, CountEnergized(field, fld.NewPos(0, col), fld.Down))
		ans = max(ans, CountEnergized(field, fld.NewPos(field.Rows()-1, col), fld.Up))
	}
	for row := 0; row < field.Rows(); row++ {
		ans = max(ans, CountEnergized(field, fld.NewPos(row, 0), fld.Right))
		ans = max(ans, CountEnergized(field, fld.NewPos(row, field.Cols()-1), fld.Right))
	}
	return ans
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
