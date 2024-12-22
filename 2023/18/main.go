package main

import (
	"context"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
	"github.com/theyoprst/adventofcode/must"
)

type Command struct {
	dir   fld.Pos
	steps int
}

var dirs map[string]fld.Pos = map[string]fld.Pos{
	"R": fld.Right,
	"L": fld.Left,
	"U": fld.Up,
	"D": fld.Down,

	"0": fld.Right,
	"1": fld.Down,
	"2": fld.Left,
	"3": fld.Up,
}

func ParseCommands(lines []string) []Command {
	var commands []Command
	for _, line := range lines {
		dirCode, stepsS, _ := must.Split3(line, " ")
		commands = append(commands, Command{dir: dirs[dirCode], steps: must.Atoi(stepsS)})
	}
	return commands
}

func ParseCommands2(lines []string) []Command {
	var commands []Command
	for _, line := range lines {
		_, _, hexStr := must.Split3(line, " ")
		steps := parseHex(hexStr[2:7])
		commands = append(commands, Command{dir: dirs[string(hexStr[7])], steps: steps})
	}
	return commands
}

func SolvePart1(_ context.Context, lines []string) any {
	commands := ParseCommands(lines)
	pos, rows, cols := GetStartAndSize(commands)
	field := NewFieldBySize(rows, cols, '.')
	field.Set(pos, '#')
	for _, cmd := range commands {
		for i := 0; i < cmd.steps; i++ {
			pos = pos.Add(cmd.dir)
			field.Set(pos, '#')
		}
	}
	// TODO: Refactor dfs function to Fill field's method, reuse in day10
	// Also reuse in day10 solution.
	var dfs func(p fld.Pos)
	dfs = func(p fld.Pos) {
		if !field.Inside(p) {
			return
		}
		if field.Get(p) != '.' {
			return
		}
		field.Set(p, '*')
		for _, dir := range []fld.Pos{fld.Left, fld.Right, fld.Up, fld.Down} {
			dfs(p.Add(dir))
		}
	}
	for row := 0; row < rows; row++ {
		dfs(fld.NewPos(row, 0))
		dfs(fld.NewPos(row, cols-1))
	}
	for col := 0; col < cols; col++ {
		dfs(fld.NewPos(0, col))
		dfs(fld.NewPos(rows-1, col))
	}

	ans := 0
	for _, line := range field {
		for _, ch := range line {
			if ch == '.' || ch == '#' {
				ans++
			}
		}
	}
	return ans
}

func SolvePart1ByCompression(_ context.Context, lines []string) any {
	commands := ParseCommands(lines)
	return SolveByCompression(commands)
}

func SolvePart2(_ context.Context, lines []string) any {
	commands := ParseCommands2(lines)
	return SolveByCompression(commands)
}

func parseHex(s string) int {
	n := 0
	for _, ch := range s {
		n = n*16 + hexVal(ch)
	}
	return n
}

func hexVal(ch rune) int {
	if '0' <= ch && ch <= '9' {
		return int(ch - '0')
	}
	if 'a' <= ch && ch <= 'f' {
		return int(ch-'a') + 10
	}
	panic(ch)
}

func NewFieldBySize(rows, cols int, fill byte) fld.ByteField {
	field := make(fld.ByteField, rows)
	for row := range field {
		field[row] = aoc.MakeSlice(fill, cols)
	}
	return field
}

func GetStartAndSize(commands []Command) (start fld.Pos, rows, cols int) {
	minP := fld.Zero
	maxP := fld.Zero
	pos := fld.Zero
	for _, cmd := range commands {
		pos = pos.Add(cmd.dir.Mult(cmd.steps))
		minP.Row = min(minP.Row, pos.Row)
		minP.Col = min(minP.Col, pos.Col)
		maxP.Row = max(maxP.Row, pos.Row)
		maxP.Col = max(maxP.Col, pos.Col)
	}
	must.Equal(pos, fld.Zero)
	return minP.Mult(-1), maxP.Row - minP.Row + 1, maxP.Col - minP.Col + 1
}

func SolveByCompression(commands []Command) any {
	pos, _, _ := GetStartAndSize(commands)
	rowSet := containers.NewSet[int](pos.Row)
	colSet := containers.NewSet[int](pos.Col)
	for _, cmd := range commands {
		pos = pos.Add(cmd.dir.Mult(cmd.steps))
		rowSet.Add(pos.Row, pos.Row+1)
		colSet.Add(pos.Col, pos.Col+1)
	}

	rowVals := rowSet.Slice()
	colVals := colSet.Slice()
	slices.Sort(rowVals)
	slices.Sort(colVals)

	rows := len(rowVals) - 1
	cols := len(colVals) - 1
	field := NewFieldBySize(rows, cols, '.')
	ipos := fld.Zero
	for rowVals[ipos.Row] < pos.Row {
		ipos.Row++
	}
	for colVals[ipos.Col] < pos.Col {
		ipos.Col++
	}
	field.Set(ipos, '#')
	for _, cmd := range commands {
		npos := ipos
		for aoc.Abs(rowVals[npos.Row]-rowVals[ipos.Row])+aoc.Abs(colVals[npos.Col]-colVals[ipos.Col]) < cmd.steps {
			npos = npos.Add(cmd.dir)
			field.Set(npos, '#')
		}
		ipos = npos
	}
	var dfs func(p fld.Pos)
	dfs = func(p fld.Pos) {
		if !field.Inside(p) {
			return
		}
		if field.Get(p) != '.' {
			return
		}
		field.Set(p, '*')
		for _, dir := range []fld.Pos{fld.Left, fld.Right, fld.Up, fld.Down} {
			dfs(p.Add(dir))
		}
	}
	for row := 0; row < rows; row++ {
		dfs(fld.NewPos(row, 0))
		dfs(fld.NewPos(row, cols-1))
	}
	for col := 0; col < cols; col++ {
		dfs(fld.NewPos(0, col))
		dfs(fld.NewPos(rows-1, col))
	}

	ans := 0
	for row, line := range field {
		for col, ch := range line {
			if ch == '.' || ch == '#' {
				ans += (rowVals[row+1] - rowVals[row]) * (colVals[col+1] - colVals[col])
			}
		}
	}
	return ans
}

var (
	solvers1 = []aoc.Solver{SolvePart1, SolvePart1ByCompression, SolvePart1Gauss}
	solvers2 = []aoc.Solver{SolvePart2, SolvePart2Gauss}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
