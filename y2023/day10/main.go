package main

import (
	"bytes"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

var Dirs map[byte]containers.Set[fld.Pos] = map[byte]containers.Set[fld.Pos]{
	'S': containers.NewSet[fld.Pos](fld.North, fld.South, fld.East, fld.West),
	'|': containers.NewSet[fld.Pos](fld.North, fld.South),
	'-': containers.NewSet[fld.Pos](fld.East, fld.West),
	'L': containers.NewSet[fld.Pos](fld.East, fld.North),
	'J': containers.NewSet[fld.Pos](fld.West, fld.North),
	'7': containers.NewSet[fld.Pos](fld.South, fld.West),
	'F': containers.NewSet[fld.Pos](fld.South, fld.East),
}

func SolvePart1(lines []string) any {
	f := fld.NewByteField(lines).AddBorder('*')
	start := f.FindFirst('S')
	p := start
	noway := fld.Pos{}
	steps := 0
	for steps == 0 || f[p.Row][p.Col] != 'S' {
		steps++
		ch := f[p.Row][p.Col]
		for dir := range Dirs[ch] {
			rev := dir.Mult(-1)
			np := p.Add(dir)
			if dir != noway && Dirs[f[np.Row][np.Col]].Has(rev) {
				p = np
				noway = rev
				break
			}
		}
	}
	return steps / 2
}

func SolvePart2(lines []string) any {
	f := fld.ByteField(make([][]byte, 2*len(lines)))
	for row := 0; row < len(f); row++ {
		f[row] = bytes.Repeat([]byte{' '}, 2*len(lines[0]))
	}
	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[row]); col++ {
			f[2*row][2*col] = lines[row][col]
		}
	}
	f = f.AddBorder('*').AddBorder('*')

	start := f.FindFirst('S')

	p := start
	noway := fld.Pos{}
	steps := 0
	for steps == 0 || f[p.Row][p.Col] != 'S' {
		steps++
		ch := f[p.Row][p.Col]
		f[p.Row][p.Col] = 'S'
		for dir := range Dirs[ch] {
			rev := dir.Mult(-1)
			np := p.Add(dir)
			np2 := p.Add(dir.Mult(2))
			if dir != noway && Dirs[f[np2.Row][np2.Col]].Has(rev) {
				p = np2
				f.Set(np, 'S')
				noway = rev
				break
			}
		}
	}

	var fill func(p fld.Pos)
	fill = func(p fld.Pos) {
		ch := f.Get(p)
		if ch == '*' || ch == 'S' {
			return
		}
		f.Set(p, '*')
		for _, dir := range []fld.Pos{fld.East, fld.West, fld.South, fld.North} {
			fill(p.Add(dir))
		}
	}
	for row := 2; row < len(f)-2; row++ {
		fill(fld.NewPos(row, 2))
		fill(fld.NewPos(row, f.Cols()-3))
	}
	for col := 2; col < len(f[0])-2; col++ {
		fill(fld.NewPos(2, col))
		fill(fld.NewPos(f.Rows()-3, col))
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
