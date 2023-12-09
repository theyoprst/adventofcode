package main

import (
	"fmt"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func Next(a []int) int {
	a = aoc.Reversed(a)
	next := 0
	for size := len(a); size >= 2; size-- {
		next += a[0]
		for i := 0; i < size-1; i++ {
			a[i] -= a[i+1]
		}
	}
	must.Equal(a[0], 0)
	return next
}

func main() {
	lines := aoc.ReadInputLines()
	var ans1, ans2 int
	for _, line := range lines {
		a := aoc.Ints(line)
		ans1 += Next(a)
		ans2 += Next(aoc.Reversed(a))
	}
	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}
