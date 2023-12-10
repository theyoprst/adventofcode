package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
)

type Point struct {
	row, col int
}

var (
	E Point = Point{0, 1}
	W Point = Point{0, -1}
	N Point = Point{-1, 0}
	S Point = Point{1, 0}

	Dirs map[byte]map[Point]bool = map[byte]map[Point]bool{
		'S': aoc.ToSet([]Point{N, S, E, W}),
		'|': aoc.ToSet([]Point{N, S}),
		'-': aoc.ToSet([]Point{E, W}),
		'L': aoc.ToSet([]Point{E, N}),
		'J': aoc.ToSet([]Point{W, N}),
		'7': aoc.ToSet([]Point{S, W}),
		'F': aoc.ToSet([]Point{S, E}),
	}
)

func AddBorder2D(a [][]byte, b byte) [][]byte {
	cols := len(a[0]) + 2
	res := make([][]byte, 0, len(a)+2)
	res = append(res, bytes.Repeat([]byte{b}, cols))
	for _, s := range a {
		line := append(append([]byte{b}, s...), b)
		res = append(res, line)
	}
	res = append(res, bytes.Repeat([]byte{b}, cols))
	return res
}

func SolvePart1(lines []string) any {
	f := aoc.AddBorder2D(lines, '*')
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
	f := make([][]byte, 2*len(lines))
	for row := 0; row < len(f); row++ {
		f[row] = bytes.Repeat([]byte{' '}, 2*len(lines[0]))
	}
	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[row]); col++ {
			f[2*row][2*col] = lines[row][col]
		}
	}
	f = AddBorder2D(f, '*')
	f = AddBorder2D(f, '*')

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
		for _, dir := range []Point{E, W, S, N} {
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

func main() {
	lines1 := aoc.ReadInputLines()
	lines2 := slices.Clone(lines1)
	var cmd string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	if cmd != "part2" || cmd == "part1" {
		fmt.Println("Part 1:", SolvePart1(lines1))
	}
	if cmd != "part1" || cmd == "part2" {
		fmt.Println("Part 2:", SolvePart2(lines2))
	}
}
