package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

// Cheatsheet:
//
// Human readable regex:
//   rex.New(rex.Common.RawVerbose(``)).MustCompile()
//

func SolvePart1(lines []string) any {
	const search = "XMAS"
	f := fld.NewByteField(lines)
	count := 0
	for row := range f.Rows() {
		for col := range f.Cols() {
			pos := fld.NewPos(row, col)
		dirsLoop:
			for _, dir := range fld.DirsAll {
				for i := range len(search) {
					npos := pos.Add(dir.Mult(i))
					if !f.Inside(npos) || f.Get(npos) != search[i] {
						continue dirsLoop
					}
				}
				count++
			}
		}
	}
	return count
}

func SolvePart2(lines []string) any {
	const (
		search  = "MAS"
		halfLen = len(search) / 2
	)
	f := fld.NewByteField(lines)
	count := 0
	for row := halfLen; row < f.Rows()-halfLen; row++ {
		for col := halfLen; col < f.Cols()-halfLen; col++ {
			pos := fld.NewPos(row, col)
			matches := 0
			for _, dir := range fld.DirsDiag {
				word := make([]byte, 0, len(search))
				for i := range len(search) {
					npos := pos.Add(dir.Mult(i - halfLen))
					word = append(word, f.Get(npos))
				}
				matches += aoc.BoolToInt(string(word) == search)
			}
			count += aoc.BoolToInt(matches == 2)
		}
	}
	return count
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
