package main

import (
	"fmt"

	"github.com/theyoprst/adventofcode/aoc"
)

func main() {
	ans := 1
	lines := aoc.ReadInputLines()
	tt := aoc.Ints(lines[0])
	rr := aoc.Ints(lines[1])
	for i := range tt {
		wins := 0
		for a := 0; a <= tt[i]; a++ {
			b := tt[i] - a
			wins += aoc.BoolToInt(a*b > rr[i])
		}
		ans *= wins
	}

	fmt.Println("Answer:", ans)
}
