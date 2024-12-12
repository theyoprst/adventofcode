package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

func SolvePart1(lines []string) any {
	field := fld.NewByteField(lines).NewFieldWithBorder('.')
	visited := containers.NewSet[fld.Pos]()
	var dfs func(fld.Pos) (int, int)
	dfs = func(pos fld.Pos) (s int, p int) {
		s += 1
		visited.Add(pos)
		for _, dir := range fld.DirsSimple {
			next := pos.Add(dir)
			if field.Get(next) == field.Get(pos) && !visited.Has(next) {
				sNext, pNext := dfs(next)
				s += sNext
				p += pNext
			} else if field.Get(next) != field.Get(pos) {
				p += 1
			}
		}
		return s, p
	}
	sum := 0
	for row := 1; row < field.Rows()-1; row++ {
		for col := 1; col < field.Cols()-1; col++ {
			pos := fld.NewPos(row, col)
			if !visited.Has(pos) {
				s, p := dfs(pos)
				sum += s * p
			}
		}
	}
	return sum
}

func SolvePart2(lines []string) any {
	field := fld.NewByteField(lines).NewFieldWithBorder('.')

	type FenceSection struct {
		double     fld.Pos // Position of center of the fence section. Double to avoid floating point.
		outsideDir fld.Pos
	}

	visited := containers.NewSet[fld.Pos]()
	var dfs func(fld.Pos) (int, []FenceSection)
	dfs = func(pos fld.Pos) (s int, p []FenceSection) {
		s += 1
		visited.Add(pos)
		for _, dir := range fld.DirsSimple {
			next := pos.Add(dir)
			if field.Get(next) == field.Get(pos) && !visited.Has(next) {
				sNext, pNext := dfs(next)
				s += sNext
				p = append(p, pNext...)
			} else if field.Get(next) != field.Get(pos) {
				p = append(p, FenceSection{
					double:     next.Add(pos),
					outsideDir: next.Sub(pos),
				})
			}
		}
		return s, p
	}
	countSides := func(perimeter []FenceSection) int {
		restSections := containers.NewSet(perimeter...)
		count := 0
		for len(restSections) > 0 {
			section := restSections.Any()
			restSections.Remove(section)
			count++
			var lookup []fld.Pos
			if section.outsideDir.Row == 0 {
				// This is a vertical side in the same row. Lookup other vertical sides above and below.
				lookup = []fld.Pos{fld.Up, fld.Down}
			} else {
				// This is a horizontal side in the same column. Lookup other horizontal sides left and right.
				lookup = []fld.Pos{fld.Left, fld.Right}
			}
			for _, dir := range lookup {
				dir = dir.Mult(2)
				next := section
				next.double = next.double.Add(dir)
				for restSections.Has(next) {
					restSections.Remove(next)
					next.double = next.double.Add(dir)
				}
			}
		}
		return count
	}
	sum := 0
	for row := 1; row < field.Rows()-1; row++ {
		for col := 1; col < field.Cols()-1; col++ {
			pos := fld.NewPos(row, col)
			if !visited.Has(pos) {
				s, p := dfs(pos)
				sum += s * countSides(p)
			}
		}
	}
	return sum
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
