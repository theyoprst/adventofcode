package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

const (
	Round = 'O'
	Cube  = '#'
	Empty = '.'
)

func TiltNorth(field fld.ByteField) {
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
				field.Swap(
					fld.Pos{Row: stopRow, Col: col},
					fld.Pos{Row: row, Col: col},
				)
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
	field := fld.NewByteField(lines)
	TiltNorth(field)
	return NorthLoad(field)
}

func SolvePart2(lines []string) any {
	field := fld.NewByteField(lines)

	seen := map[string]int{}
	var cycle int
	const iters = 1000000000
	for i := 1; i <= iters; i++ {
		for k := 0; k < 4; k++ {
			TiltNorth(field)
			field = field.NewFieldRotatedRight()
		}
		fieldStr := fld.ToString(field)
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
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
