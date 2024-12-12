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
		insidePos  fld.Pos
		outsideDir fld.Pos
	}

	visited := containers.NewSet[fld.Pos]()
	var dfs func(fld.Pos, containers.Set[FenceSection]) int
	dfs = func(pos fld.Pos, perimeter containers.Set[FenceSection]) int {
		area := 1
		visited.Add(pos)
		for _, dir := range fld.DirsSimple {
			next := pos.Add(dir)
			if field.Get(next) == field.Get(pos) && !visited.Has(next) {
				area += dfs(next, perimeter)
			} else if field.Get(next) != field.Get(pos) {
				perimeter.Add(FenceSection{
					insidePos:  pos,
					outsideDir: dir,
				})
			}
		}
		return area
	}

	countSides := func(perimeter containers.Set[FenceSection]) int {
		count := 0
		for len(perimeter) > 0 {
			count++
			section := perimeter.Any()
			perimeter.Remove(section)
			perpendiculars := []fld.Pos{
				section.outsideDir.RotateClockwise(),
				section.outsideDir.RotateCounterClockwise(),
			}
			for _, dir := range perpendiculars {
				next := section
				next.insidePos = next.insidePos.Add(dir)
				for perimeter.Has(next) {
					perimeter.Remove(next)
					next.insidePos = next.insidePos.Add(dir)
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
				perimeter := containers.NewSet[FenceSection]()
				area := dfs(pos, perimeter)
				sum += area * countSides(perimeter)
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
