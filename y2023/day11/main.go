package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

type Point struct {
	row int
	col int
}

const (
	Star  = '#'
	Space = '.'
)

func SolvePart1(lines []string) any {
	return solveGenericPartSum(lines, 2)
}

func SolvePart2(lines []string) any {
	return solveGenericPartSum(lines, 1000000)
}

func SolvePart1Naive(lines []string) any {
	return solveGenericNaive(lines, 2)
}

func SolvePart2Naive(lines []string) any {
	return solveGenericNaive(lines, 1000000)
}

func solveGenericNaive(lines []string, expandFactor int) int {
	field := fld.NewByteField(lines)
	rowStars := map[int]int{}
	colStars := map[int]int{}
	var stars []Point
	for row, line := range field {
		for col, ch := range line {
			if ch == Star {
				stars = append(stars, Point{row, col})
				rowStars[row]++
				colStars[col]++
			}
		}
	}

	ans := 0
	for i, first := range stars {
		for _, second := range stars[i+1:] {
			rowDist := 0
			minRow := min(first.row, second.row)
			maxRow := max(first.row, second.row)
			for row := minRow + 1; row <= maxRow; row++ {
				if rowStars[row] == 0 {
					rowDist += expandFactor
				} else {
					rowDist++
				}
			}

			colDist := 0
			minCol := min(first.col, second.col)
			maxCol := max(first.col, second.col)
			for col := minCol; col < maxCol; col++ {
				if colStars[col] == 0 {
					colDist += expandFactor
				} else {
					colDist++
				}
			}

			ans += rowDist + colDist
		}
	}
	return ans
}

func solveGenericPartSum(lines []string, expandFactor int) int {
	field := fld.NewByteField(lines)
	rows := len(field)
	cols := len(field[0])

	rowSizes := aoc.MakeSlice(expandFactor, rows)
	colSizes := aoc.MakeSlice(expandFactor, cols)
	var stars []Point
	for row, line := range field {
		for col, ch := range line {
			if ch == Star {
				stars = append(stars, Point{row, col})
				rowSizes[row] = 1
				colSizes[col] = 1
			}
		}
	}

	rowSums := aoc.PartialSum(rowSizes)
	colSums := aoc.PartialSum(colSizes)
	ans := 0
	for i, first := range stars {
		for _, second := range stars[i+1:] {
			ans += aoc.Abs(rowSums[first.row] - rowSums[second.row])
			ans += aoc.Abs(colSums[first.col] - colSums[second.col])
		}
	}
	return ans
}

var solversPart1 = []aoc.Solver{
	SolvePart1,
	SolvePart1Naive,
}

var solversPart2 = []aoc.Solver{
	SolvePart2,
	SolvePart2Naive,
}

func main() {
	aoc.Main([]aoc.Solver{SolvePart1}, []aoc.Solver{SolvePart2})
}
