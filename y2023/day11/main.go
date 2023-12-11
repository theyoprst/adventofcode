package main

import (
	"github.com/theyoprst/adventofcode/aoc"
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
	field := aoc.ToBytesField(lines)
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
					rowDist += 1
				}
			}

			colDist := 0
			minCol := min(first.col, second.col)
			maxCol := max(first.col, second.col)
			for col := minCol; col < maxCol; col++ {
				if colStars[col] == 0 {
					colDist += expandFactor
				} else {
					colDist += 1
				}
			}

			ans += rowDist + colDist
		}
	}
	return ans
}

func solveGenericPartSum(lines []string, expandFactor int) int {
	field := aoc.ToBytesField(lines)
	rowStars := map[int]bool{}
	colStars := map[int]bool{}
	var stars []Point
	for row, line := range field {
		for col, ch := range line {
			if ch == Star {
				stars = append(stars, Point{row, col})
				rowStars[row] = true
				colStars[col] = true
			}
		}
	}

	sum := 0
	var rowsPart []int
	for row := 0; row < len(field); row++ {
		if rowStars[row] {
			sum += 1
		} else {
			sum += expandFactor
		}
		rowsPart = append(rowsPart, sum)
	}

	sum = 0
	var colsPart []int
	for col := 0; col < len(field[0]); col++ {
		if colStars[col] {
			sum += 1
		} else {
			sum += expandFactor
		}
		colsPart = append(colsPart, sum)
	}

	ans := 0
	for i, first := range stars {
		for _, second := range stars[i+1:] {
			ans += aoc.Abs(rowsPart[first.row] - rowsPart[second.row])
			ans += aoc.Abs(colsPart[first.col] - colsPart[second.col])
		}
	}
	return ans
}

var solversPart1 []aoc.Solver = []aoc.Solver{
	SolvePart1,
	SolvePart1Naive,
}

var solversPart2 []aoc.Solver = []aoc.Solver{
	SolvePart2,
	SolvePart2Naive,
}

func main() {
	aoc.Main([]aoc.Solver{SolvePart1}, []aoc.Solver{SolvePart2})
}
