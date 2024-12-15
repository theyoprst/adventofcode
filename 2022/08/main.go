// 15:05-15:12-15:20 - 15:31 (26 min to solve)
package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

func SolvePart1(lines []string) any {
	field := fld.NewByteField(lines)
	visible := containers.NewSet[fld.Pos]()
	for row := 0; row < field.Rows(); row++ {
		maxTree := byte(0)
		for col := 0; col < field.Cols(); col++ {
			p := fld.NewPos(row, col)
			if field.Get(p) > maxTree {
				visible.Add(p)
				maxTree = field.Get(p)
			}
		}

		maxTree = byte(0)
		for col := field.Cols() - 1; col >= 0; col-- {
			p := fld.NewPos(row, col)
			if field.Get(p) > maxTree {
				visible.Add(p)
				maxTree = field.Get(p)
			}
		}
	}

	for col := 0; col < field.Cols(); col++ {
		maxTree := byte(0)
		for row := 0; row < field.Rows(); row++ {
			p := fld.NewPos(row, col)
			if field.Get(p) > maxTree {
				visible.Add(p)
				maxTree = field.Get(p)
			}
		}

		maxTree = byte(0)
		for row := field.Rows() - 1; row >= 0; row-- {
			p := fld.NewPos(row, col)
			if field.Get(p) > maxTree {
				visible.Add(p)
				maxTree = field.Get(p)
			}
		}
	}
	return len(visible)
}

func SolvePart2(lines []string) any {
	field := fld.NewByteField(lines)
	maxProduct := 0
	for pos := range field.IterPositions() {
		p := 1
		for _, dir := range fld.DirsSimple {
			c := 0
			next := pos
			for {
				next = next.Add(dir)
				if !field.Inside(next) {
					break
				}
				c++
				if field.Get(next) >= field.Get(pos) {
					break
				}
			}
			p *= c
		}
		maxProduct = max(maxProduct, p)
	}
	return maxProduct
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
