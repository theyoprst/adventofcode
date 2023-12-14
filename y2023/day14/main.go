package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

const (
	Round = 'O'
	Cube  = '#'
	Empty = '.'
)

func TiltNorth(field [][]byte) {
	rows := len(field)
	cols := len(field[0])
	for col := 0; col < cols; col++ {
		stopRow := -1
		for row := 0; row < rows; row++ {
			ch := field[row][col]
			if ch == Cube {
				stopRow = row
			} else if ch == Round {
				stopRow++
				field[stopRow][col], field[row][col] = field[row][col], field[stopRow][col]
			}
		}
	}
}

func NorthLoad(field [][]byte) int {
	rows := len(field)
	cols := len(field[0])
	ans := 0
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			ch := field[row][col]
			if ch == Round {
				ans += rows - row
			}
		}
	}
	return ans
}

func Solve1(lines []string) any {
	field := aoc.MakeByteField(lines)
	TiltNorth(field)
	return NorthLoad(field)
}

func TiltCycle(field aoc.ByteField) aoc.ByteField {
	for i := 0; i < 4; i++ {
		TiltNorth(field)
		field = field.RotateRight()
	}
	return field
}

func Solve2(lines []string) any {
	field := aoc.MakeByteField(lines)

	seen := map[string]int{}
	i := 1
	cycle := 0
	for ; true; i++ {
		field = TiltCycle(field)

		fieldStr := field.String()
		if seen[fieldStr] == 0 {
			seen[fieldStr] = i
		} else {
			cycle = i - seen[fieldStr]
			break
		}
	}
	rest := (1000000000 - i) % cycle
	for i := 0; i < rest; i++ {
		field = TiltCycle(field)
		must.Greater(seen[field.String()], 0)
	}

	return NorthLoad(field)
}

var (
	solvers1 []aoc.Solver = []aoc.Solver{Solve1}
	solvers2 []aoc.Solver = []aoc.Solver{Solve2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
