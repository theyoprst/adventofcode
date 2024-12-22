package main

import (
	"context"
	"math"
	"strconv"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(_ context.Context, lines []string) any {
	return solve(lines, 25)
}

func SolvePart2(_ context.Context, lines []string) any {
	return solve(lines, 75)
}

func solve(lines []string, blinksCount int) any {
	cache := map[int]int{}
	var blinkRec func(x, n int) int
	blinkRec = func(x, n int) (result int) {
		must.Less(x, math.MaxInt/100)
		cacheKey := x*100 + n
		if v, ok := cache[cacheKey]; ok {
			return v
		}
		defer func() {
			cache[cacheKey] = result
		}()
		if n == 0 {
			return 1
		}
		if x == 0 {
			return blinkRec(1, n-1)
		}
		sx := strconv.Itoa(x)
		if len(sx)&1 == 1 {
			must.LessOrEqual(x, math.MaxInt/2024)
			return blinkRec(x*2024, n-1)
		}
		half := len(sx) / 2
		left := must.Atoi(sx[:half])
		right := must.Atoi(sx[half:])
		return blinkRec(left, n-1) + blinkRec(right, n-1)
	}
	sum := 0
	for _, x := range aoc.Ints(lines[0]) {
		sum += blinkRec(x, blinksCount)
	}
	return sum
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
