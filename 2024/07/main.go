package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
)

func SolvePart1(lines []string) any {
	return solveGeneric(lines, false)
}

func SolvePart2(lines []string) any {
	return solveGeneric(lines, true)
}

func solveGeneric(lines []string, allowConcats bool) int {
	ans := 0
	for _, line := range lines {
		nn := aoc.Ints(line)
		testValue := nn[0]
		nn = nn[1:]

		want := containers.NewSet(testValue)
		for i := len(nn) - 1; i > 0; i-- {
			// Go backwards trying to fugure out whether the test value can be obtained by applying some operation
			// on some number x and the current number y. Because the expressions are evaluated from left to right,
			// and we're rolling them back (undo), we have to go backwards.
			//
			// For the multiplication and concatenation, that doesn't generate irrelevant branches because of
			// the modulo checks. That runs dozens times faster than the naive approach of generating all possible
			// expressions: 0.01s for both parts vs 0.05s + 0.45s for naive approach.
			y := nn[i]
			newWant := containers.NewSet[int]()
			for w := range want {
				// Undo summation: w = x + y
				if x := w - y; x > 0 {
					newWant.Add(x)
				}
				// Undo product: w = x * y
				if w%y == 0 {
					newWant.Add(w / y)
				}
				if allowConcats {
					// Undo concatenation:
					// w = x || y
					// w = x * 10^k + y
					k := pow10gt(y)
					x := w / k
					if x > 0 && w%k == y {
						newWant.Add(x)
					}
				}
			}
			want = newWant
		}
		if want.Has(nn[0]) {
			ans += testValue
		}
	}
	return ans
}

func pow10gt(x int) int {
	for i := 10; ; i *= 10 {
		if i > x {
			return i
		}
	}
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
