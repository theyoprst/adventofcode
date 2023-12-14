package main

import (
	"github.com/theyoprst/adventofcode/aoc"
)

const (
	Round = 'O'
	Cube  = '#'
	Empty = '.'
)

func TiltNorth(field aoc.ByteField) {
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
				field.Swap(stopRow, col, row, col)
			}
		}
	}
}

func NorthLoad(field [][]byte) int {
	rows := len(field)
	cols := len(field[0])
	ans := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			ch := field[row][col]
			if ch == Round {
				ans += rows - row
			}
		}
	}
	return ans
}

func SolvePart1(lines []string) any {
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

func SolvePart2(lines []string) any {
	field := aoc.MakeByteField(lines)

	seen := map[string]int{}
	cycle := 0
	const iters = 1000000000
	for i := 1; i <= iters; i++ {
		field = TiltCycle(field)
		fieldStr := field.String()
		if seen[fieldStr] == 0 {
			seen[fieldStr] = i
		} else {
			cycle = i - seen[fieldStr]
			i += cycle * ((iters - i) / cycle)
		}
	}

	return NorthLoad(field)
}

var (
	solvers1 []aoc.Solver = []aoc.Solver{SolvePart1}
	solvers2 []aoc.Solver = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
