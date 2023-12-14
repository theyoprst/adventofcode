package main

import (
	"bytes"

	"github.com/theyoprst/adventofcode/aoc"
)

type Point struct {
	row, col int
}

var (
	East  Point = Point{0, 1}
	West  Point = Point{0, -1}
	North Point = Point{-1, 0}
	South Point = Point{1, 0}

	Dirs map[byte]map[Point]bool = map[byte]map[Point]bool{
		'S': aoc.ToSet([]Point{North, South, East, West}),
		'|': aoc.ToSet([]Point{North, South}),
		'-': aoc.ToSet([]Point{East, West}),
		'L': aoc.ToSet([]Point{East, North}),
		'J': aoc.ToSet([]Point{West, North}),
		'7': aoc.ToSet([]Point{South, West}),
		'F': aoc.ToSet([]Point{South, East}),
	}
)

func SolvePart1(lines []string) any {
	f := aoc.MakeByteField(lines).AddBorder('*')
	var start Point
	for row := range f {
		for col := range f[row] {
			if f[row][col] == 'S' {
				start.row = row
				start.col = col
			}
		}
	}
	p := start
	noway := Point{}
	steps := 0
	for steps == 0 || f[p.row][p.col] != 'S' {
		steps++
		ch := f[p.row][p.col]
		for dir := range Dirs[ch] {
			rev := Point{-dir.row, -dir.col}
			np := Point{p.row + dir.row, p.col + dir.col}
			if dir != noway && Dirs[f[np.row][np.col]][rev] {
				p = np
				noway = rev
				break
			}
		}
	}
	return steps / 2
}

func SolvePart2(lines []string) any {
	f := aoc.ByteField(make([][]byte, 2*len(lines)))
	for row := 0; row < len(f); row++ {
		f[row] = bytes.Repeat([]byte{' '}, 2*len(lines[0]))
	}
	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[row]); col++ {
			f[2*row][2*col] = lines[row][col]
		}
	}
	f = f.AddBorder('*').AddBorder('*')

	var start Point
	for row := range f {
		for col := range f[row] {
			if f[row][col] == 'S' {
				start.row = row
				start.col = col
			}
		}
	}

	p := start
	noway := Point{}
	steps := 0
	for steps == 0 || f[p.row][p.col] != 'S' {
		steps++
		ch := f[p.row][p.col]
		f[p.row][p.col] = 'S'
		for dir := range Dirs[ch] {
			rev := Point{-dir.row, -dir.col}
			np := Point{p.row + dir.row, p.col + dir.col}
			np2 := Point{p.row + 2*dir.row, p.col + 2*dir.col}
			if dir != noway && Dirs[f[np2.row][np2.col]][rev] {
				p = np2
				f[np.row][np.col] = 'S'
				noway = rev
				break
			}
		}
	}

	var fill func(p Point)
	fill = func(p Point) {
		ch := f[p.row][p.col]
		if ch == '*' || ch == 'S' {
			return
		}
		f[p.row][p.col] = '*'
		for _, dir := range []Point{East, West, South, North} {
			fill(Point{
				row: p.row + dir.row,
				col: p.col + dir.col,
			})
		}
	}
	for row := 2; row < len(f)-2; row++ {
		fill(Point{row: row, col: 2})
		fill(Point{row: row, col: len(f[row]) - 3})
	}
	for col := 2; col < len(f[0])-2; col++ {
		fill(Point{row: 2, col: col})
		fill(Point{row: len(f) - 3, col: col})
	}

	var ans2 int
	for row := range f {
		for col := range f[row] {
			ch := f[row][col]
			ans2 += aoc.BoolToInt(ch != 'S' && ch != '*' && ch != ' ')
		}
	}
	return ans2
}

var solversPart1 []aoc.Solver = []aoc.Solver{
	SolvePart1,
}

var solversPart2 []aoc.Solver = []aoc.Solver{
	SolvePart2,
	// TODO: try Shoelace formula and Pick's theorem: https://www.reddit.com/r/adventofcode/comments/18evyu9/comment/kcqu687/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
}

func main() {
	aoc.Main(solversPart1, solversPart2)
}
