package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(lines []string) any {
	return solve(lines, 0)
}

func SolvePart2(lines []string) any {
	return solve(lines, 10000000000000)
}

func solve(lines []string, shift int) int {
	sum := 0
	for _, block := range aoc.Blocks(lines) {
		ax, ay := parseTwoInts(block[0])
		bx, by := parseTwoInts(block[1])
		tx, ty := parseTwoInts(block[2])
		tx += shift
		ty += shift

		// ay * n + by * m = ty // *= ax
		// ax * n + bx * m = tx // *= ay
		// =>
		// ax * ay * n + ax * by * m = ty * ax
		// ax * ay * n + ay * bx * m = tx * ay
		// (minus) =>
		// m * (ax * by - ay * bx) = ty * ax - tx * ay

		m := divExact(ty*ax-tx*ay, ax*by-ay*bx)
		if m == -1 {
			continue
		}

		n := divExact(tx-bx*m, ax)
		if n == -1 {
			continue
		}

		sum += 3*n + m
	}
	return sum
}

func parseTwoInts(str string) (int, int) {
	a := aoc.Ints(str)
	must.Equal(len(a), 2)
	return a[0], a[1]
}

func divExact(a, b int) int {
	if b == 0 || a%b != 0 {
		return -1
	}
	return a / b
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
