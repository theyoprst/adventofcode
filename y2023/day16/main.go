package main

import (
	"github.com/theyoprst/adventofcode/aoc"
)

type Dir aoc.FieldPos

var (
	Right Dir = Dir{0, 1}
	Left  Dir = Dir{0, -1}
	Up    Dir = Dir{-1, 0}
	Down  Dir = Dir{1, 0}

	Map = map[byte]map[Dir][]Dir{
		'-': {
			Right: {Right},
			Left:  {Left},
			Up:    {Right, Left},
			Down:  {Right, Left},
		},
		'|': {
			Right: {Up, Down},
			Left:  {Up, Down},
			Up:    {Up},
			Down:  {Down},
		},
		'\\': {
			Right: {Down},
			Left:  {Up},
			Up:    {Left},
			Down:  {Right},
		},
		'/': {
			Right: {Up},
			Left:  {Down},
			Up:    {Right},
			Down:  {Left},
		},
		'.': {
			Right: {Right},
			Left:  {Left},
			Up:    {Up},
			Down:  {Down},
		},
	}
)

func CountEnergized(field aoc.ByteField, start aoc.FieldPos, startDir Dir) int {
	type State struct {
		p   aoc.FieldPos
		dir Dir
	}
	seen := map[State]bool{}
	seenPoints := map[aoc.FieldPos]bool{}
	var dfs func(p aoc.FieldPos, dir Dir)
	dfs = func(p aoc.FieldPos, dir Dir) {
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
			np := p.Add(aoc.FieldPos(ndir))
			dfs(np, ndir)
		}
	}
	dfs(start, startDir)
	return len(seenPoints)
}

func SolvePart1(lines []string) any {
	field := aoc.MakeByteField(lines)
	return CountEnergized(field, aoc.FieldPos{Row: 0, Col: 0}, Right)
}

func SolvePart2(lines []string) any {
	field := aoc.MakeByteField(lines)
	maxEn := 0
	for col := 0; col < field.Cols(); col++ {
		maxEn = max(maxEn, CountEnergized(field, aoc.FieldPos{Row: 0, Col: col}, Down))
		maxEn = max(maxEn, CountEnergized(field, aoc.FieldPos{Row: field.Rows() - 1, Col: col}, Up))
	}
	for row := 0; row < field.Rows(); row++ {
		maxEn = max(maxEn, CountEnergized(field, aoc.FieldPos{Row: row, Col: 0}, Right))
		maxEn = max(maxEn, CountEnergized(field, aoc.FieldPos{Row: row, Col: field.Cols() - 1}, Right))
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
