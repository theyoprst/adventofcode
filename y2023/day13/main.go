package main

import (
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
)

func MirrirPoint(a []int, ignoreI int) int {
	for i := 1; i < len(a); i++ {
		if i == ignoreI {
			continue
		}
		p := slices.Clone(a[:i])
		slices.Reverse(p)
		s := a[i:]
		minLen := min(len(p), len(s))
		p = p[:minLen]
		s = s[:minLen]
		if slices.Equal(p, s) {
			// fmt.Printf("MirrorPoint %v: %d\n", a, i)
			return i
		}
	}
	return 0
}

func MirrorPoints(field [][]byte, ignoreHor, ignoreVert int) (int, int) {
	cols := make([]int, len(field[0]))
	rows := make([]int, len(field))
	for row, line := range field {
		for _, ch := range line {
			rows[row] *= 2
			if ch == '#' {
				rows[row]++
			}
		}
	}
	for _, line := range field {
		for col, ch := range line {
			cols[col] *= 2
			if ch == '#' {
				cols[col]++
			}
		}
	}
	horLine := MirrirPoint(cols, ignoreHor)
	vertLine := MirrirPoint(rows, ignoreVert)
	return horLine, vertLine
}

func SolvePart1(lines []string) any {
	var ans int
	for _, pattern := range aoc.Split(lines, "") {
		field := aoc.ToBytesField(pattern)
		hor, vert := MirrorPoints(field, -1, -1)
		ans += hor + 100*vert
	}
	return ans
}

func SolvePart2(lines []string) any {
	var ans int
	for _, pattern := range aoc.Split(lines, "") {
		field := aoc.ToBytesField(pattern)
		hor, vert := MirrorPoints(field, -1, -1)
		var newHor, newVert int
	fieldLoop:
		for row, line := range field {
			for col, ch := range line {
				var newCh byte
				if ch == '#' {
					newCh = '.'
				} else {
					newCh = '#'
				}
				field[row][col] = newCh
				newHor, newVert = MirrorPoints(field, hor, vert)
				if newHor != hor && newHor > 0 {
					break fieldLoop
				}
				if newVert != vert && newVert > 0 {
					break fieldLoop
				}
				field[row][col] = ch
			}
		}
		if newHor == hor {
			newHor = 0
		}
		if newVert == vert {
			newVert = 0
		}
		if newHor == 0 && newVert == 0 {
			panic("oops")
		}
		ans += newHor + 100*newVert
	}
	return ans
}

var solversPart1 []aoc.Solver = []aoc.Solver{
	SolvePart1,
}

var solversPart2 []aoc.Solver = []aoc.Solver{
	SolvePart2,
}

func main() {
	aoc.Main([]aoc.Solver{SolvePart1}, []aoc.Solver{SolvePart2})
}
