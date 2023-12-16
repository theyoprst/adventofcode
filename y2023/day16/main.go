package main

import (
	"github.com/theyoprst/adventofcode/aoc"
)

type Point struct {
	row, col int
}

type Dir Point

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

func CountEnergized(field aoc.ByteField, start Point, startDir Dir) int {
	type State struct {
		p   Point
		dir Dir
	}
	seen := map[State]bool{}
	seenPoints := map[Point]bool{}
	var dfs func(p Point, dir Dir)
	dfs = func(p Point, dir Dir) {
		if seen[State{p, dir}] {
			return
		}
		seen[State{p, dir}] = true
		ch := field[p.row][p.col]
		if ch == '*' {
			return
		}
		seenPoints[p] = true

		for _, ndir := range Map[ch][dir] {
			np := Point{
				row: p.row + ndir.row,
				col: p.col + ndir.col,
			}
			dfs(np, ndir)
		}
	}
	dfs(start, startDir)
	return len(seenPoints)
}

func SolvePart1(lines []string) any {
	field := aoc.MakeByteField(lines).AddBorder('*')
	return CountEnergized(field, Point{1, 1}, Right)
}

func SolvePart2(lines []string) any {
	field := aoc.MakeByteField(lines).AddBorder('*')
	rows := len(field) - 2
	cols := len(field[0]) - 2
	maxEn := 0
	for col := 1; col < 1+cols; col++ {
		maxEn = max(maxEn, CountEnergized(field, Point{row: 1, col: col}, Down))
		maxEn = max(maxEn, CountEnergized(field, Point{row: rows, col: col}, Up))
	}
	for row := 1; row < 1+rows; row++ {
		maxEn = max(maxEn, CountEnergized(field, Point{row: row, col: 1}, Right))
		maxEn = max(maxEn, CountEnergized(field, Point{row: row, col: cols}, Right))
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
