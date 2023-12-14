package main

import (
	"fmt"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

const (
	Round = 'O'
	Cube  = '#'
	Empty = '.'
)

func Solve1(lines []string) any {
	field := aoc.MakeByteField(lines)
	rows := len(field)
	cols := len(field[0])
	ans := 0
	for col := 0; col < cols; col++ {
		stopRow := -1
		for row := 0; row < rows; row++ {
			ch := field[row][col]
			if ch == Cube {
				stopRow = row
			} else if ch == Round {
				stopRow++
				ans += (rows - stopRow)
			}
		}
	}
	return ans
}

func TiltWest(field [][]byte) {
	rows := len(field)
	cols := len(field[0])
	for row := 0; row < rows; row++ {
		stopCol := -1
		for col := 0; col < cols; col++ {
			ch := field[row][col]
			if ch == Cube {
				stopCol = col
			} else if ch == Round {
				stopCol++
				field[row][stopCol], field[row][col] = field[row][col], field[row][stopCol]
			}
		}
	}
}

func PrintField(field [][]byte) {
	for _, line := range field {
		fmt.Println(string(line))
	}
}

func ToString(field [][]byte) string {
	s := ""
	for _, line := range field {
		s += string(line)
	}
	return s
}

func TiltCycle(field aoc.ByteField) aoc.ByteField {
	// North:
	field = field.Transpose()
	TiltWest(field)
	field = field.Transpose()

	// West
	TiltWest(field)

	// South
	field = field.Transpose()
	field.ReverseColumns()
	TiltWest(field)
	field.ReverseColumns()
	field = field.Transpose()

	// East
	field.ReverseColumns()
	TiltWest(field)
	field.ReverseColumns()

	return field
}

func NorthLoad(field [][]byte) int {
	rows := len(field)
	cols := len(field[0])
	ans := 0
	for col := 0; col < cols; col++ {
		stopRow := -1
		for row := 0; row < rows; row++ {
			ch := field[row][col]
			if ch == Cube {
				stopRow = row
			} else if ch == Round {
				stopRow++
				ans += rows - stopRow
			}
		}
	}
	return ans
}

func NorthLoad2(field [][]byte) int {
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

func Solve2(lines []string) any {
	field := aoc.MakeByteField(lines)

	seen := map[string]int{}
	i := 1
	cycle := 0
	for ; true; i++ {
		field = TiltCycle(field)

		fieldStr := ToString(field)
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
		must.Greater(seen[ToString(field)], 0)
	}

	return NorthLoad2(field)
}

var (
	solvers1 []aoc.Solver = []aoc.Solver{Solve1}
	solvers2 []aoc.Solver = []aoc.Solver{Solve2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
