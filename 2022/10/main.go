package main

import (
	"context"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(_ context.Context, lines []string) any {
	registerX := 1
	cycle := 0
	signalStrengthSum := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		cycle++
		if cycle%40 == 20 {
			signalStrengthSum += cycle * registerX
		}
		if fields[0] == "addx" {
			cycle++
			if cycle%40 == 20 {
				signalStrengthSum += cycle * registerX
			}
			registerX += must.Atoi(fields[1])
		}
	}
	return signalStrengthSum
}

func SolvePart2(_ context.Context, lines []string) any {
	registerX := 1
	cycle := 0
	crt := make([][]byte, 6)
	for i := range crt {
		crt[i] = make([]byte, 40)
	}
	for _, line := range lines {
		fields := strings.Fields(line)
		drawPixel(crt, cycle, registerX)
		cycle++
		if fields[0] == "addx" {
			drawPixel(crt, cycle, registerX)
			cycle++
			registerX += must.Atoi(fields[1])
		}
	}
	var output strings.Builder
	for _, row := range crt {
		output.Write(row)
		output.WriteByte('\n')
	}
	return output.String()
}

func drawPixel(crt [][]byte, cycle, x int) {
	row := cycle / 40
	col := cycle % 40
	if aoc.Abs(col-x) <= 1 {
		crt[row][col] = '#'
	} else {
		crt[row][col] = '.'
	}
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
