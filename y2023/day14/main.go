package main

import (
	"fmt"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

const (
	Round = 'O'
	Cube  = '#'
	Empty = '.'
)

func Solve1(lines []string) any {
	field := aoc.ToBytesField(lines)
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

func ReverseColumns(field [][]byte) {
	for i := range field {
		slices.Reverse(field[i])
	}
}

func TiltCycle(field [][]byte) [][]byte {
	// North:
	field = aoc.Transpose(field)
	TiltWest(field)
	field = aoc.Transpose(field)
	// fmt.Println("\nNorth:")
	// PrintField(field)

	// West
	TiltWest(field)
	// fmt.Println("\nWest:")
	// PrintField(field)

	// South
	field = aoc.Transpose(field)
	ReverseColumns(field)
	TiltWest(field)
	ReverseColumns(field)
	field = aoc.Transpose(field)
	// fmt.Println("\nSouth:")
	// PrintField(field)

	// East
	ReverseColumns(field)
	TiltWest(field)
	ReverseColumns(field)
	// fmt.Println("\nEast:")
	// PrintField(field)

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
	field := aoc.ToBytesField(lines)

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
