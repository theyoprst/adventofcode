package main

import (
	"github.com/theyoprst/adventofcode/aoc"
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
	seen := map[State]bool{}
	seenPoints := map[fld.Pos]bool{}
	var dfs func(p, dir fld.Pos)
	dfs = func(p, dir fld.Pos) {
		if !field.Inside(p) {
			return
		}
		if seen[State{p, dir}] {
			return
		}
		seen[State{p, dir}] = true
		ch := field[p.Row][p.Col]
		seenPoints[p] = true

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
	return CountEnergized(field, fld.Pos{Row: 0, Col: 0}, fld.Right)
}

func SolvePart2(lines []string) any {
	field := fld.NewByteField(lines)
	maxEn := 0
	for col := 0; col < field.Cols(); col++ {
		maxEn = max(maxEn, CountEnergized(field, fld.Pos{Row: 0, Col: col}, fld.Down))
		maxEn = max(maxEn, CountEnergized(field, fld.Pos{Row: field.Rows() - 1, Col: col}, fld.Up))
	}
	for row := 0; row < field.Rows(); row++ {
		maxEn = max(maxEn, CountEnergized(field, fld.Pos{Row: row, Col: 0}, fld.Right))
		maxEn = max(maxEn, CountEnergized(field, fld.Pos{Row: row, Col: field.Cols() - 1}, fld.Right))
	}
	return maxEn
}

var (
	solvers1 []aoc.Solver = []aoc.Solver{SolvePart1}
	solvers2 []aoc.Solver = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
